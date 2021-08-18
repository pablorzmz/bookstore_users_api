package users

import (
	"github.com/pablorzmz/bookstore_users_api/src/datasources/mysql/users_db"
	"github.com/pablorzmz/bookstore_users_api/src/utils/errors"
)

const (
	queryInsertUser       = "INSERT INTO users(first_name,last_name,email,date_created, password, status) VALUES(?,?,?,?,?,?);"
	queryGetUser          = "SELECT id, first_name, last_name, email, date_created, status FROM users WHERE id=?;"
	queryUpdateUser       = "UPDATE users SET first_name=?, last_name=?, email=? WHERE id=?;"
	queryDeleteUser       = "DELETE FROM users WHERE id = ?;"
	queryFindUserByStatus = "SELECT id, first_name, last_name, email, date_created, status FROM users WHERE status = ?;"
)

func (user *User) Save() *errors.RestErr {

	stmt, err := users_db.Client.Prepare(queryInsertUser)
	if err != nil {
		return errors.NewInteralServerError(err.Error())
	}

	defer stmt.Close()

	inserResult, err := stmt.Exec(user.FirstName, user.LastName, user.Email, user.DateCreated, user.Password, user.Status)
	if err != nil {
		return errors.NewInteralServerError("Error when saving user." + err.Error())
	}

	userId, err := inserResult.LastInsertId()
	if err != nil {
		return errors.NewInteralServerError("Could not get last insert index.")
	}

	user.Id = userId

	return nil
}

func (user *User) Get() *errors.RestErr {

	stmt, err := users_db.Client.Prepare(queryGetUser)
	if err != nil {
		return errors.NewInteralServerError(err.Error())
	}

	defer stmt.Close()

	result := stmt.QueryRow(user.Id)

	if err := result.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.DateCreated, &user.Status); err != nil {
		return errors.NewInteralServerError(err.Error())
	}

	return nil
}

func (user *User) Update() *errors.RestErr {

	stmt, err := users_db.Client.Prepare(queryUpdateUser)
	if err != nil {
		return errors.NewInteralServerError(err.Error())
	}

	defer stmt.Close()

	_, err = stmt.Exec(user.FirstName, user.LastName, user.Email, user.Id)
	if err != nil {
		return errors.NewInteralServerError(err.Error())
	}

	return nil
}

func (user *User) Delete() *errors.RestErr {
	stmt, err := users_db.Client.Prepare(queryDeleteUser)

	if err != nil {
		return errors.NewInteralServerError(err.Error())
	}

	_, err = stmt.Exec(user.Id)
	if err != nil {
		return errors.NewInteralServerError(err.Error())
	}

	defer stmt.Close()

	return nil
}

func (user *User) FindByStatus(status string) ([]User, *errors.RestErr) {
	stmt, err := users_db.Client.Prepare(queryFindUserByStatus)

	if err != nil {
		return nil, errors.NewInteralServerError(err.Error())
	}

	// Statement is no nil here
	defer stmt.Close()

	rows, err := stmt.Query(status)
	if err != nil {
		return nil, errors.NewBadRequestError(err.Error())
	}

	// Rows are not nil here
	defer rows.Close()

	results := make([]User, 0)
	for rows.Next() {
		var user User
		if err := rows.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.DateCreated, &user.Status); err != nil {
			return nil, errors.NewInteralServerError(err.Error())
		}
		results = append(results, user)
	}

	if len(results) == 0 {
		return nil, errors.NotFoundError("no users found matching status = " + status)
	}

	return results, nil
}

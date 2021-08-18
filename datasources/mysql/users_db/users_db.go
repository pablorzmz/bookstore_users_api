package users_db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var (
	Client *sql.DB
)

const (
	mysql_users_username = "root"
	mysql_users_password = "admin"
	mysql_users_host     = "127.0.0.1:3306"
	mysql_users_schema   = "users_db"
)

func init() {
	datasourceName := fmt.Sprintf("%s:%s@tcp(%s)/%s",
		mysql_users_username,
		mysql_users_password,
		mysql_users_host,
		mysql_users_schema,
	)

	var err error
	Client, err = sql.Open("mysql", datasourceName)
	if err != nil {
		panic(err)
	}

	if err = Client.Ping(); err != nil {
		panic(err)
	}

	log.Println("Database sucessfully configured")

}

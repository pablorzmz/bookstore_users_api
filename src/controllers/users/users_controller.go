package users

import (
	"net/http"
	"strconv"

	"github.com/pablorzmz/bookstore_users_api/src/services"
	"github.com/pablorzmz/bookstore_users_api/src/utils/errors"

	"github.com/gin-gonic/gin"
	"github.com/pablorzmz/bookstore_users_api/src/domain/users"
)

func getUserId(userIdParam string) (int64, *errors.RestErr) {
	userId, userErr := strconv.ParseInt(userIdParam, 10, 64)
	if userErr != nil {
		return 0, errors.NewBadRequestError("Invalid user id")
	}
	return userId, nil
}

func GetUser(c *gin.Context) {
	userId, userErr := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if userErr != nil {
		err := errors.NewBadRequestError("Invalid user id")
		c.JSON(err.Code, err)
	}

	user, getErr := services.UserService.GetUser(userId)
	if getErr != nil {
		c.JSON(getErr.Code, getErr)
		return
	}

	c.JSON(http.StatusOK, user.Marshall(c.GetHeader("X-Public") == "true"))
}

func CreateUser(c *gin.Context) {

	var user users.User

	if err := c.ShouldBindJSON(&user); err != nil {
		restErr := errors.NewBadRequestError(err.Error())
		c.JSON(restErr.Code, restErr)
		return
	}

	result, saveErr := services.UserService.CreateUser(user)
	if saveErr != nil {
		c.JSON(saveErr.Code, saveErr)
		return
	}

	c.JSON(http.StatusCreated, result.Marshall(c.GetHeader("X-Public") == "true"))
}

func UpdateUser(c *gin.Context) {

	var user users.User

	if err := c.ShouldBindJSON(&user); err != nil {
		restErr := errors.NewBadRequestError(err.Error())
		c.JSON(restErr.Code, restErr)
		return
	}

	isPartial := c.Request.Method == http.MethodPatch

	result, err := services.UserService.UpdateUser(isPartial, user)
	if err != nil {
		c.JSON(err.Code, err)
	}

	c.JSON(http.StatusOK, result.Marshall(c.GetHeader("X-Public") == "true"))
}

func Delete(c *gin.Context) {
	userId, userErr := getUserId(c.Param("user_id"))
	if userErr != nil {
		c.JSON(userErr.Code, userErr)
		return
	}
	if err := services.UserService.DeleteUser(userId); err != nil {
		c.JSON(userErr.Code, userErr)
		return
	}
	c.JSON(http.StatusOK, map[string]string{"status": "deleted"})
}

func Search(c *gin.Context) {
	status := c.Query("status")

	users, err := services.UserService.FindByStatus(status)

	if err != nil {
		c.JSON(err.Code, err)
		return
	}

	c.JSON(http.StatusOK, users.Marshall(c.GetHeader("X-Public") == "true"))
}

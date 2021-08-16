package app

import (
	"github.com/pablorzmz/bookstore_users_api/controllers/ping"
	"github.com/pablorzmz/bookstore_users_api/controllers/user"
)

func mapUrls() {
	router.GET("/ping", ping.Ping)
	router.POST("/users", user.CreateUser)
	router.GET("/users/:user_id", user.GetUser)
	router.GET("/users/search", user.SearchUser)
}

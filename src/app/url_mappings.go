package app

import (
	ping "github.com/pablorzmz/bookstore_users_api/src/controllers/ping"
	user "github.com/pablorzmz/bookstore_users_api/src/controllers/users"
)

func mapUrls() {
	router.GET("/ping", ping.Ping)
	router.POST("/users", user.CreateUser)
	router.GET("/users/:user_id", user.GetUser)
	router.PUT("/users/:user_id", user.UpdateUser)
	router.PATCH("/users/:user_id", user.UpdateUser)
	router.DELETE("/users/:user_id", user.Delete)
	router.GET("/internal/users/search", user.Search)
}

package app

import (
	"github.com/pablorzmz/bookstore_users_api/controllers"
)

func mapUrls() {
	router.GET("/ping", controllers.Ping)
}

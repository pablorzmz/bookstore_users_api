package app

import (
	"github.com/gin-gonic/gin"
	"github.com/pablorzmz/bookstore_users_api/src/logger"
)

var (
	router = gin.Default()
)

func StarApplication() {
	mapUrls()
	logger.Info("About to start the aplicacion")
	router.Run(":8080")
}

package mappings

import (
	"github.com/yunpeng1234/go-api/GoBackend/controllers"

	"github.com/gin-gonic/gin"
)

var Router *gin.Engine

func CreateUrlMappings() {
	Router = gin.Default()
	Router.Use(controllers.Cors())
	api := Router.Group("/api")
	{
		api.POST("/register", controllers.registerStudent)
		api.GET("/commonstudents", controllers.getCommonStudents)
		api.POST("/suspend", controllers.suspendStudent)
		api.POST("/retrievefornotifications", controllers.retrieveNotifications)
	}
}

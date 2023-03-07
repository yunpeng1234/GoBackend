package mappings

import (
	"github.com/yunpeng1234/GoBackend/controllers"

	"github.com/gin-gonic/gin"
)

var Router *gin.Engine

func CreateUrlMappings() {
	Router = gin.Default()
	Router.Use(controllers.Cors())
	api := Router.Group("/api")
	{
		api.POST("/register", controllers.RegisterStudent)
		api.GET("/commonstudents", controllers.GetCommonStudents)
		api.POST("/suspend", controllers.SuspendStudent)
		api.POST("/retrievefornotifications", controllers.RetrieveNotifications)
	}
}

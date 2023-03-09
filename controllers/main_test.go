package controllers

import (
	"github.com/gin-gonic/gin"
)

// func TestMain(m *testing.M) {
// 	gin.SetMode(gin.TestMode)
// 	setup()
// 	exitCode := m.Run()
// 	teardown()

// 	os.Exit(exitCode)
// }

func SetUpRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/api/commonstudents", func(c *gin.Context) {
		c.String(200, "pong")
	})
	r.POST("/api/register", func(c *gin.Context) {
		c.String(204, "pong")
	})
	r.POST("/api/suspend", func(c *gin.Context) {
		c.String(204, "pong")
	})
	r.POST("/api/retrievenotification", func(c *gin.Context) {
		c.String(200, "pong")
	})
	return r
}

func MainTest() {
	r := SetUpRouter()
	r.Run(":8080")
}

// func setup() {
// 	err := godotenv.Load()
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	controllers.InitDb()
// 	// Listen and server on 0.0.0.0:8080
// 	mappings.Router.Run(":8080")
// 	fmt.Println("Stop")
// }

// func teardown() {
// 	controllers.TearDown()
// }

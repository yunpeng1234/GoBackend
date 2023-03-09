package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/yunpeng1234/GoBackend/controllers"
	"github.com/yunpeng1234/GoBackend/mappings"
)

func TestMain(m *testing.M) {
	gin.SetMode(gin.TestMode)
	setup()
	exitCode := m.Run()
	teardown()

	os.Exit(exitCode)
}

func router() *gin.Engine {
	router := gin.Default()

	testRoutes := router.Group("/api")
	testRoutes.POST("/register", controllers.RegisterStudent)
	testRoutes.GET("/commonstudents", controllers.GetCommonStudents)
	testRoutes.POST("/suspend", controllers.SuspendStudent)
	testRoutes.POST("/retrievefornotifications", controllers.RetrieveNotifications)

	return router
}

func setup() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	controllers.InitDb()
	// Listen and server on 0.0.0.0:8080
	mappings.Router.Run(":8080")
	fmt.Println("Stop")
}

func teardown() {
	controllers.TearDown()
}

func MakeRequest(method, url string, body interface{}) *httptest.ResponseRecorder {
	requestBody, _ := json.Marshal(body)
	request, _ := http.NewRequest(method, url, bytes.NewBuffer(requestBody))

	writer := httptest.NewRecorder()
	router().ServeHTTP(writer, request)
	return writer
}

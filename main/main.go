package main

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"github.com/yunpeng1234/GoBackend/controllers"
	"github.com/yunpeng1234/GoBackend/mappings"
)

func main() {
	godotenv.Load()
	mappings.CreateUrlMappings()
	controllers.InitDb()
	// Listen and server on 0.0.0.0:8080
	mappings.Router.Run(":8080")
	fmt.Println("Stop")
}

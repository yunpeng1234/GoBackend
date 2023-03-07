package controllers

import (
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/yunpeng1234/GoBackend/models"
	"xorm.io/xorm"
)

var Engine *xorm.Engine

func InitDb() *xorm.Engine {
	user := os.Getenv("DB_USER")
	pw := os.Getenv("DB_PASSWORD")
	port := os.Getenv("DB_PORT")
	name := os.Getenv("DB_NAME")
	dsn := fmt.Sprintf("%s:%s@tcp(localhost:%s)/%s", user, pw, port, name)
	fmt.Println(dsn)
	engine, err := xorm.NewEngine("mysql", dsn)
	Engine = engine
	fmt.Println("Faggggot")
	checkErr(err, "sql.Open failed")
	if isExist, err := Engine.IsTableExist("Student"); err == nil {
		if !isExist {
			err2 := Engine.CreateTables(models.Student{})
			fmt.Println(err2)

		}
	} else {
		fmt.Println(err)
	}
	return Engine
}

func checkErr(err error, msg string) {
	if err != nil {
		log.Fatalln(msg, err)
	}
}

func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Add("Access-Control-Allow-Origin", "*")
		c.Next()
	}
}

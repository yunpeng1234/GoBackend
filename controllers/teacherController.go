package controllers

import (
	"log"
	"strconv"

	"github.com/yunpeng1234/go-api/GoBackend/models"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func getCommonStudents(c *gin.Context) {
	var user []models.User
	_, err := dbmap.Select(&user, "select * from user")

	if err == nil {
		c.JSON(200, user)
	} else {
		c.JSON(404, gin.H{"error": "user not found"})
	}
}

func registerStudent(c *gin.Context) {
	id := c.Params.ByName("id")
	var user models.User
	err := dbmap.SelectOne(&user, "SELECT * FROM user WHERE id=? LIMIT 1", id)

	if err == nil {
		user_id, _ := strconv.ParseInt(id, 0, 64)

		content := &models.User{
			Id:        user_id,
			Username:  user.Username,
			Password:  user.Password,
			Firstname: user.Firstname,
			Lastname:  user.Lastname,
		}
		c.JSON(200, content)
	} else {
		c.JSON(404, gin.H{"error": "user not found"})
	}
}

func suspendStudent(c *gin.Context) {
	var user models.User
	c.Bind(&user)
	err := dbmap.SelectOne(&user, "select * from user where Username=? LIMIT 1", user.Username)

	if err == nil {
		user_id := user.Id

		content := &models.User{
			Id:        user_id,
			Username:  user.Username,
			Password:  user.Password,
			Firstname: user.Firstname,
			Lastname:  user.Lastname,
		}
		c.JSON(200, content)
	} else {
		c.JSON(404, gin.H{"error": "user not found"})
	}

}

func retrieveNotifications(c *gin.Context) {
	var user models.User
	c.Bind(&user)

	log.Println(user)

	if user.Username != "" && user.Password != "" && user.Firstname != "" && user.Lastname != "" {

		if insert, _ := dbmap.Exec(`INSERT INTO user (Username, Password, Firstname, Lastname) VALUES (?, ?, ?, ?)`, user.Username, user.Password, user.Firstname, user.Lastname); insert != nil {
			user_id, err := insert.LastInsertId()
			if err == nil {
				content := &models.User{
					Id:        user_id,
					Username:  user.Username,
					Password:  user.Password,
					Firstname: user.Firstname,
					Lastname:  user.Lastname,
				}
				c.JSON(201, content)
			} else {
				checkErr(err, "Insert failed")
			}
		}

	} else {
		c.JSON(400, gin.H{"error": "Fields are empty"})
	}

}

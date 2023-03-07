package controllers

import (
	"fmt"
	"strings"

	"github.com/yunpeng1234/GoBackend/models"
	"xorm.io/xorm"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

// GET /api/commonstudents?teacher=teacherken%40gmail.com
func GetCommonStudents(c *gin.Context) {
	vals := c.Request.URL.Query()
	var teachers []string
	if x, found := vals["teacher"]; found {
		teachers = x
	} else {
		c.JSON(400, gin.H{"message": "Bad Request"})
	}
	var tempEngine *xorm.Session
	var students []models.Student
	fmt.Println(teachers)
	if len(teachers) >= 1 {
		tempEngine = Engine.Where("teacher=?", teachers[0])

		for i := 1; i < len(teachers); i++ {
			tempEngine = tempEngine.And("teacher=?", teachers[i])
		}
		err := tempEngine.Find(&students)
		if err != nil {
			c.JSON(404, gin.H{"error": "user not found"})
		}
	}
	fmt.Print(students)
	var commonStudents []string
	for _, value := range students {
		commonStudents = append(commonStudents, value.Email)
	}
	fmt.Print(commonStudents)

	c.JSON(200, gin.H{"students": commonStudents})
}

func RegisterStudent(c *gin.Context) {

	var reqBody models.RequestRegister
	if err := c.BindJSON(&reqBody); err != nil {
		c.JSON(400, gin.H{"message": "Bad Request"})
	}
	var students = reqBody.Students
	var teacherEmail = reqBody.Teacher
	fmt.Println(students, teacherEmail)

	session := Engine.NewSession()
	defer session.Close()

	err2 := session.Begin()
	if err2 != nil {
		c.JSON(500, gin.H{"message": "Server Unavailable"})
	}
	for _, value := range students {
		var newStudent = models.Student{}
		newStudent.Email = value
		newStudent.Teacher = teacherEmail
		newStudent.IsSuspended = false
		if _, err4 := Engine.Insert(&newStudent); err4 != nil {
		} else {
			fmt.Println(err4)
			session.Rollback()
			c.JSON(404, gin.H{"error": "Unable to register student"})
			return
		}
	}
	err3 := session.Commit()
	if err3 != nil {
		session.Rollback()
		c.JSON(500, gin.H{"message": "Server Unavailable"})
		return
	}
	c.JSON(204, "")
}

func SuspendStudent(c *gin.Context) {
	var reqBody models.RequestSuspend
	if err := c.BindJSON(&reqBody); err != nil {
		c.JSON(400, gin.H{"message": "Bad Request"})
	}
	fmt.Println(reqBody)
	var student = reqBody.Student
	var newStudent = models.Student{}
	newStudent.Email = student
	newStudent.IsSuspended = true
	_, err := Engine.ID(student).Update(newStudent)
	if err == nil {
		c.JSON(200, "")
	} else {
		c.JSON(500, gin.H{"error": "Unable to suspend student"})
	}
}

func getEmails(str string) []string {
	var res = strings.Split(str, " ")
	arr := []string{}
	for i := 1; i < len(res); i++ {
		arr = append(arr, res[i][1:])
	}
	return arr
}

func RetrieveNotifications(c *gin.Context) {
	var reqBody models.RequestNotification
	if err := c.BindJSON(&reqBody); err != nil {
		c.JSON(400, gin.H{"message": "Bad Request"})
	}
	fmt.Println(reqBody)
	var notifEmail = reqBody.Notification
	var teacher = reqBody.Teacher
	var studentEmail []string
	var students []models.Student
	err := Engine.Where("student.teacher=?", teacher).And("student.is_suspended=false").Find(&students)
	if err != nil {
		c.JSON(500, "Server unavailable")
	}
	studentEmail = getEmails(notifEmail)
	for i := 0; i < len(students); i++ {
		studentEmail = append(studentEmail, students[i].Email)
	}

	c.JSON(200, gin.H{"recipients": studentEmail})
}

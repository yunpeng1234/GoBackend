package controllers

import (
	"fmt"
	"strings"

	"github.com/yunpeng1234/GoBackend/models"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func intersect(a []models.Student, b []models.Student) []models.Student {
	set := make([]models.Student, 0)

	for _, v := range a {
		if contains(b, v) {
			set = append(set, v)
		}
	}

	return set
}

func contains(s []models.Student, stu models.Student) bool {
	for _, v := range s {
		if v.Email == stu.Email {
			return true
		}
	}

	return false
}

func removeDuplicateStr(strSlice []string) []string {
	allKeys := make(map[string]bool)
	list := []string{}
	for _, item := range strSlice {
		if _, value := allKeys[item]; !value {
			allKeys[item] = true
			list = append(list, item)
		}
	}
	return list
}

// GET /api/commonstudents?teacher=teacherken%40gmail.com
func GetCommonStudents(c *gin.Context) {
	vals := c.Request.URL.Query()
	var teachers []string
	if x, found := vals["teacher"]; found {
		teachers = x
	} else {
		c.JSON(400, gin.H{"message": "Bad Request"})
	}
	var students []models.Student
	var tempStudents []models.Student
	fmt.Println(teachers)
	if len(teachers) >= 1 {
		err := Engine.Table("Student").Where("teacher=?", teachers[0]).Find(&students)

		for i := 1; i < len(teachers); i++ {
			_ = Engine.Table("Student").Where("teacher=?", teachers[i]).Find(&tempStudents)
			students = intersect(students, tempStudents)
		}
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
		if test, err4 := Engine.Table("Student").Insert(&newStudent); err4 == nil {
			fmt.Println(test)
		} else {
			fmt.Println(test)
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
	var newStudent = models.Suspend{}
	newStudent.Email = student
	_, err := Engine.Table("Suspend").Insert(&newStudent)
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
	var temp []models.Suspend
	var tempS models.Suspend
	var nonSuspendedStudents []models.Student
	err := Engine.Table("Student").Where("teacher=?", teacher).Find(&students)
	if err != nil {
		c.JSON(500, "Server unavailable")
	}
	errT := Engine.Table("Suspend").Find(&temp)
	fmt.Println(errT)
	fmt.Println(temp)
	for _, val := range students {
		has, _ := Engine.Table("Suspend").ID(val.Email).Get(&tempS)
		if !has {
			nonSuspendedStudents = append(nonSuspendedStudents, val)
		}
	}
	fmt.Println(nonSuspendedStudents)
	studentEmail = getEmails(notifEmail)
	for i := 0; i < len(nonSuspendedStudents); i++ {
		studentEmail = append(studentEmail, nonSuspendedStudents[i].Email)
	}
	studentEmail = removeDuplicateStr(studentEmail)
	c.JSON(200, gin.H{"recipients": studentEmail})
}

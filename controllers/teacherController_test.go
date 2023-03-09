package controllers

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yunpeng1234/GoBackend/main"
)

type RegisterBody struct {
	teacher  string
	students []string
}

type SuspendBody struct {
	student string
}

type NotificationBody struct {
	teacher      string
	notification string
}

func TestRegister(t *testing.T) {
	body := RegisterBody{"a", []string{"b", "c"}}
	writer := main.MakeRequest("POST", "/api/register", body)
	assert.Equal(t, 204, writer.Code)
}

func TestRetrieveCommonStudents(t *testing.T) {
	writer := main.MakeRequest("GET", "/api/commonstudents?teacher=a", nil)

	assert.Equal(t, http.StatusOK, writer.Code)
}

func TestSuspendStudent(t *testing.T) {
	body := SuspendBody{"a"}

	writer := main.MakeRequest("POST", "/api/suspend", body)

	assert.Equal(t, http.StatusOK, writer.Code)
}

func TestGetNotification(t *testing.T) {
	body := NotificationBody{"a", "a @b @d"}

	writer := main.MakeRequest("POST", "/api/retrievenotification", body)

	assert.Equal(t, http.StatusOK, writer.Code)
}

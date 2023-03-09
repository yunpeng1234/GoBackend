package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
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
	router := SetUpRouter()
	body := RegisterBody{"a", []string{"b", "c"}}

	w := httptest.NewRecorder()
	var buf bytes.Buffer
	err := json.NewEncoder(&buf).Encode(body)
	if err != nil {
		fmt.Println(err)
	}
	req, _ := http.NewRequest("POST", "/api/register", &buf)
	router.ServeHTTP(w, req)

	assert.Equal(t, 204, w.Code)
	assert.Equal(t, "", w.Body.String())
}

func TestRetrieveCommonStudents(t *testing.T) {
	router := SetUpRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/commonstudents?teacher=a", nil)
	router.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "pong", w.Body.String())
}

func TestSuspendStudent(t *testing.T) {
	router := SetUpRouter()
	body := SuspendBody{"a"}

	w := httptest.NewRecorder()
	var buf bytes.Buffer
	err := json.NewEncoder(&buf).Encode(body)
	if err != nil {
		fmt.Println(err)
	}
	req, _ := http.NewRequest("POST", "/api/suspend", &buf)
	router.ServeHTTP(w, req)

	assert.Equal(t, 204, w.Code)
	assert.Equal(t, "", w.Body.String())
}

func TestGetNotification(t *testing.T) {
	router := SetUpRouter()
	body := NotificationBody{"a", "a @b @d"}
	w := httptest.NewRecorder()
	var buf bytes.Buffer
	err := json.NewEncoder(&buf).Encode(body)
	if err != nil {
		fmt.Println(err)
	}
	req, _ := http.NewRequest("POST", "/api/retrievenotification", &buf)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "pong", w.Body.String())
}

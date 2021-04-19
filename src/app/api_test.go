package app

import (
	"bytes"
	"github.com/gorilla/mux"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"time"
)

func executeRequest(req *http.Request, router *mux.Router) *httptest.ResponseRecorder {
	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	return rr
}

func checkResponseCode(t *testing.T, expected, actual int) {
	if expected != actual {
		t.Errorf("Expected response code %d. Got %d\n", expected, actual)
	}
}

var router = mux.NewRouter()

func TestMain(m *testing.M) {
	go HandleRequests(router, ":4000")

	code := m.Run()

	os.Exit(code)
}

func TestSortTasksEndpoint(t *testing.T) {
	jsonBody := []byte(`
{
	"tasks":[
		{
			"name":"task-1",
			"command":"touch /tmp/file1",
			"requires":[
				"task-2"
			]
		},
		{
			"name":"task-2",
			"command":"cat /tmp/file1"
		}
	]
}
`)

	// Test fails without sleep, not sure why
	time.Sleep(time.Millisecond)
	req, _ := http.NewRequest("POST", "/sortTasks", bytes.NewBuffer(jsonBody))
	response := executeRequest(req, router)

	checkResponseCode(t, http.StatusOK, response.Code)

	expected := `[
  {
    "Name": "task-2",
    "Command": "cat /tmp/file1"
  },
  {
    "Name": "task-1",
    "Command": "touch /tmp/file1"
  }
]`
	if body := response.Body.String(); body != expected {
		t.Errorf("Expected %s Got %s", expected, body)
	}
}

func TestBashEndpoint(t *testing.T) {
	jsonBody := []byte(`
{
	"tasks":[
		{
			"name":"task-1",
			"command":"touch /tmp/file1",
			"requires":[
				"task-2"
			]
		},
		{
			"name":"task-2",
			"command":"cat /tmp/file1"
		}
	]
}
`)

	req, _ := http.NewRequest("POST", "/bash", bytes.NewBuffer(jsonBody))
	response := executeRequest(req, router)

	checkResponseCode(t, http.StatusOK, response.Code)

	expected := "#!/usr/bin/env bash \n\ncat /tmp/file1\ntouch /tmp/file1\n"
	if body := response.Body.String(); body != expected {
		t.Errorf("Expected %s Got %s", expected, body)
	}
}

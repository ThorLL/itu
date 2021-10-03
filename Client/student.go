package main

import (
	"bytes"
	"encoding/json"
	"net/http"
)

type student struct {
	ID       string   `json:"id" binding:"required"`
	Name     string   `json:"name"`
	Workload int      `json:"workload"`
	Courses  []course `json:"courses"`
}

func updateStudent(student student) {
	URL := URLBASE + "students/"

	jsonValue, _ := json.Marshal(student)
	req, _ := http.NewRequest("PUT", URL, bytes.NewBuffer(jsonValue))
	doRequest(req)
}

func deleteStudent(ID string) {
	URL := URLBASE + "students/" + ID
	req, _ := http.NewRequest("DELETE", URL, nil)
	doRequest(req)
}

func createStudent(student student) {
	URL := URLBASE + "students/"

	jsonValue, _ := json.Marshal(student)
	req, _ := http.NewRequest("POST", URL, bytes.NewBuffer(jsonValue))
	doRequest(req)

}

func getStudentByID(ID string) student {
	URL := URLBASE + "students/" + ID

	req, _ := http.NewRequest("GET", URL, nil)

	var responseObject student

	json.Unmarshal(getBodyBytes(req), &responseObject)

	return responseObject
}

func getStudents() []student {
	URL := URLBASE + "students/"

	req, _ := http.NewRequest("GET", URL, nil)

	var responseObject []student

	json.Unmarshal(getBodyBytes(req), &responseObject)

	return responseObject
}

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
	request, _ := http.NewRequest("PUT", URL, bytes.NewBuffer(jsonValue))
	doRequest(request)
}

func deleteStudent(ID string) {
	URL := URLBASE + "students/" + ID
	request, _ := http.NewRequest("DELETE", URL, nil)
	doRequest(request)
}

func createStudent(student student) {
	URL := URLBASE + "students/"

	jsonValue, _ := json.Marshal(student)
	request, _ := http.NewRequest("POST", URL, bytes.NewBuffer(jsonValue))
	doRequest(request)
}

func getStudentByID(ID string) student {
	URL := URLBASE + "students/" + ID

	request, _ := http.NewRequest("GET", URL, nil)
	var responseObject student
	json.Unmarshal(getBodyBytes(request), &responseObject)
	return responseObject
}

func getStudents() []student {
	URL := URLBASE + "students/"

	request, _ := http.NewRequest("GET", URL, nil)

	var responseObject []student

	json.Unmarshal(getBodyBytes(request), &responseObject)

	return responseObject
}

package main

import (
	"bytes"
	"encoding/json"
	"net/http"
)

type teacher struct {
	Name    string   `json:"name" binding:"required"`
	Rating  int      `json:"rating"`
	Courses []course `json:"courses"`
}

func updateTeacher(teacher teacher) {
	URL := URLBASE + "teachers/"

	jsonValue, _ := json.Marshal(teacher)
	req, _ := http.NewRequest("PUT", URL, bytes.NewBuffer(jsonValue))
	doRequest(req)
}

func deleteTeacher(ID string) {
	URL := URLBASE + "teachers/" + ID
	req, _ := http.NewRequest("DELETE", URL, nil)
	doRequest(req)
}

func createTeacher(teacher teacher) {
	URL := URLBASE + "teachers/"

	jsonValue, _ := json.Marshal(teacher)
	req, _ := http.NewRequest("POST", URL, bytes.NewBuffer(jsonValue))
	doRequest(req)

}

func getTeacherByName(name string) teacher {
	URL := URLBASE + "teachers/" + name

	req, _ := http.NewRequest("GET", URL, nil)

	var responseObject teacher

	json.Unmarshal(getBodyBytes(req), &responseObject)

	return responseObject
}

func getTeachers() []teacher {
	URL := URLBASE + "teachers/"

	req, _ := http.NewRequest("GET", URL, nil)

	var responseObject []teacher

	json.Unmarshal(getBodyBytes(req), &responseObject)

	return responseObject
}

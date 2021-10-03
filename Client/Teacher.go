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
	request, _ := http.NewRequest("PUT", URL, bytes.NewBuffer(jsonValue))
	doRequest(request)
}

func deleteTeacher(ID string) {
	URL := URLBASE + "teachers/" + ID
	request, _ := http.NewRequest("DELETE", URL, nil)
	doRequest(request)
}

func createTeacher(teacher teacher) {
	URL := URLBASE + "teachers/"

	jsonValue, _ := json.Marshal(teacher)
	request, _ := http.NewRequest("POST", URL, bytes.NewBuffer(jsonValue))
	doRequest(request)

}

func getTeacherByName(name string) teacher {
	URL := URLBASE + "teachers/" + name

	request, _ := http.NewRequest("GET", URL, nil)

	var responseObject teacher

	json.Unmarshal(getBodyBytes(request), &responseObject)

	return responseObject
}

func getTeachers() []teacher {
	URL := URLBASE + "teachers/"

	request, _ := http.NewRequest("GET", URL, nil)

	var responseObject []teacher

	json.Unmarshal(getBodyBytes(request), &responseObject)

	return responseObject
}

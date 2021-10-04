package main

import (
	"bytes"
	"encoding/json"
	"net/http"
)

type course struct {
	ID     string
	Name   string
	Rating int
}

func updateCourse(course course) {
	URL := URLBASE + "courses/"

	jsonValue, _ := json.Marshal(course)
	request, _ := http.NewRequest("PUT", URL, bytes.NewBuffer(jsonValue))
	doRequest(request)
}

func deleteCourse(ID string) {
	URL := URLBASE + "courses/" + ID
	request, _ := http.NewRequest("DELETE", URL, nil)
	doRequest(request)
}

func createCourse(course course) {
	URL := URLBASE + "courses/"

	jsonValue, _ := json.Marshal(course)
	request, _ := http.NewRequest("POST", URL, bytes.NewBuffer(jsonValue))
	doRequest(request)

}

func getCourseByID(id string) course {
	URL := URLBASE + "courses/" + id

	request, _ := http.NewRequest("GET", URL, nil)

	var responseObject course

	json.Unmarshal(getBodyBytes(request), &responseObject)

	return responseObject
}

func getCourses() []course {
	URL := URLBASE + "courses/"

	request, _ := http.NewRequest("GET", URL, nil)

	var responseObject []course

	json.Unmarshal(getBodyBytes(request), &responseObject)

	return responseObject
}

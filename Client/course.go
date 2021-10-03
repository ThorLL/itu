package main

import (
	"bytes"
	"encoding/json"
	"net/http"
)

type course struct {
	ID     string `json:"id" binding:"required"`
	Name   string `json:"name"`
	Rating int    `json:"rating"`
}

func updateCourse(course course) {
	URL := URLBASE + "courses/"

	jsonValue, _ := json.Marshal(course)
	req, _ := http.NewRequest("PUT", URL, bytes.NewBuffer(jsonValue))
	doRequest(req)
}

func deleteCourse(ID string) {
	URL := URLBASE + "courses/" + ID
	req, _ := http.NewRequest("DELETE", URL, nil)
	doRequest(req)
}

func createCourse(course course) {
	URL := URLBASE + "courses/"

	jsonValue, _ := json.Marshal(course)
	req, _ := http.NewRequest("POST", URL, bytes.NewBuffer(jsonValue))
	doRequest(req)

}

func getCourseByID(id string) course {
	URL := URLBASE + "courses/" + id

	req, _ := http.NewRequest("GET", URL, nil)

	var responseObject course

	json.Unmarshal(getBodyBytes(req), &responseObject)

	return responseObject
}

func getCourses() []course {
	URL := URLBASE + "courses/"

	req, _ := http.NewRequest("GET", URL, nil)

	var responseObject []course

	json.Unmarshal(getBodyBytes(req), &responseObject)

	return responseObject
}

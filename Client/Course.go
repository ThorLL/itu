package main

import (
	ITU "./ITU"
	"bytes"
	"context"
	"encoding/json"
	"google.golang.org/grpc"
	"log"
	"net/http"
	"time"
)

type course struct {
	ID     string `json:"id" binding:"required"`
	Name   string `json:"name"`
	Rating int    `json:"rating"`
}

func updateCourse(course course) {
	URL := URLBASE + "courses/"

	conn, _ := grpc.Dial(URL)
	c := ITU.NewCourseClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	_, err := c.updateCourse(ctx, course)
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
}

/*
func updateCourse(course course) {
	URL := URLBASE + "courses/"

	jsonValue, _ := json.Marshal(course)
	request, _ := http.NewRequest("PUT", URL, bytes.NewBuffer(jsonValue))
	doRequest(request)
}
*/

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

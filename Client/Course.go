package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/ThorLL/itu"
	"google.golang.org/grpc"
	"net/http"
	"time"
)

type course struct {
	ID     string
	Name   string
	Rating int32
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

/*
func createCourse(course course) {
	URL := URLBASE + "courses/"

	jsonValue, _ := json.Marshal(course)
	request, _ := http.NewRequest("POST", URL, bytes.NewBuffer(jsonValue))
	doRequest(request)
}


*/
const address = "localhost:50051"

func CreateCourse(course course) {
	//URL := URLBASE + "courses/"

	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		fmt.Printf("did not connect: %v\n", err)
	}
	defer conn.Close()
	c := itu.NewCourseMethodsClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	messageCourse := itu.Course{
		ID:     &course.ID,
		Name:   &course.Name,
		Rating: &course.Rating,
	}
	_, err = c.CreateCourse(ctx, &messageCourse)
	if err != nil {
		fmt.Println(err)
	}
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

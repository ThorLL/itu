package main

import (
	"context"
	"fmt"
	"github.com/ThorLL/itu"
	"github.com/gin-gonic/gin"
	"net/http"
	//"net/http"
)

type course struct {
	ID     string `json:"id" binding:"required"`
	Name   string `json:"name"`
	Rating int32  `json:"rating"`
}

var courses = []course{
	{ID: "1", Name: "BDSA", Rating: 5},
	{ID: "2", Name: "IDBD", Rating: 7},
	{ID: "3", Name: "DISYS", Rating: 10},
	{ID: "4", Name: "GRPRO", Rating: 2},
	{ID: "5", Name: "DMAT", Rating: 4},
	{ID: "6", Name: "ALGO", Rating: 8},
}

func (s *server) createCourse(ctx context.Context, in *itu.Course) itu.Empty {
	fmt.Printf("Received: %v\n", in.GetID())
	var newCourse = course{
		ID:     in.GetID(),
		Name:   in.GetName(),
		Rating: in.GetRating(),
	}
	courses = append(courses, newCourse)
	return itu.Empty{}
}

func getCourses(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, courses)
}

func getCourseByID(c *gin.Context) {
	id := c.Param("id")

	for _, course := range courses {
		if course.ID == id {
			c.IndentedJSON(http.StatusOK, course)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "course not found"})
}

func deleteCourseByID(c *gin.Context) {
	id := c.Param("id")

	for i := 0; i < len(courses); i++ {
		if courses[i].ID == id {
			courses = removeCourses(courses, i)
			removeCourseForStudents(students, id)
			removeCourseForTeachers(teachers, id)
			c.IndentedJSON(http.StatusOK, gin.H{"message": "course deleted"})
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "course not found"})
}

func removeCourses(courses []course, index int) []course {
	return append(courses[:index], courses[index+1:]...)
}

func removeCourse(c []course, id string) []course {
	for i := 0; i < len(c); i++ {
		if c[i].ID == id {
			return removeCourses(c, i)
		}
	}
	return c
}

func removeCourseForStudents(s []student, id string) {
	for i := 0; i < len(s); i++ {
		s[i].Courses = removeCourse(s[i].Courses, id)
	}
}

func removeCourseForTeachers(t []teacher, id string) {
	for i := 0; i < len(t); i++ {
		t[i].Courses = removeCourse(t[i].Courses, id)
	}
}

func updateCourse(c *gin.Context) {
	var course course

	if err := c.BindJSON(&course); err != nil {
		return
	}

	for i := 0; i < len(courses); i++ {
		if course.ID == courses[i].ID {
			courses[i] = course
			c.IndentedJSON(http.StatusNotFound, gin.H{"message": "course updated"})
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "course not found"})
}

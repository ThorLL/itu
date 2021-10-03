package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type course struct {
	ID     string `json:"id" binding:"required"`
	Name   string `json:"name"`
	Rating int    `json:"rating"`
}

var courses = []course{
	{ID: "1", Name: "BDSA", Rating: 5},
	{ID: "2", Name: "IDBD", Rating: 7},
	{ID: "3", Name: "DISYS", Rating: 10},
	{ID: "4", Name: "GRPRO", Rating: 2},
	{ID: "5", Name: "DMAT", Rating: 4},
	{ID: "6", Name: "ALGO", Rating: 8},
}

func getCourses(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, courses)
}

func postCourses(c *gin.Context) {
	var newCourse course

	if err := c.BindJSON(&newCourse); err != nil {
		return
	}

	courses = append(courses, newCourse)
	c.IndentedJSON(http.StatusCreated, newCourse)
}

func getCourseByID(c *gin.Context) {
	id := c.Param("id")

	for _, s := range courses {
		if s.ID == id {
			c.IndentedJSON(http.StatusOK, s)
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

func removeCourses(s []course, index int) []course {
	return append(s[:index], s[index+1:]...)
}

func removeCourse(s []course, id string) []course {
	for i := 0; i < len(s); i++ {
		if s[i].ID == id {
			return removeCourses(s, i)
		}
	}
	return s
}

func removeCourseForStudents(s []student, id string) {
	for i := 0; i < len(s); i++ {
		s[i].Courses = removeCourse(s[i].Courses, id)
	}
}

func removeCourseForTeachers(s []teacher, id string) {
	for i := 0; i < len(s); i++ {
		s[i].Courses = removeCourse(s[i].Courses, id)
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

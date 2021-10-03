package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type teacher struct {
	Name    string   `json:"name" binding:"required"`
	Rating  int      `json:"rating"`
	Courses []course `json:"courses"`
}

var teachers = []teacher{
	{Name: "Peter", Rating: 9, Courses: []course{courses[0], courses[1], courses[2]}},
	{Name: "Flemming", Rating: 4, Courses: []course{courses[2], courses[3], courses[4]}},
	{Name: "Bob", Rating: 6, Courses: []course{courses[5], courses[2], courses[1]}},
}

func getTeachers(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, teachers)
}

func postTeachers(c *gin.Context) {
	var newTeacher teacher

	if err := c.BindJSON(&newTeacher); err != nil {
		return
	}

	teachers = append(teachers, newTeacher)
	c.IndentedJSON(http.StatusCreated, newTeacher)
}

func getTeacherByName(c *gin.Context) {
	name := c.Param("name")

	for _, s := range teachers {
		if s.Name == name {
			c.IndentedJSON(http.StatusOK, s)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "teacher not found"})
}

func deleteTeacherByName(c *gin.Context) {
	name := c.Param("name")

	for i := 0; i < len(teachers); i++ {
		if teachers[i].Name == name {
			teachers = removeTeacher(teachers, i)
			c.IndentedJSON(http.StatusOK, gin.H{"message": "teacher deleted"})
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "teacher not found"})
}

func removeTeacher(s []teacher, index int) []teacher {
	return append(s[:index], s[index+1:]...)
}

func updateTeacher(c *gin.Context) {
	var teacher teacher

	if err := c.BindJSON(&teacher); err != nil {
		return
	}

	for i := 0; i < len(teachers); i++ {
		if teacher.Name == teachers[i].Name {
			teachers[i] = teacher
			c.IndentedJSON(http.StatusNotFound, gin.H{"message": "teacher updated"})
			return
		}
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "teacher not found"})
}

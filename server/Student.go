package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type student struct {
	ID       string   `json:"id" binding:"required"`
	Name     string   `json:"name"`
	Workload int      `json:"workload"`
	Courses  []course `json:"courses"`
}

var students = []student{
	{ID: "1", Name: "My", Workload: 5, Courses: []course{courses[0], courses[1], courses[2]}},
	{ID: "2", Name: "Alex", Workload: 10, Courses: []course{courses[2], courses[3], courses[4]}},
	{ID: "3", Name: "Rasmus", Workload: 7, Courses: []course{courses[5], courses[2], courses[1]}},
}

func getStudents(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, students)
}

func postStudents(c *gin.Context) {
	var newStudent student

	if err := c.BindJSON(&newStudent); err != nil {
		return
	}

	students = append(students, newStudent)
	c.IndentedJSON(http.StatusCreated, newStudent)
}

func getStudentByID(c *gin.Context) {
	id := c.Param("id")

	for _, s := range students {
		if s.ID == id {
			c.IndentedJSON(http.StatusOK, s)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "student not found"})
}

func deleteStudentByID(c *gin.Context) {
	id := c.Param("id")
	for i := 0; i < len(students); i++ {
		if students[i].ID == id {
			students = removeStudent(students, i)
			c.IndentedJSON(http.StatusOK, gin.H{"message": "student deleted"})
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "student not found"})
}

func removeStudent(s []student, index int) []student {
	return append(s[:index], s[index+1:]...)
}

func updateStudent(c *gin.Context) {
	var student student

	if err := c.BindJSON(&student); err != nil {
		return
	}

	for i := 0; i < len(students); i++ {
		if student.ID == students[i].ID {
			students[i] = student
			c.IndentedJSON(http.StatusOK, gin.H{"message": "student updated"})
			return
		}
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "student not found"})
}

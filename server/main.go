package main

import "github.com/gin-gonic/gin"

func main() {
	router := gin.Default()
	router.GET("/students", getStudents)
	router.GET("/students/:id", getStudentByID)
	router.POST("/students", postStudents)
	router.DELETE("/students/:id", deleteStudentByID)
	router.PUT("/students/", updateStudent)

	router.GET("/courses", getCourses)
	router.GET("/courses/:id", getCourseByID)
	router.POST("/courses", postCourses)
	router.DELETE("/courses/:id", deleteCourseByID)
	router.PUT("/courses/", updateCourse)

	router.GET("/teachers", getTeachers)
	router.GET("/teachers/:name", getTeacherByName)
	router.POST("/teachers", postTeachers)
	router.DELETE("/teachers/:name", deleteTeacherByName)
	router.PUT("/teachers/", updateTeacher)

	router.Run("localhost:3000")
}

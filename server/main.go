package main

import (
	"fmt"
	"github.com/ThorLL/itu"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"net"
)

type server struct {
	itu.UnimplementedCourseMethodsServer
}

func main() {
	go startRest("3000")
	go startGRPC("50051")
}

func startRest(port string) {
	router := gin.Default()
	router.GET("/students", getStudents)
	router.GET("/students/:id", getStudentByID)
	router.POST("/students", postStudents)
	router.DELETE("/students/:id", deleteStudentByID)
	router.PUT("/students/", updateStudent)

	router.GET("/courses", getCourses)
	router.GET("/courses/:id", getCourseByID)
	router.DELETE("/courses/:id", deleteCourseByID)
	router.PUT("/courses/", updateCourse)

	router.GET("/teachers", getTeachers)
	router.GET("/teachers/:name", getTeacherByName)
	router.POST("/teachers", postTeachers)
	router.DELETE("/teachers/:name", deleteTeacherByName)
	router.PUT("/teachers/", updateTeacher)

	router.Run("localhost:" + port)
}

func startGRPC(port string) {
	lis, err := net.Listen("tcp", ":"+port)
	fmt.Printf("et eller andet")
	if err != nil {
		fmt.Printf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	itu.RegisterCourseMethodsServer(s, &server{})
	fmt.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		fmt.Printf("failed to serve: %v", err)
	}
}

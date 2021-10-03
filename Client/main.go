package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const URLBASE = "http://localhost:8080/"

func main() {
	testStudentRequests()
	testTeacherRequests()
	testCoursesRequests()
}

func testStudentRequests() {
	fmt.Printf("%+v\n", getStudentByID("1").Name)
	fmt.Printf("%+v\n", getStudents())
	student1 := student{
		ID:       "10",
		Name:     "My",
		Workload: 10,
		Courses:  nil,
	}
	student2 := student{
		ID:       "2",
		Name:     "Bob",
		Workload: 10,
		Courses:  nil,
	}
	createStudent(student1)
	fmt.Printf("%+v\n", getStudents())
	deleteStudent("1")
	fmt.Printf("%+v\n", getStudents())
	updateStudent(student2)
	fmt.Printf("%+v\n", getStudents())
}

func testTeacherRequests() {
	fmt.Printf("%+v\n", getTeacherByName("Peter").Name)
	fmt.Printf("%+v\n", getTeachers())
	teacher1 := teacher{
		Name:   "Lars",
		Rating: 10,
		Courses: []course{{
			ID:     "10",
			Name:   "Course10",
			Rating: 10,
		}, {
			ID:     "11",
			Name:   "Course11",
			Rating: 5,
		},
		},
	}
	teacher2 := teacher{
		Name:   "Bob",
		Rating: 10,
		Courses: []course{{
			ID:     "12",
			Name:   "Course12",
			Rating: 2,
		}, {
			ID:     "11",
			Name:   "Course11",
			Rating: 5,
		},
		},
	}
	createTeacher(teacher1)
	fmt.Printf("%+v\n", getTeachers())
	deleteTeacher("Peter")
	fmt.Printf("%+v\n", getTeachers())
	updateTeacher(teacher2)
	fmt.Printf("%+v\n", getTeachers())
}

func testCoursesRequests() {
	fmt.Printf("%+v\n", getCourseByID("1").Name)
	fmt.Printf("%+v\n", getCourses())
	course1 := course{
		ID:     "7",
		Name:   "Course 7",
		Rating: 3,
	}
	course2 := course{
		ID:     "2",
		Name:   "Course 2",
		Rating: 10,
	}
	createCourse(course1)
	fmt.Printf("%+v\n", getCourses())
	deleteCourse("1")
	fmt.Printf("%+v\n", getCourses())
	updateCourse(course2)
	fmt.Printf("%+v\n", getCourses())
}

func getBodyBytes(req *http.Request) []byte {
	client := &http.Client{}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	resp, _ := client.Do(req)
	defer resp.Body.Close()
	bodyBytes, _ := ioutil.ReadAll(resp.Body)
	return bodyBytes
}

func doRequest(req *http.Request) {
	client := &http.Client{}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	resp, _ := client.Do(req)
	defer resp.Body.Close()
}

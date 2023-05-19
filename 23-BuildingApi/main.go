package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

type Course struct {
	CourseId    string  `json:"courseid"`
	CourseName  string  `json:"coursename"`
	CoursePrice int64   `json:"price"`
	Author      *Author `json:"author"`
}

type Author struct {
	FullName string `json:"fullname"`
	Website  string `json:"website"`
}

var courses []Course

func (c *Course) isEmpty() bool {
	return c.CourseName == ""
}

func main() {

	addData()
	router := mux.NewRouter()

	router.HandleFunc("/", HomeRouter).Methods("GET")
	router.HandleFunc("/courses", getAllCourses).Methods("GET")
	router.HandleFunc("/course/{id}", getCourse).Methods("GET")
	router.HandleFunc("/add", addCourse).Methods("POST")
	router.HandleFunc("/update/{id}", updateCourse).Methods("PUT")
	router.HandleFunc("/delete/{id}", deleteCourse).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000", router))
}

func addData() {
	courses = append(courses, Course{CourseId: "1", CourseName: "Golang Bootcamp", CoursePrice: 200,
		Author: &Author{FullName: "Nikhil Rana", Website: "udemy.com"}})

	courses = append(courses, Course{CourseId: "2", CourseName: "Python Bootcamp", CoursePrice: 900,
		Author: &Author{FullName: "Ajay Rana", Website: "coursera.com"}})

	courses = append(courses, Course{CourseId: "3", CourseName: "React.js Bootcamp", CoursePrice: 740,
		Author: &Author{FullName: "Raman Rana", Website: "coursera.com"}})

	courses = append(courses, Course{CourseId: "4", CourseName: "Next.js Bootcamp", CoursePrice: 940,
		Author: &Author{FullName: "Sikha Rana", Website: "scaler.com"}})

}

func HomeRouter(w http.ResponseWriter, r *http.Request) {
	fmt.Println("This is home route")
	w.Write([]byte("<h1>Hello This is first api<h1>"))
}

func getAllCourses(w http.ResponseWriter, r *http.Request) {

	fmt.Println("Get all courses")

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(courses)
}

func getCourse(w http.ResponseWriter, r *http.Request) {

	fmt.Println("Get a course with courseId")

	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	for _, course := range courses {
		if course.CourseId == params["id"] {
			json.NewEncoder(w).Encode(course)
			return
		}
	}

	json.NewEncoder(w).Encode("No course found with given id")
	return
}

func addCourse(w http.ResponseWriter, r *http.Request) {

	fmt.Println("Adding a course")

	w.Header().Set("Content-Type", "application/json")
	if r.Body == nil {
		json.NewEncoder(w).Encode("Please provide the course")
		return
	}

	var course Course
	err := json.NewDecoder(r.Body).Decode(&course)

	if err != nil {
		panic(err)
	}

	if course.isEmpty() {
		json.NewEncoder(w).Encode("Please provide the corrent content data")
		return
	}

	rand.Seed(time.Now().UnixNano())
	course.CourseId = strconv.Itoa(rand.Intn(1000))
	courses = append(courses, course)
	json.NewEncoder(w).Encode(course)
	return

}

func updateCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Update a course")
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	for index, course := range courses {
		if course.CourseId == params["id"] {
			courses = append(courses[:index], courses[index+1:]...)
			var course Course
			err := json.NewDecoder(r.Body).Decode(&course)
			if err != nil {
				panic(err)
			}
			course.CourseId = params["id"]
			courses = append(courses, course)
			json.NewEncoder(w).Encode(course)
			return
		}
	}

	json.NewEncoder(w).Encode("data not found")
	return

}

func deleteCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Deleting a course")
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	for index, course := range courses {
		if course.CourseId == params["id"] {
			courses = append(courses[:index], courses[index+1:]...)
			break
		}
	}

	json.NewEncoder(w).Encode("Course deleted successfully")
	return
}

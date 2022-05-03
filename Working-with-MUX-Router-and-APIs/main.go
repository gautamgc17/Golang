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

// Model for course - file
type Course struct {
	CourseId    string  `json:"courseid"`
	CourseName  string  `json:"coursename"`
	CoursePrice int     `json:"price"`
	Author      *Author `json:"author"`
}

type Author struct {
	Fullname string `json:"fullname"`
	Website  string `json:"website"`
}

// injecting fake values into DB - seeding
var courses []Course

// middleware, helper - file
func (c *Course) IsEmpty() bool {
	// return c.CourseId == "" && c.CourseName == ""
	return c.CourseName == ""
}

func main() {
	fmt.Println("API Tutorial")

	// seeding
	courses = append(courses, Course{CourseId: "1", CourseName: "ReactJS", CoursePrice: 499,
		Author: &Author{Fullname: "Gauti", Website: "dev.in"}})
	courses = append(courses, Course{CourseId: "2", CourseName: "JS", CoursePrice: 299,
		Author: &Author{Fullname: "Samarth", Website: "jinja.co"}})
		
	// routing 
	r := mux.NewRouter()

	r.HandleFunc("/" , serveHome).Methods("GET")
	r.HandleFunc("/courses" , getAllCourses).Methods("GET")
	r.HandleFunc("/course/{id}", getOneCourse).Methods("GET")
	r.HandleFunc("/course", createOneCourse).Methods("POST")
	r.HandleFunc("/course/{id}", updateOneCourse).Methods("PUT")
	r.HandleFunc("/course/{id}", deleteOneCourse).Methods("DELETE")

	// listening to a port
	if err := http.ListenAndServe(":4000", r); err != nil {
		log.Fatal(err)
	}
}

// controllers - file

// server home route
func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1> Welcome to building API tutorial </h1>"))
}

// sending a API json response for all courses
func getAllCourses(w http.ResponseWriter, r *http.Request) {
	// Header is the metadata associated with API
	fmt.Println("Get All Courses")
	w.Header().Set("Content-Type", "application/json")
	// NewEncoder returns a new encoder that writes to w
	// Writes the JSON encoding of v to the stream, followed by a newline character
	json.NewEncoder(w).Encode(courses)
}

// get one course based on request id
func getOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get One Course")
	w.Header().Set("Content-Type", "application/json")

	// all key-value pairs, could use url package here to extract parameters
	params := mux.Vars(r)
	fmt.Println(params)
	fmt.Printf("Type is: %T", params)

	// loop through courses, find matching id and return the response
	for _, course := range courses {
		if course.CourseId == params["id"] {
			json.NewEncoder(w).Encode(course)
			return
		}
	}
	json.NewEncoder(w).Encode("No course found with given ID")
}

// add a course controller
func createOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Create One Course")
	w.Header().Set("Content-Type", "application/json")

	// what if body is empty
	if r.Body == nil {
		json.NewEncoder(w).Encode("Please send some data")
	}

	var course Course
	// new decoder that reads from r, Decode reads the next JSON-encoded value
	_ = json.NewDecoder(r.Body).Decode(&course)
	if course.IsEmpty() {
		json.NewEncoder(w).Encode("No data inside JSON")
		return
	}

	// generate unique string id and append course into courses
	rand.Seed(time.Now().UnixNano())
	course.CourseId = strconv.Itoa(rand.Intn(100))
	courses = append(courses, course)
	json.NewEncoder(w).Encode(course)
}

// update a course controller
func updateOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Create One Course")
	w.Header().Set("Content-Type", "applications/json")

	// first grab id from request
	params := mux.Vars(r)

	for index, course := range courses {
		if course.CourseId == params["id"] {
			courses = append(courses[:index], courses[index+1:]...)
			var course Course
			json.NewDecoder(r.Body).Decode(&course)
			course.CourseId = params["id"]
			courses = append(courses, course)
			json.NewEncoder(w).Encode(course)
			return
		}
	}
	json.NewEncoder(w).Encode("Course to be updated with given ID not found")
}

// delete a course controller
func deleteOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Delete One Course")
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	for index, course := range courses {
		if course.CourseId == params["id"] {
			courses = append(courses[:index], courses[index+1:]...)
			json.NewEncoder(w).Encode("Deleted course with given ID")
			break
		}
	}
}

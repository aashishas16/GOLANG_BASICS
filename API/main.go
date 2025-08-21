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

// Course struct represents a course
type Course struct {
	CourseId    string  `json:"course_id"`
	CourseName  string  `json:"course_name"`
	CoursePrice int     `json:"course_price"`
	Author      *Author `json:"author"` // Nested struct pointer
}

// Author struct represents the author of a course
type Author struct {
	Fullname string `json:"fullname"`
	Website  string `json:"website"`
}

// Global variable acting as a "fake DB"
var courses []Course

// Helper method to check if a course struct is empty
func (c *Course) IsEmpty() bool {
	return c.CourseName == ""
}

func main() {
	fmt.Println("API is running on port 4000")
	r := mux.NewRouter()

	// Seeding the fake database
	courses = append(courses, Course{
		CourseId:    "1",
		CourseName:  "React JS",
		CoursePrice: 299,
		Author: &Author{
			Fullname: "Aashish Singune",
			Website:  "go.dev",
		},
	})
	courses = append(courses, Course{
		CourseId:    "2",
		CourseName:  "Django",
		CoursePrice: 199,
		Author: &Author{
			Fullname: "Aashish Singune",
			Website:  "go.dev",
		},
	})

	// API routes
	r.HandleFunc("/", serveHome).Methods("GET")
	r.HandleFunc("/courses", getAllCourses).Methods("GET")
	r.HandleFunc("/courses/{course_id}", getOneCourse).Methods("GET")
	r.HandleFunc("/courses", createOneCourse).Methods("POST")
	r.HandleFunc("/courses/{course_id}", updateOneCourse).Methods("PUT")
	r.HandleFunc("/courses/{course_id}", deleteOneCourse).Methods("DELETE")

	// Listen and serve on port 4000
	log.Fatal(http.ListenAndServe(":4000", r))
}

// Serve home route handler
func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Welcome to API</h1>"))
}

// Get all courses handler
func getAllCourses(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get all courses")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(courses)
}

// Get one course by ID handler
func getOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get one course")
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	for _, course := range courses {
		if course.CourseId == params["course_id"] {
			json.NewEncoder(w).Encode(course)
			return
		}
	}
	// Return a JSON object for consistency
	json.NewEncoder(w).Encode(map[string]string{"message": "No course found with the given ID"})
}

// Create one course handler
func createOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Create one course")
	w.Header().Set("Content-Type", "application/json")

	if r.Body == nil {
		json.NewEncoder(w).Encode(map[string]string{"message": "Please send some data"})
		return
	}

	var course Course
	err := json.NewDecoder(r.Body).Decode(&course)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if course.IsEmpty() {
		json.NewEncoder(w).Encode(map[string]string{"message": "No data inside the request body"})
		return
	}

	// Generate a unique ID
	rand.Seed(time.Now().UnixNano())
	course.CourseId = strconv.Itoa(rand.Intn(100))
	courses = append(courses, course)
	json.NewEncoder(w).Encode(course)
}

// Update one course by ID handler
func updateOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Update one course")
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	// Decode the new data from the request body
	var updatedCourse Course
	err := json.NewDecoder(r.Body).Decode(&updatedCourse)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	for index, course := range courses {
		if course.CourseId == params["course_id"] {
			// Update the fields of the found course
			courses[index].CourseName = updatedCourse.CourseName
			courses[index].CoursePrice = updatedCourse.CoursePrice
			courses[index].Author = updatedCourse.Author

			json.NewEncoder(w).Encode(courses[index])
			return
		}
	}

	// If no course found, respond with an error message
	json.NewEncoder(w).Encode(map[string]string{"message": "No course found with the given ID"})
}

// Delete one course by ID handler
func deleteOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Delete one course")
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	for index, course := range courses {
		if course.CourseId == params["course_id"] {
			// Remove the course from the slice
			courses = append(courses[:index], courses[index+1:]...)
			json.NewEncoder(w).Encode(map[string]string{"message": "Course deleted successfully"})
			return
		}
	}

	// If no course found, respond with an error message
	json.NewEncoder(w).Encode(map[string]string{"message": "No course found with the given ID"})
}

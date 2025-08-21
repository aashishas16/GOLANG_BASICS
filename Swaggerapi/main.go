package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	_ "GOLANG_BASICS/docs" // Import the generated docs using your module name

	"github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger/v2"
)

// Course struct represents a course
type Course struct {
	CourseId    string  `json:"course_id"`
	CourseName  string  `json:"course_name"`
	CoursePrice int     `json:"course_price"`
	Author      *Author `json:"author"`
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

// @title Go API with Swagger
// @version 1.0
// @description This is a sample API for a course management system.
// @host localhost:4000
// @BasePath /
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
	courses = append(courses, Course{
		CourseId:    "3",
		CourseName:  "Go Backend",
		CoursePrice: 350,
		Author: &Author{
			Fullname: "Jane Doe",
			Website:  "jdoe.dev",
		},
	})
	courses = append(courses, Course{
		CourseId:    "4",
		CourseName:  "Advanced Python",
		CoursePrice: 450,
		Author: &Author{
			Fullname: "John Smith",
			Website:  "jsmith.com",
		},
	})
	courses = append(courses, Course{
		CourseId:    "5",
		CourseName:  "Cloud Essentials",
		CoursePrice: 150,
		Author: &Author{
			Fullname: "Mary Williams",
			Website:  "cloudy.org",
		},
	})

	// API routes
	r.HandleFunc("/", serveHome).Methods("GET")
	r.HandleFunc("/courses", getAllCourses).Methods("GET")
	r.HandleFunc("/courses/{course_id}", getOneCourse).Methods("GET")
	r.HandleFunc("/courses", createOneCourse).Methods("POST")
	r.HandleFunc("/courses/{course_id}", updateOneCourse).Methods("PUT")
	r.HandleFunc("/courses/{course_id}", deleteOneCourse).Methods("DELETE")

	// Swagger route
	r.PathPrefix("/swagger").Handler(httpSwagger.WrapHandler)

	// Listen and serve on port 4000
	log.Fatal(http.ListenAndServe(":4000", r))
}

// Serve home route handler
func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Welcome to API with Swagger</h1>"))
}

// @Summary Get all courses
// @Description Retrieves a list of all courses
// @Produce json
// @Success 200 {array} Course
// @Router /courses [get]
func getAllCourses(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(courses)
}

// @Summary Get a single course by ID
// @Description Get details for a specific course by its unique ID
// @ID get-one-course
// @Produce json
// @Param course_id path string true "Course ID"
// @Success 200 {object} Course
// @Failure 404 {object} map[string]string "No course found with the given ID"
// @Router /courses/{course_id} [get]
func getOneCourse(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, course := range courses {
		if course.CourseId == params["course_id"] {
			json.NewEncoder(w).Encode(course)
			return
		}
	}
	w.WriteHeader(http.StatusNotFound)
	json.NewEncoder(w).Encode(map[string]string{"message": "No course found with the given ID"})
}

// @Summary Create a new course
// @Description Adds a new course to the system
// @Accept json
// @Produce json
// @Param course body Course true "Course data to create"
// @Success 201 {object} Course "Successfully created course"
// @Failure 400 {object} map[string]string "Please send some data"
// @Failure 400 {object} map[string]string "No data inside the request body"
// @Router /courses [post]
func createOneCourse(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Body == nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"message": "Please send some data"})
		return
	}
	var course Course
	if err := json.NewDecoder(r.Body).Decode(&course); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}

	if course.IsEmpty() {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"message": "No data inside the request body"})
		return
	}

	rand.Seed(time.Now().UnixNano())
	course.CourseId = strconv.Itoa(rand.Intn(100000000))
	courses = append(courses, course)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(course)
}

// @Summary Update an existing course
// @Description Updates the details of a course by its ID
// @Accept json
// @Produce json
// @Param course_id path string true "Course ID"
// @Param course body Course true "Updated course data"
// @Success 200 {object} Course
// @Failure 404 {object} map[string]string "No course found with the given ID"
// @Router /courses/{course_id} [put]
func updateOneCourse(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	var updatedCourse Course
	if err := json.NewDecoder(r.Body).Decode(&updatedCourse); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}

	for index, course := range courses {
		if course.CourseId == params["course_id"] {
			courses[index] = updatedCourse
			courses[index].CourseId = params["course_id"] // Ensure ID doesn't change
			json.NewEncoder(w).Encode(courses[index])
			return
		}
	}

	w.WriteHeader(http.StatusNotFound)
	json.NewEncoder(w).Encode(map[string]string{"message": "No course found with the given ID"})
}

// @Summary Delete a course
// @Description Deletes a course by its unique ID
// @Produce json
// @Param course_id path string true "Course ID"
// @Success 200 {object} map[string]string "Course deleted successfully"
// @Failure 404 {object} map[string]string "No course found with the given ID"
// @Router /courses/{course_id} [delete]
func deleteOneCourse(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	for index, course := range courses {
		if course.CourseId == params["course_id"] {
			courses = append(courses[:index], courses[index+1:]...)
			json.NewEncoder(w).Encode(map[string]string{"message": "Course deleted successfully"})
			return
		}
	}
	w.WriteHeader(http.StatusNotFound)
	json.NewEncoder(w).Encode(map[string]string{"message": "No course found with the given ID"})
}

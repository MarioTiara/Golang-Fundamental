package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

// Model for course - file
type Course struct {
	CourseId    string  `json:"courseid"`
	CourseName  string  `json:"coursename"`
	CoursePrice int     `json:"price"`
	Author      *Author `json:""`
}

type Author struct {
	Fullname string `json:"fullname"`
	Website  string `json:"website"`
}

// fake db
var courses []Course

// midleware, helper -file
func (c *Course) IsEmpty() bool {
	return c.CourseId == "" && c.CourseName == ""
}
func main() {

}

//controller - file

// server home route
func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1> Welcome to first mario API using Golang"))
}

func getAllCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get all course")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(courses)
}

func getOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("GetCourse")
	w.Header().Set("Content-Type", "applicaion/json")

	//garb id from request
	params:=mux.Vars(r)

	//loop through courses, find matching id and return the response
	for _, course :=range courses {
		if course.CourseId ==params["id"]{
			json.NewEncoder(w).Encode(course)
			return 
		}
	}

	json.NewEncoder(w).Encode("No Course found with given id")
	return

}

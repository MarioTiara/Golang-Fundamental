package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

//Model for course - file
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

//server home route
func serveHome (w http.ResponseWriter, r *http.Request){
	w.Write([] byte("<h1> Welcome to first mario API using Golang"))
}

func getAllCourse (w http.ResponseWriter, r *http.Request){
	fmt.Println("Get all course")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(courses)
}

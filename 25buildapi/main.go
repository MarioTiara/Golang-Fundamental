package main

import (
	"encoding/json"
	"fmt"
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
	//return c.CourseId == "" && c.CourseName == ""
	return c.CourseName==""
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

func  createOneCourse (w http.ResponseWriter, r * http.Request){
	fmt.Println("Create one course")
	w.Header().Set("Content-Type", "application/json")

	//what is : body is empty
	if r.Body == nil{
		json.NewEncoder(w).Encode("please send some data")
	}

	//what abut
	var course Course
	_ =json.NewDecoder(r.Body).Decode(&course)
	if course.IsEmpty(){
		json.NewEncoder(w).Encode("No data inside JSON")
		return
	}

	//generate unique id, string 
	//append course into course

	rand.Seed(time.Now().UnixMicro())
	course.CourseId=  strconv.Itoa(rand.Intn(100))
	courses= append(courses, course)
	json.NewEncoder(w).Encode(course)
	return
}

func updateOneCourse (w http.ResponseWriter, r * http.Request){
	fmt.Println("Update one course")
	w.Header().Set("Content-Type", "application/json")

	//first - grab id from req
	params:=mux.Vars(r)

	//loop, id, remove, add with my Id
	for index , course := range courses {
		if course.CourseId== params["id"]{
			courses = append(courses[:index], courses[index+1:]...)
			var course Course
			_ =json.NewDecoder(r.Body).Decode(& course)
			course.CourseId=params["id"]
			courses=append(courses, course)
			json.NewEncoder(w).Encode(course)
			return
		}
	}

	//TODO : send a response when is is not found

}

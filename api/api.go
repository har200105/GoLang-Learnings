package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("Making API calls")
	// performAPIRequest()
	encodeJSON()
}

func performAPIRequest() {

	const myURL = "http://127.0.0.1:8000/api/common/session_id"

	response, err := http.Get(myURL)
	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	content, _ := ioutil.ReadAll(response.Body)
	// fmt.Println(string(content))

	var responseBuilder strings.Builder
	responseBuilder.Write(content)

	fmt.Println(responseBuilder.String())

}

func performPostFormRequest() {
	data := url.Values{}
	data.Add("s", "s")
}

type courses struct {
	Name string `json:"namedata"`
	Age  int    `json:"age,omitempty"`
}

func encodeJSON() {

	coursesData := []courses{
		{"name", 1},
		{"name1", 5},
	}

	finalJson, _ := json.MarshalIndent(coursesData, "", "\t")

	fmt.Printf("%s\n", finalJson)
}

type Course struct {
	CourseId    string  `json:"courseid"`
	CourseName  string  `json:"coursename"`
	CoursePrice int     `json:"price"`
	Author      *Author `json:"author"`
}

type Author struct {
	FullName string `json:"fullname"`
	Website  string `json:"website"`
}

var coursesSlice []Course

func (c *Course) isEmpty() bool {
	return c.CourseName == ""
}

func getAllCourses(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(coursesSlice)
}

func getOneCourse(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	for _, course := range coursesSlice {
		if course.CourseId == params["courseId"] {
			json.NewEncoder(w).Encode(course)
		}
	}
	json.NewEncoder(w).Encode("No course found with this ID")

}

func createOneCourse(w http.ResponseWriter, r *http.Request) {
	if r.Body == nil {

	}

	var course Course
	json.NewDecoder(r.Body).Decode(&course)
	if course.isEmpty() {

	}
	rand.Seed(time.Now().UnixNano())
	course.CourseId = strconv.Itoa(rand.Intn(100))
	coursesSlice = append(coursesSlice, course)
	json.NewEncoder(w).Encode(course)

}

func updateOneCourse(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	for index, course := range coursesSlice {
		if course.CourseId == params["id"] {
			coursesSlice = append(coursesSlice[:index], coursesSlice[index+1:]...)
			var course Course
			json.NewDecoder(r.Body).Decode(&course)
			course.CourseId = params["id"]
			coursesSlice = append(coursesSlice, course)
			json.NewEncoder(w).Encode(course)
		}
	}
}

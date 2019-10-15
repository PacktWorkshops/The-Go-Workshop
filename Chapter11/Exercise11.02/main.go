package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type student struct {
	StudentId     int      `json:"id"`
	LastName      string   `json:"lname"`
	MiddleInitial string   `json:"mname,omitempty"`
	FirstName     string   `json:"fname"`
	IsMarried     bool     `json:"-"`
	IsEnrolled    bool     `json:",omitempty"`
	Courses       []course `json:"classes"`
}

type course struct {
	Name   string `json:"coursename"`
	Number int    `json:"coursenum"`
	Hours  int    `json:"coursehours"`
}

func main() {
	s := newStudent(1, "Williams", "s", "Felicia", false, false)
	student1, err := json.MarshalIndent(s, "", "    ")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(string(student1))
	fmt.Println()

	s2 := newStudent(2, "Washington", "", "Bill", true, true)

	c := course{Name: "World Lit", Number: 101, Hours: 3}
	s2.Courses = append(s2.Courses, c)
	c = course{Name: "Biology", Number: 201, Hours: 4}
	s2.Courses = append(s2.Courses, c)
	c = course{Name: "Intro to Go", Number: 101, Hours: 4}
	s2.Courses = append(s2.Courses, c)

	student2, err := json.MarshalIndent(s2, "", "    ")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(string(student2))
}

func newStudent(studentID int, lastName, middleInitial, firstName string,
	isMarried, isEnrolled bool) student {
	s := student{StudentId: studentID,
		LastName:      lastName,
		MiddleInitial: middleInitial,
		FirstName:     firstName,
		IsMarried:     isMarried,
		IsEnrolled:    isEnrolled,
	}
	return s
}

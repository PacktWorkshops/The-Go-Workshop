package main

import (
	"encoding/json"
	"fmt"
)

type student struct {
	StudentId     int      `json:"id"`
	LastName      string   `json:"lname"`
	MiddleInitial string   `json:"minitial"`
	FirstName     string   `json:"fname"`
	IsEnrolled    bool     `json:"enrolled"`
	Courses       []course `json:"classes"`
}

type course struct {
	Name   string `json:"coursename"`
	Number int    `json:"coursenum"`
	Hours  int    `json:"coursehours"`
}

func main() {
	data := []byte(`
	    {
	    	"id": 123,
	    	"lname": "Smith",
	    	"minitial": null,
	    	"fname": "John",
	    	"enrolled": true,
	    	"classes": [{
	    		"coursename": "Intro to Golang",
	    		"coursenum": 101,
	    		"coursehours": 4
	    	},
		{
	    		"coursename": "English Lit",
	    		"coursenum": 101,
	    		"coursehours": 3
	    	},
		{
	    		"coursename": "World History",
	    		"coursenum": 101,
	    		"coursehours": 3
	    	}
	
	]
	    }
	`)

	var s student
	err := json.Unmarshal(data, &s)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(s)
}

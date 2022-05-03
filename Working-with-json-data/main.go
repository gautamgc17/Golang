package main

import (
	"encoding/json"
	"fmt"
)


type course struct {
	// setting aliases for data members in json format
	// `` string template, set any type or variable as string
	Name     string `json:"coursename"`
	Price    int
	Platform string   `json:"website"`
	Password string   `json:"-"` // remove the field while consuming json data
	Tags     []string `json:"tags,omitempty"`
}


func EncodeJSON() {

	lcoCourses := []course{
		{"ReactJS Bootcamp", 299, "Udemy", "abc12", []string{"web-dev", "react"}},
		{"Mern Bootcamp", 499, "Udemy", "too12", []string{"web", "mern"}},
		{"Angular Bootcamp", 199, "Udemy", "mro12", []string{"dev", "angular"}},
		{"Vue Bootcamp", 199, "Udemy", "abc12", nil},
	}

	// package this data as JSON data - Marshal returns the JSON encoding of object
	// finalJSON, err := json.Marshal(lcoCourses)

	// MarshalIndent is like Marshal but applies Indent to format the output
	// Each JSON element in the output will begin on a new line beginning with prefix
	finalJSON, err := json.MarshalIndent(lcoCourses, "", "\t")
	if err != nil {
		panic(err)
	}
	fmt.Println(finalJSON)
	fmt.Printf("JSON Data %s \n", finalJSON)
	fmt.Printf("Type is: %T", finalJSON)
}


func DecodeJSON() {
	// byte is an alias for uint8
	// data from web is in byte format so we wrap it around string format
	// here we consume byte data nd decode it in json format
	jsonDATA := []byte(`
	{
		"coursename": "Angular Bootcamp",
		"Price": 199,
		"website": "Udemy",
		"tags": ["dev", "angular"]
	}
	`)

	var lcoCourse course
	// Valid reports whether data is a valid JSON encoding.
	checkValid := json.Valid(jsonDATA)

	if checkValid {
		fmt.Println("JSON Data was valid")
		// Unmarshal parses the JSON-encoded data and stores the result in the value pointed to by v.
		// Reference is passed so as to store result into it and not its copy
		json.Unmarshal(jsonDATA, &lcoCourse)
		// to print structs, this format is used
		fmt.Printf("%#v \n", lcoCourse)
	} else {
		fmt.Println("Not a valid JSON Data format")
	}

	// some cases where you want to add data to key-value pairs and construct structs
	// interface is another name for structs, since we dont know what will be the value so we make it interface
	var myJSONdata map[string]interface{}
	json.Unmarshal(jsonDATA, &myJSONdata)
	fmt.Printf("%#v \n" , myJSONdata)

	for key, val := range myJSONdata {
		fmt.Printf("Key is %v , Val is %v and type of data is %T \n", key, val, val)
	}
}


func main() {
	fmt.Println("Creating, handling and consuming JSON Data")
	// EncodeJSON()
	DecodeJSON()
}

package main

import (
	"fmt"
	"net/http"
	"log"
)

func main() {
	fmt.Println("Building a simple web server")

	// FileServer returns a handler that serves (mapping) HTTP requests with the contents of the file system rooted at root
	fileServer := http.FileServer(http.Dir("./static"))
	fmt.Printf("Type of File handler is: %T\n", fileServer)
	
	// Register an object(type) with http server that responds to HTTP request
	http.Handle("/", fileServer)

	// Register function with http server that responds to HTTP request
	http.HandleFunc("/hello", helloHandler)
	http.HandleFunc("/form", formHandler)

	fmt.Println("Starting server at port 8080")

	// ListenAndServe starts an HTTP server with a given address and handler. The handler is usually nil, which means to use DefaultServeMux.
	if err := http.ListenAndServe(":8080", nil); err!=nil{
		// Fatal is equivalent to Print() followed by a call to os.Exit(1).
		log.Fatal(err)
	}
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	// A ResponseWriter interface is used by an HTTP handler to construct an HTTP response
	if r.URL.Path != "/hello"{
		// Error replies to the request with the specified error message and HTTP code.
		http.Error(w, "404 Not Found", http.StatusNotFound)
		return
	}

	if r.Method != "GET"{
		http.Error(w, "Method is not Supported", http.StatusBadRequest)
		return 
	}
	// Fprintf formats according to a format specifier and writes to writer object w.
	fmt.Fprintf(w, "Hello World from Golang Server !!")
}

func formHandler(w http.ResponseWriter, r *http.Request) {
	// It reads the request body, parses it as a form and puts the results into both r.PostForm and r.Form
	if err := r.ParseForm(); err != nil{
		fmt.Fprintf(w, "Error in Parsing Form: %v", err)
	}
	fmt.Fprintf(w, "POST Request Successful !! \n\n")
	name := r.FormValue("name")
	address := r.FormValue("address")
	fmt.Fprintf(w, "Name: %s \n", name)
	fmt.Fprintf(w, "Address: %s", address)
}

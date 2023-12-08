package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

type Person struct {
	Id   int
	Name string
}
type PeopleResponse struct {
	People []Person
}

func ReturnJson(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	p := Person{
		Id:   1,
		Name: "a person",
	}

	people := PeopleResponse{People: []Person{p}}

	fmt.Println(people.People)

	// json.NewEncoder(w).Encode(people)

	/* alternative way for json response */
	data, _ := json.Marshal(people)
	w.Write(data)
}

func hello(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "hello\n")

}

func headers(w http.ResponseWriter, req *http.Request) {
	for name, headers := range req.Header {
		for _, h := range headers {
			fmt.Fprintf(w, "%v: %v\n", name, h)
		}
	}
}

func getImage(w http.ResponseWriter, r *http.Request) {
	f, _ := os.Open("./cardbox.jpeg")

	// Read the entire JPG file into memory.
	reader := bufio.NewReader(f)
	content, _ := io.ReadAll(reader)

	// Set the Content Type header.
	w.Header().Set("Content-Type", "image/jpeg")

	// Write image to the response.
	w.Write(content)
}

func main() {
	http.HandleFunc("/hello", hello)
	http.HandleFunc("/headers", headers)
	http.HandleFunc("/image", getImage)
	http.HandleFunc("/json", ReturnJson)

	http.ListenAndServe(":8090", nil)
}

// todo, JSON body
// router params
// query params
// sqlite?

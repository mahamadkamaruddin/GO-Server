package main

import (
	"fmt"
	"log"
	"net/http"
)

func funcHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hello" {
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}

	if r.Method != "GET" {
		http.Error(w, "Method is not supported", http.StatusNotFound)
		return
	}
	fmt.Fprintf(w, "hello!")
	return
}

func formHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method == "GET" {
		http.ServeFile(w, r, "./static/form.html")
		return
	}
	if r.Method == "POST" {
		if err := r.ParseForm(); err != nil {
			fmt.Fprintf(w, "ParseForm() error %v", err)
			return
		}
		fmt.Fprintf(w, "POST request successful\n")
		name := r.FormValue("name")
		address := r.FormValue("address")
		fmt.Fprintf(w, "Name = %v\n", name)
		fmt.Fprintf(w, "address = %v\n", address)
		return
	}

}

func main() {
	fileserver := http.FileServer(http.Dir("./static"))

	http.Handle("/", fileserver)
	http.HandleFunc("/form", formHandler)
	http.HandleFunc("/hello", funcHandler)

	fmt.Printf("Starting the server at port 9002\n")

	if err := http.ListenAndServe(":9002", nil); err != nil {
		log.Fatal(err)
	}

}

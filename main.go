package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func formHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		return
	}
	fmt.Fprintf(w, "Post request successful\n")
	name := r.FormValue("name")
	address := r.FormValue("address")
	fmt.Fprintf(w, "Name=%s, address=%s", name, address)
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hello" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}
	if r.Method != "GET" {
		http.Error(w, "the method is not supported", http.StatusNotFound)
		return
	}
	fmt.Fprintf(w, "Hello")
}

func parseTxt(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/txt" {
		http.Error(w, "Information of my block not found. Error ", http.StatusNotFound)
		return
	}
	if r.Method != "GET" {
		http.Error(w, "the method is not supported", http.StatusNotFound)
		return
	}
	file, err := os.OpenFile("./static/info.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Fprintf(w, "Error opening file")
		return
	}
	defer file.Close()

	text := "Danya pidor\n"

	if _, err1 := file.WriteString(text); err1 != nil {
		fmt.Fprintf(w, "Error writing to file")
		return
	}
	fmt.Fprintf(w, "File has been saved successfully\n")
}

func main() {
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)
	http.HandleFunc("/form", formHandler)
	http.HandleFunc("/hello", helloHandler)
	http.HandleFunc("/txt", parseTxt)

	fmt.Printf("Starting our server at port 8080\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}

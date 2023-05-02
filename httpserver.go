package main

import (
	"io"
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/", version)
	http.HandleFunc("/next", version1)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func version(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "golang version 1.12")
}
func version1(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "golang version 2.0")
}

package main

import (
	"io"
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/", version)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func version(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "golang version 1.12")
}

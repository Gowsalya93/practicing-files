package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Methods struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

func method(W http.ResponseWriter, r *http.Request) {
	method := `{"Name":"GET","Description":"gets all or a specified data"}`
	Bytes := []byte(method)

	var m Methods
	json.Unmarshal(Bytes, &m)
	fmt.Println(m.Name, m.Description)
}
func main() {
	http.HandleFunc("/", method)
	http.ListenAndServe(":8000", nil)
}

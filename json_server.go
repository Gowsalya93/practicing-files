package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type BMW struct {
	Name   string `json:"name"`
	Colour string `json:"colour"`
	Engine string `json:"engine"`
}

bmw := `{"name":"BMW 5series sedan","colour":"metallic blue","engine":"diesel engine"}`,
func CreateNewData(W http.ResponseWriter, r *http.Request) {
	
	var bmw BMW
	json.Unmarshal(reqBody, &bmw)
	json.NewEncoder(W).Encode(bmw)
	newData, err := json.Marshal(bmw)
	if err != nil {
		fmt.Println(err)

	} else {
		fmt.Println(string(newData))
	}
}
func main() {
	
	http.HandleFunc("bmw", CreateNewData).Methods("POST")
	log.Fatal(http.ListenAndServe(":8000", nil))
}

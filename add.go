package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Request struct {
	Operand string `json:"operand"`
	Number1 int    `json:"number1"`
	Number2 int    `json:"number2"`
}
type Response struct {
	Error  string `json:"error"`
	Result int    `json:"result"`
}

func mathOperation(w http.ResponseWriter, r *http.Request) {
	request := &Request{}
	err := json.NewDecoder(r.Body).Decode(request)
	if err != nil {
		fmt.Println(err)
	}
	response := &Response{}
	switch request.Operand {
	case "+":
		response.Result = (request.Number1 + request.Number2)
	default:
		response.Error = fmt.Sprintf("unknown operation: %s", request.Operand)
	}
	w.Header().Set("content-Type", "application/json")
	if response.Error != "" {
		w.WriteHeader(http.StatusBadRequest)
	}
	json.NewEncoder(w).Encode(response)

}
func main() {
	http.HandleFunc("/", mathOperation)
	http.ListenAndServe(":8080", nil)
}

package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	api "github.com/practicing-files/api/qm"
)

// Define a struct to hold your handler dependencies, such as the API client
type Handler struct {
	Client *api.Client
}

// Example handler function for handling requests to create an issue
func (h *Handler) CreateIssueHandler(w http.ResponseWriter, r *http.Request) {
	// Decode the request body
	var reqBody api.CreateIssueJSONRequestBody
	err := json.NewDecoder(r.Body).Decode(&reqBody)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Call the API client method to create the issue
	resp, err := h.Client.CreateIssue(context.Background(), reqBody)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	// Check the response status code
	if resp.StatusCode != http.StatusCreated {
		// Decode the error response
		var errResp api.Error
		err = json.NewDecoder(resp.Body).Decode(&errResp)
		if err != nil {
			http.Error(w, err.Error(), resp.StatusCode)
			return
		}

		// Return the error message
		http.Error(w, errResp.Message, resp.StatusCode)
		return
	}

	// Decode the successful response
	var issueResp api.Issue
	err = json.NewDecoder(resp.Body).Decode(&issueResp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Write the response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(issueResp)
}

func (h *Handler) GetAllIssuesHandler(w http.ResponseWriter, r *http.Request) {
	// Parse query parameters
	pageNumber, err := strconv.Atoi(r.URL.Query().Get("page"))
	if err != nil {
		http.Error(w, "Invalid page number", http.StatusBadRequest)
		return
	}

	pageSizeNumber, err := strconv.Atoi(r.URL.Query().Get("pageSize"))
	if err != nil {
		http.Error(w, "Invalid page size", http.StatusBadRequest)
		return
	}

	params := &api.GetAllIssuesParams{
		Page:     (*api.PageNumber)(&pageNumber),
		PageSize: (*api.PageSize)(&pageSizeNumber),
	}

	// Call the API client method to get all issues
	resp, err := h.Client.GetAllIssues(context.Background(), params)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	// Check the response status code
	if resp.StatusCode != http.StatusOK {
		// Decode the error response
		var errResp api.Error
		err = json.NewDecoder(resp.Body).Decode(&errResp)
		if err != nil {
			http.Error(w, err.Error(), resp.StatusCode)
			return
		}

		// Return the error message
		http.Error(w, errResp.Message, resp.StatusCode)
		return
	}

	// Decode the successful response
	var issuesResp []api.Issue
	err = json.NewDecoder(resp.Body).Decode(&issuesResp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Write the response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(issuesResp)
}

func main() {
	// Create an instance of your API client
	client, err := api.NewClient("https://api.example.com")
	if err != nil {
		fmt.Println("Error creating client:", err)
		return
	}

	// Create an instance of the handler with the client dependency
	handler := &Handler{
		Client: client,
	}

	// Define your HTTP routes
	http.HandleFunc("/create-issue", handler.CreateIssueHandler)
	http.HandleFunc("/get-all-issues", handler.GetAllIssuesHandler)

	// Start the HTTP server
	fmt.Println("Server listening on port 8080...")
	http.ListenAndServe(":8080", nil)
}

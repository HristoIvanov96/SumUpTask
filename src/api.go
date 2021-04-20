package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func HandleRequests(router *mux.Router, portNumber string) {
	//Expose endpoints
	router.HandleFunc("/tasks/json", sortTasks)
	router.HandleFunc("/tasks/bash", outputBashCommands)
	fmt.Printf("Serving requests ...")
	log.Fatal(http.ListenAndServe(portNumber, router))
}

func sortTasks(w http.ResponseWriter, r *http.Request) {
	var tasks Tasks
	//Decode tasks from request body
	err := json.NewDecoder(r.Body).Decode(&tasks)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	//Sort and format tasks
	formattedTasks, err := FormatSortedTasks(SortTasks(tasks.Tasks))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	//Return response with a json representation of the formatted tasks
	formattedOutput, err := json.MarshalIndent(formattedTasks, "", "  ")
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	fmt.Fprintf(w, string(formattedOutput))
}

func outputBashCommands(w http.ResponseWriter, r *http.Request) {
	var tasks Tasks
	//Decode tasks from request body
	err := json.NewDecoder(r.Body).Decode(&tasks)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	//Return response with sorted and formatted tasks
	sortedTasks := SortTasks(tasks.Tasks)
	output, err := FormatCommands(sortedTasks)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	fmt.Fprintf(w, output)
}



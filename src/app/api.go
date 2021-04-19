package app

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func HandleRequests(router *mux.Router, portNumber string) {
	router.HandleFunc("/sortTasks", sortTasks)
	router.HandleFunc("/bash", bash)
	fmt.Printf("Serving requests ...")
	log.Fatal(http.ListenAndServe(portNumber, router))
}

func sortTasks(w http.ResponseWriter, r *http.Request) {
	var tasks Tasks
	err := json.NewDecoder(r.Body).Decode(&tasks)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	formattedTasks := FormatSortedTasks(SortTasks(tasks.Tasks))
	formattedOutput, err := json.MarshalIndent(formattedTasks, "", "  ")
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	fmt.Fprintf(w, string(formattedOutput))
}

func bash(w http.ResponseWriter, r *http.Request) {
	var tasks Tasks
	err := json.NewDecoder(r.Body).Decode(&tasks)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	sortedTasks := SortTasks(tasks.Tasks)
	fmt.Fprintf(w, FormatCommands(sortedTasks))
}



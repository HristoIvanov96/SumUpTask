package app

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func HandleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/sortTasks", sortTasks)
	myRouter.HandleFunc("/bash", bash)
	log.Fatal(http.ListenAndServe(":4000", myRouter))
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



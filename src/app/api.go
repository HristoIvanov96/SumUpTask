package app

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func getTasks(w http.ResponseWriter, r *http.Request){
	var tasks Tasks
	err := json.NewDecoder(r.Body).Decode(&tasks)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	formattedTasks := FormatSortedTasks(SortTasks(tasks.Tasks))
	formattedOutput, err := json.MarshalIndent(formattedTasks, "", "  ")
	fmt.Fprintf(w, string(formattedOutput))
}

func HandleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/tasks", getTasks)
	log.Fatal(http.ListenAndServe(":4000", myRouter))
}



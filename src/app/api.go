package app

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func getTasks(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "Sorting tasks..")
	fmt.Println("Endpoint Hit: getTasks")
	var tasks Tasks
	err := json.NewDecoder(r.Body).Decode(&tasks)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	fmt.Fprintf(w, "%+v", SortTasks(tasks.Tasks))

}

func HandleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/tasks", getTasks)
	log.Fatal(http.ListenAndServe(":4000", myRouter))
}



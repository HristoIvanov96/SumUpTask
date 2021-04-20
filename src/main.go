package main

import (
	"SumUpTask/http"
	"github.com/gorilla/mux"
)

func main() {
	//Create a new router and set port number for api
	http.HandleRequests(mux.NewRouter(), ":4000")
}

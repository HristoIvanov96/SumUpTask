package main

import (
	"github.com/gorilla/mux"
)

func main() {
	//Create a new router and set port number for api
	HandleRequests(mux.NewRouter(), ":4000")
}

package main

import (
	"github.com/gorilla/mux"
)

func main() {
	HandleRequests(mux.NewRouter(), ":4000")
}

package main

import (
	"awesomeProject/src/app"
	"github.com/gorilla/mux"
)

func main() {
	app.HandleRequests(mux.NewRouter(), ":4000")
}

package main

import (
	// golang internal packages
	"log"
	"net/http"

	// local packages
	"github.com/programmingbunny/books/configs"
	routes "github.com/programmingbunny/books/service"

	// external packages
	"github.com/gorilla/mux"
)

func main() {
	// create a new router instance
	router := mux.NewRouter()

	// configure routes in service package into router just created
	routes.Routes(router)

	// create server
	log.Fatal(http.ListenAndServe(configs.Port, router))
}

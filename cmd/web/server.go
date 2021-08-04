package main

import (
	"fmt"
	//_handler "handling-json-in-go/pkg/handlers"
	"log"
	"net/http"

	"github.com/gorilla/mux"

)

func homePage(w http.ResponseWriter, re *http.Request) {
	fmt.Fprintf(w, "HomePage  one love")
}

func handleRequest() {

	myRouter := mux.NewRouter().StrictSlash(true)
	//myRouter.HandleFunc("/test", handler.HandleTest).Methods("GET")
	myRouter.HandleFunc("/", homePage)

	log.Fatal(http.ListenAndServe(":8081", myRouter))
}

func main() {
	handleRequest()
}

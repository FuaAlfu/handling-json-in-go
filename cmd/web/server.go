package main

import (
	"fmt"
	handler "handling-json-in-go/pkg/handlers"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func homePage(w http.ResponseWriter, re *http.Request) {
	fmt.Fprintf(w, "HomePage  one love")
}

func handleRequest() {

	port := ":8080"
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/test", handler.HandleTest).Methods("GET")
	myRouter.HandleFunc("/another", handler.HandleAnotherTest("testing")).Methods("GET")
	myRouter.HandleFunc("/user", handler.HandleUser).Methods("POST")
	myRouter.HandleFunc("/", homePage)

	fmt.Println("serving:\t at port [", port, "]")
	log.Fatal(http.ListenAndServe(port, myRouter))
}

func main() {
	handleRequest()
}

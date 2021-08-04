package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

)

type (
	User struct {
		Email    string `json: "email"`
		Password string `json: "password"`
	}
)

func HandleUser(w http.ResponseWriter, r *http.Request) {
	var user User
	w.Header().Set("Content-Type", "application/json")
	json.NewDecoder(r.Body).Decode(&user)
	fmt.Println(user) //for test in console

	json.NewEncoder(w).Encode(user)
}

func HandleTest(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "reply")
}

func HandleAnotherTest(s string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// fmt.Printf("reply again to: %s", s)
		fmt.Fprintf(w, "reply again to: %s", s)
	}
}

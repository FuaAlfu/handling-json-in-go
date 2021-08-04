package handler

import (
	"fmt"
	"net/http"

)

func HandleTest(w http.ResponseWriter, r *http.Request) {
	fmt.Println("reply")
}

package handlers

import (
	"fmt"
	"net/http"
)

func HandleHealthCheck(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Health ok! Welcome to the API!")
}

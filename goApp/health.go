package main

import (
	"fmt"
	"net/http"
)

func healthCheck(w http.ResponseWriter, r *http.Request) {
	status := "Running"
	fmt.Fprintf(w, "App is - %v", status)
}

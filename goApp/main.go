package main

import (
	"fmt"
	"log"
	"net/http"
)

func main(){
	http.HandleFunc("/", server1)
	http.HandleFunc("/login", loginHandler)
	http.HandleFunc("/health", healthCheck)
	http.HandleFunc("/logout", logout)
	fmt.Printf("Starting go_app...\n")
	if err := http.ListenAndServe(":8001", nil); err != nil {
		log.Fatal(err)
	}
}

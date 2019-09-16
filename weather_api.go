package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"net/http"
	"os"
	s "strings"
)
func weather_api(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/error" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}
	if len(r.URL.Path) > 1{
		switch r.Method {
		case "GET":
			variable := s.Split(r.URL.Path, "/")
			if variable[1] == "shutdown"{
				os.Exit(0)
			}
			fmt.Fprintf(w, "%v\n", r.Cookies())
			fmt.Fprintf(w, "You have entered - %v", variable[1])
			result := Index(variable[1])
			if len(result) > 1{
				fmt.Fprintf(w, "\nEmail - %v", result)
			} else {
				fmt.Fprintf(w, "\nEmail - Not found")
			}
			if variable[1] == "harrypotter"{
				fmt.Fprintf(w, "\nYou have cracked the magic words")
				fmt.Fprintf(w, "\nmetadata -> %v", r)
			}
		case "POST":
			fmt.Fprintf(w, "We do not accept POST requests this way")
		}
	} else {
		switch r.Method {
		case "GET":
			fmt.Fprintf(w, "Welcome to weather api.")
		case "POST":
			fmt.Fprintf(w, "We do not accept POST requests this way")
		default:
			fmt.Fprintf(w, "Sorry, only GET and POST methods are supported.")
		}
	}
}

func main() {
	http.HandleFunc("/", weather_api)

	fmt.Printf("Starting weather api server...\n")
	if err := http.ListenAndServe(":9002", nil); err != nil {
		log.Fatal(err)
	}
}
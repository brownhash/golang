package main

import (
	"fmt"
	"log"
	"net/http"
	s "strings"
)
func http_handler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/error" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}
	if len(r.URL.Path) > 1{
		switch r.Method {
		case "GET":
			variable := s.Split(r.URL.Path, "/")
			fmt.Fprintf(w, "You have entered - %v", variable[1])
			if variable[1] == "harrypotter"{
				fmt.Fprintf(w, "\nYou have cracked the magic words")
			}
		case "POST":
			fmt.Fprintf(w, "We do not accept POST requests this way")
		}
	} else {
		switch r.Method {
		case "GET":
			http.ServeFile(w, r, "form.html")
		case "POST":
			if err := r.ParseForm(); err != nil {
				fmt.Fprintf(w, "ParseForm() err: %v", err)
				return
			}
			fmt.Fprintf(w, "Post from website! r.PostFrom = %v\n", r.PostForm)
			inp1 := r.FormValue("inp1")
			inp2 := r.FormValue("inp2")
			fmt.Fprintf(w, "Input 1 = %s\n", inp1)
			fmt.Fprintf(w, "Input 2 = %s\n", inp2)
			if(inp1 == "harry" && inp2 == "potter"){
				fmt.Fprintf(w, "You have cracked the magic words")
				fmt.Fprintf(w, "\nurl -> %v", r.URL.Path)
				fmt.Fprintf(w, "\nmetadata -> %v", r)
			} else {
				fmt.Fprintf(w, "You have not cracked the magic words")
			}
		default:
			fmt.Fprintf(w, "Sorry, only GET and POST methods are supported.")
		}
	}
}

func main() {
	http.HandleFunc("/", http_handler)

	fmt.Printf("Starting Goserver...\n")
	if err := http.ListenAndServe(":9001", nil); err != nil {
		log.Fatal(err)
	}
}
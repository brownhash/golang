package main

import (
    "log"
    "net/http"
)

func main() {
	var http_path = "/Users/username/golang/go_server/http/"
	http.Handle("/", http.FileServer(http.Dir(http_path)))

	log.Fatal(http.ListenAndServe(":80", nil))
}

package main

import (
	"net/http"
)

func server1(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./source/index.html")
}

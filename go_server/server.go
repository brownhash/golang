package main

import (
    "log"
    "net/http"
    //"os/exec"
    //"os"
    //"fmt"
)

func redirect(w http.ResponseWriter, r *http.Request) {

    http.Redirect(w, r, "http://localhost:5000", 301)
}

func main() {

	//cmd := exec.Command("/usr/bin/python", "/Users/harshitsharma/Documents/golang/src/golang/go_server/backend/backend_server.py &&")
	//cmd.Stdout = os.Stdout
	//cmd.Stderr = os.Stderr
	//err := cmd.Run()
	//if err != nil{ 
	//	fmt.Println("Backend server executed")
	//}

	http.HandleFunc("/", redirect)
	err_http := http.ListenAndServe(":8080", nil)

	if err_http != nil {
		log.Fatal("ListenAndServe: ", err_http)
	}
}

package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func getMetrics(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		fmt.Fprintf(w,"We dont support get requests here!")
	case "POST":
		if err := r.ParseForm(); err != nil {
			http.ServeFile(w, r, "./templates/login_error.html")
			return
		}
		username := r.FormValue("username")
		csrf := r.FormValue("csrf")
		instanceName := r.FormValue("instanceName")
		myvar := map[string]string{
			"uname": username,
			"csrf": csrf,
			"instanceName": instanceName,
		}
		f, err := os.Create("./templates/data")
		if err != nil {
			panic(err.Error())
		}
		n2, err := f.WriteString(myvar["instanceName"])
		if err != nil {
			panic(err.Error())
		}
		fmt.Println("%v", n2)
		goUser(w, "./templates/instance_metrics.html", myvar)
	default:
		fmt.Fprintf(w,"We only serve POST requests this way!")
	}
}

func getMetricsInst(w http.ResponseWriter, r *http.Request) {
	client := &http.Client{}
	req, _ := http.NewRequest("GET", "http://localhost:5000/get-metrics-sg", nil)
	resp, err := client.Do(req)

	if err != nil {
		fmt.Println("Errored when sending request to the server")
		return
	}

	defer resp.Body.Close()
	resp_body, _ := ioutil.ReadAll(resp.Body)
	fmt.Fprintf(w, string(resp_body))
}

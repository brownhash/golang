package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

func main(){
	http.HandleFunc("/", responseHandler)

	fmt.Printf("Starting weather server...\n")
	if err := http.ListenAndServe(":9002", nil); err != nil {
		log.Fatal(err)
	}

}

func weather(lat string, lon string, apiKey string) string{
	link := "https://api.forecast.io/forecast/"

	weather_url := link + apiKey + "/" + lat + "," + lon + "?callback=?"
	resp, err := http.Get(weather_url)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Fatal(err)
	}
	data := strings.Replace(strings.Replace(strings.Split(string(body), "/**/ typeof  === 'function' && ")[1], "(", "", 1), ");", "", 1)
	return data
}


func responseHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/info" {
		fmt.Fprintf(w, "%v", "Welcome to the weather api(Powered by forecast.io), created with love by Harshit Sharma")
	}
	if len(r.URL.Path) > 1{
		switch r.Method {
		case "GET":
			variable := strings.Split(r.URL.Path, "/")
			latLon := strings.Split(variable[1], "&")
			lat := latLon[0]
			lon := latLon[1]
			api_key := "f536d4c3330c0a1391370d1443cee848"

			weather_data := weather(lat, lon, api_key)
			fmt.Fprintf(w, "%v", weather_data)

		case "POST":
			fmt.Fprintf(w, "Sorry, only GET method is supported.")
		}
	} else {
		switch r.Method {
		case "GET":
			fmt.Fprintf(w, "%v", "Welcome to the weather api(Powered by forecast.io), created with love by Harshit Sharma")
		default:
			fmt.Fprintf(w, "Sorry, only GET method is supported.")
		}
	}
}
package main
//#include <time.h>
import "C"
import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func healthCheck(w http.ResponseWriter, r *http.Request) {
	var status = map[string]string{}
	count := 0
	db := dbConn()
	selDB, err := db.Query("SHOW TABLES")
	for selDB.Next() {
		var Tables_in_goapp string
		err = selDB.Scan(&Tables_in_goapp)
		if err != nil {
			panic(err.Error())
		}
		status[string(count)]=string(Tables_in_goapp)
		count = count + 1
	}
	if len(status) > 0 {
		fmt.Fprintf(w, "Running")
	}
}

func cpuReport(w http.ResponseWriter, r *http.Request){
	client := &http.Client{}
	req, _ := http.NewRequest("GET", "http://localhost:5000/metrics-mumbai", nil)
	resp, err := client.Do(req)

	if err != nil {
		fmt.Println("Errored when sending request to the server")
		return
	}

	defer resp.Body.Close()
	resp_body, _ := ioutil.ReadAll(resp.Body)
	fmt.Fprintf(w, string(resp_body))
}

func cpuReportSg(w http.ResponseWriter, r *http.Request){
	client := &http.Client{}
	req, _ := http.NewRequest("GET", "http://localhost:5000/metrics-singapore", nil)
	resp, err := client.Do(req)

	if err != nil {
		fmt.Println("Errored when sending request to the server")
		return
	}

	defer resp.Body.Close()
	resp_body, _ := ioutil.ReadAll(resp.Body)
	fmt.Fprintf(w, string(resp_body))
}
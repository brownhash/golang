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
func http_handler(w http.ResponseWriter, r *http.Request) {
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

type Employee struct {
	Name  string
	Email string
}

func dbConn() (db *sql.DB) {
	dbDriver := "mysql"
	dbUser := "root"
	dbPass := "WIKIman1612**#"
	dbName := "harrydbst"
	db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@/"+dbName)
	if err != nil {
		panic(err.Error())
	}
	return db
}

func Index(name string) string {

	db := dbConn()
	selDB, err := db.Query("SELECT * FROM data where name=?", name)
	if err != nil {
		panic(err.Error())
	}
	Email := ""
	for selDB.Next() {
		var name, email string
		err = selDB.Scan(&name, &email)
		if err != nil {
			panic(err.Error())
		}
		Email = email
	}
	defer db.Close()

	return Email
}

func main() {
	http.HandleFunc("/", http_handler)

	fmt.Printf("Starting Goserver...\n")
	if err := http.ListenAndServe(":9001", nil); err != nil {
		log.Fatal(err)
	}
}
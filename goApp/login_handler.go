package main

import (
	"fmt"
	//"fmt"
	_ "github.com/go-sql-driver/mysql"
	"net/http"
	s "strings"
)

func loginHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		http.ServeFile(w, r, "./source/error.html")
	case "POST":
		if err := r.ParseForm(); err != nil {
			http.ServeFile(w, r, "./source/login_error.html")
			return
		}
		uname := r.FormValue("username")
		pword := r.FormValue("password")

		loginFlag := login(uname, pword)

		if loginFlag["uname"] != "0" {

			csrf, err := r.Cookie("csrftoken")
			myvar := map[string]interface{}{
				"uname": loginFlag["uname"],
				"data": csrf,
			}
			if err != nil{
				panic(err.Error())
			}
			db := dbConn()
			token, err := db.Query("UPDATE user_data SET id=? WHERE username=?", csrf.Value, myvar["uname"])
			fmt.Println(token)
			goUser(w, "./source/user_main.html", myvar)
		} else {
			http.ServeFile(w, r, "./source/index_error.html")
		}

	default:
		fmt.Fprintf(w, "Sorry, only GET and POST methods are supported.")
	}
}

func login(username string, password string) map[string]interface{} {
	db := dbConn()
	selDB, err := db.Query("SELECT * FROM user_data where username=? and password=?", username, password)

	if err != nil {
		panic(err.Error())
	}
	loginFlag := map[string]interface{}{
		"uname": "",
		"data": "",
	}

	for selDB.Next() {
		var username, name, password, id string
		err = selDB.Scan(&name, &username, &password, &id)
		if err != nil {
			panic(err.Error())
		}

		if err != nil {
			panic(err.Error())
		}
		if len(username) > 0 {
			loginFlag["uname"] = username
			loginFlag["data"] = id
		} else {
			loginFlag["uname"] = "0"
			loginFlag["data"] = "0"
		}
	}
	defer db.Close()

	return loginFlag
}

func logout(w http.ResponseWriter, r *http.Request){
	if r.Method == "POST" {
		if err := r.ParseForm(); err != nil {
			http.ServeFile(w, r, "./source/error.html")
			return
		}
	}
	data := r.FormValue("user_value")
	listData := s.Split(data, "??")
	user:= listData[0]
	csrf:= listData[1]
	flag := clear_token(string(user), string(csrf))

	if flag == "true"{
		http.ServeFile(w, r, "./source/index.html")
	} else {
		http.ServeFile(w, r, "./source/error.html")
	}
}

func clear_token(uname string, csrf string) string{
	db := dbConn()
	selDB, err := db.Query("SELECT * FROM user_data where username=? and id=?", uname, csrf)

	if err != nil {
		panic(err.Error())
	}
	var flag string
	for selDB.Next() {
		var username, name, password, id string
		err = selDB.Scan(&name, &username, &password, &id)
		if err != nil {
			panic(err.Error())
		}
		fmt.Println(username)
		if len(username) > 0 {
			token, err := db.Query("UPDATE user_data SET id=? WHERE username=?", "0", uname)
			if err != nil {
				panic(err.Error())
			}
			fmt.Println(token)
			flag = "true"
		} else{
			flag = "false"
		}
	}
	return flag
}

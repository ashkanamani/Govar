package main

import (
	"net/http"

	"github.com/ashkanamani/govar/data"
)


func signup(w http.ResponseWriter, r *http.Request) {
	var err error 
	switch r.Method {
	case "GET":
		generateHTML(w, nil, "login.layout", "public.navbar", "signup")
		
	case "POST":
		err = r.ParseForm()
		if err != nil {
			logger.SetPrefix("ERROR ")
			logger.Println(err, "cannot pasrse file")
		}
		user := data.User{
			Name: r.PostFormValue("name"),
			Email: r.PostFormValue("email"),
			Password: r.PostFormValue("password"),
		}
		err = user.Create()
		if err != nil {
			logger.SetPrefix("ERROR ")
			logger.Println(err, "cannot pasrse file")
		}
		http.Redirect(w, r, "/login", http.StatusFound)
	}
}

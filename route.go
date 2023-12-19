package main

import (
	"net/http"

	"github.com/ashkanamani/govar/data"
)


func signup(w http.ResponseWriter, r *http.Request) {
	var err error 
	switch r.Method {
	case "GET":
		generateHTML(w, nil, "layout", "public.navbar", "signup")
		
	case "POST":
		err = r.ParseForm()
		if err != nil {
			logger.SetPrefix("ERROR ")
			logger.Println(err, "cannot pasrse file")
		}
		logger.Println("22")
		user := data.User{
			Name: r.PostFormValue("name"),
			Email: r.PostFormValue("email"),
			Password: r.PostFormValue("password"),
		}
		logger.Println(user)

		err = user.Create()
		if err != nil {
			danger(err, "Cannot parse form")
		}
		http.Redirect(w, r, "/login", http.StatusFound)
	}
}

func login(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		generateHTML(w, nil, "layout", "public.navbar", "login")

	case "POST":
	}
}

func authenticate(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		err := r.ParseForm()
		if err != nil {
			danger(err, "Cannot parse form")
		}
		user, err := data.UserByEmail(r.FormValue("email"))
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			http.Redirect(w, r, "/login", http.StatusAccepted)
		}
		if user.Password == data.Encrypt(r.FormValue("password")) {
			session, err := user.CreateSession()
			if err != nil {
				danger(err, "can not create session")
			}
			cookie := http.Cookie{
				Name: "_cookie",
				Value: session.UUID,
				HttpOnly: true,
			}
			http.SetCookie(w, &cookie)
			http.Redirect(w, r, "/", http.StatusFound)
		} else {
			http.Redirect(w, r, "/login", http.StatusFound)
		}
	}
}


func logout(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("_cookie")
	if err != nil {
		warning(err, "failed to get cookie")
		return
	}
	session := data.Session{UUID: cookie.Value}
	session.DeleteByUUID()
	http.Redirect(w, r, "/", http.StatusFound)

}


func index(w http.ResponseWriter, r *http.Request) {
	_, err := session(w, r)
	if err != nil {
		generateHTML(w, nil, "layout", "public.navbar", "home")
	} else {
		generateHTML(w, nil, "layout", "private.navbar", "home")
	}
}
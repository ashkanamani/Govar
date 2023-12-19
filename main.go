package main

import (
	"flag"
	"net/http"
	"time"
)

func main() {

	
	flag.StringVar(&config.Db.Username, "dbuser", "postgres", "username to connect to database")
	flag.StringVar(&config.Db.Password, "dbpassword", "postgres", "password to connect to database")
	flag.Parse()
	
	mux := http.NewServeMux()

	files := http.FileServer(http.Dir("/public"))
	mux.Handle("/static/", http.StripPrefix("/static/", files))
	
	mux.HandleFunc("/", index)
	mux.HandleFunc("/signup", signup)
	mux.HandleFunc("/login", login)
	mux.HandleFunc("/authenticate", authenticate)
	mux.HandleFunc("/logout", logout)
	server := http.Server{
		Addr: config.Address,
		Handler: mux,
		ReadTimeout: time.Duration(config.ReadTimeout * int64(time.Second)),
		WriteTimeout: time.Duration(config.WriteTimeout * int64(time.Second)),
	}
	logger.Println("Starting application on port:", config.Address)
	server.ListenAndServe()
}

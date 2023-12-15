package main

import (
	"net/http"
	"time"
)

func main() {
	mux := http.NewServeMux()
	
	files := http.FileServer(http.Dir("/public"))
	mux.Handle("/static/", http.StripPrefix("/static/", files))
	
	mux.HandleFunc("/signup", signup)

	
	server := http.Server{
		Addr: config.Address,
		Handler: mux,
		ReadTimeout: time.Duration(config.ReadTimeout * int64(time.Second)),
		WriteTimeout: time.Duration(config.WriteTimeout * int64(time.Second)),
	}
	logger.Println("Starting application on port:", config.Address)
	server.ListenAndServe()
}

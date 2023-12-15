package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
)

type Configuration struct {
	Address      string
	ReadTimeout  int64
	WriteTimeout int64
}
var config Configuration
var logger *log.Logger

func init() {
	file, err := os.OpenFile("app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println("failed to create or open app.log file: ", err)
		os.Exit(1)
	}
	logger = log.New(file, "INFO", log.Ldate|log.Ltime|log.Lshortfile)
	loadConfig()

}
func loadConfig() {
	file, err := os.Open("config.json")
	if err != nil {
		logger.Fatalln("can not open config.json file ", err)
	}
	defer file.Close()
	decoder := json.NewDecoder(file)
	config = Configuration{}
	err = decoder.Decode(&config)
	if err != nil {
		logger.Fatalln("can not get configuration from file", err)
	}
}
func generateHTML(writer http.ResponseWriter, data interface{}, filenames ...string) {
	var files []string 
	for _, file := range filenames {
		files = append(files, fmt.Sprintf("templates/%s.gohtml", file))
	}
	templates := template.Must(template.ParseFiles(files...))
	templates.ExecuteTemplate(writer, "layout", data)
}

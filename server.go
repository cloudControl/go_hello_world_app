package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
)

func defaultHandler(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("templates/hello.html"))
	pageModel := make(map[string]interface{})
	pageModel["Domain"] = getEnv("DOMAIN", "cloudcontrolled.com")
	if err := t.Execute(w, pageModel); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func main() {
	http.Handle("/static/", http.FileServer(http.Dir(".")))
	http.HandleFunc("/", defaultHandler)

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", getEnv("PORT", "8080")), nil))
}

func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value != "" {
		return value
	}
	return defaultValue
}

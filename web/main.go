package main

import (
	"html/template"
	"log"
	"net/http"
)

func viewHandler(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("template.html")
	t.Execute(w, nil)
}

func main() {
	http.HandleFunc("/view/", viewHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

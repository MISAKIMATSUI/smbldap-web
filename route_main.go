package main

import (
	"html/template"
	"net/http"
)

func index(writer http.ResponseWriter, request *http.Request) {
	t := template.Must(template.ParseFiles("templates/index.html"))
	t.Execute(writer, "")
}

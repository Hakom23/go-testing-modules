package main

import (
	"net/http"
)

func Index(w http.ResponseWriter, req *http.Request) {
	renderTemplate(w, "index.html")
}
func Elements(w http.ResponseWriter, req *http.Request) {
	renderTemplate(w, "about.page.html")
}
func Generic(w http.ResponseWriter, req *http.Request) {
	renderTemplate(w, "generic.html")
}

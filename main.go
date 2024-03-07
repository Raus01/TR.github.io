package main

import (
	"fmt"
	"html/template"
	"net/http"
)

var tpl = template.Must(template.ParseGlob("templates/*.html"))

func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("index handler ", r.RequestURI)
	if r.URL.Path != "/" {
		http.Error(w, "404 Address Not Found", http.StatusNotFound)
		return
	}
	tpl.ExecuteTemplate(w, "index.html", nil)
}

func kruusHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("kruus handler ", r.RequestURI)
	if r.URL.Path != "/galerii" {
		http.Error(w, "404 Address Not Found", http.StatusNotFound)
		return
	}
	tpl.ExecuteTemplate(w, "galerii.html", nil)
}

func parkHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("kruus handler ", r.RequestURI)
	if r.URL.Path != "/park" {
		http.Error(w, "404 Address Not Found", http.StatusNotFound)
		return
	}
	tpl.ExecuteTemplate(w, "park.html", nil)
}

func müükHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("kruus handler ", r.RequestURI)
	if r.URL.Path != "/müük" {
		http.Error(w, "404 Address Not Found", http.StatusNotFound)
		return
	}
	tpl.ExecuteTemplate(w, "müük.html", nil)
}

func main() {
	port := 8080

	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/galerii", kruusHandler)
	http.HandleFunc("/park", parkHandler)
	http.HandleFunc("/müük", müükHandler)

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	fmt.Printf("Server listening on http://localhost:%d/\n", port)

	http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
}

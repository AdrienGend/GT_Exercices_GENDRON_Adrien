package main

import (
	"net/http"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "html/index.html")
}

func page1Handler(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/page1", http.StatusFound)
}

func page2Handler(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/page2", http.StatusFound)
}

func page3Handler(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/page3", http.StatusFound)
}

func notFoundHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	http.ServeFile(w, r, "html/error.html")
}

func main() {
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/page1", page1Handler)
	http.HandleFunc("/page2", page2Handler)
	http.HandleFunc("/page3", page3Handler)
	http.HandleFunc("/notfound", notFoundHandler)
	http.ListenAndServe(":8080", nil)
}

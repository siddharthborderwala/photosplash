package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "text/html")
	fmt.Fprint(w, "<h1>Welcome to photosplash</h1>")
}

func contact(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "text/html")
	fmt.Fprint(w, "Get in touch at <a href=\"mailto:support@photosplash.com\">support@photosplash.com</a>")
}

func hello(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "text/html")
	params := mux.Vars(r)
	fmt.Fprintf(w, "Hello %s", params["name"])
}

func notfound(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	fmt.Fprint(w, "The page you are looking for does not exist or has been moved")
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", home)
	r.HandleFunc("/contact", contact)
	r.HandleFunc("/hello/{name}", hello)

	r.NotFoundHandler = http.HandlerFunc(notfound)
	http.ListenAndServe(":3000", r)
	fmt.Println("Server listening at http://localhost:3000")
}

package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", Index)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// fmt.Printf("Hello, %q", html.EscapeString(r.URL.path))
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}


func Index(w http.ResponseWriter, r *http.Request) {
	fmt.
}
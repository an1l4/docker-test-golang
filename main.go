package main

import (
	"fmt"
	"html"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
	})

	http.HandleFunc("/hai", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hai")
	})
	fmt.Println("server running at port 8000")
	log.Fatal(http.ListenAndServe(":8000", nil))

}

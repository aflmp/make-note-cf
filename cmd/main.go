package main

import (
	"log"
	"net/http"

	function "github.com/aflmp/make-note-cf"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/make-note", function.MakeNote)
	err := http.ListenAndServe(":8080", mux)
	log.Fatal(err)
}

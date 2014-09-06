package main

import (
	"log"
	"net/http"
	"os"
)

func main() {
	// TODO(ttacon): allow user to specify directory and just default to cwd

	curr, err := os.Getwd()
	if err != nil {
		log.Fatal("os.Getwd(): ", err)
	}

	http.Handle("/", http.StripPrefix("/", http.FileServer(http.Dir(curr))))
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

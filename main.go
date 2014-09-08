package main

import (
	"flag"
	"log"
	"net/http"
	"os"
)

var (
	dir = flag.String("b", "", "base directory to serve from")
)

func main() {
	flag.Parse()

	curr, err := os.Getwd()
	if err != nil {
		log.Fatal("os.Getwd(): ", err)
	}

	if len(*dir) > 0 {
		if _, err := os.Stat(*dir); err != nil {
			log.Fatal("failed to detect dir, err: ", err)
		}
		curr = *dir
	}

	fs := http.FileServer(http.Dir(curr))

	http.Handle("/", http.StripPrefix("/", LoggerHandler{fs}))
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

type LoggerHandler struct {
	fs http.Handler
}

func (l LoggerHandler) ServeHTTP(resp http.ResponseWriter, req *http.Request) {
	log.Printf("%s %s\n", req.Method, req.RequestURI)
	l.fs.ServeHTTP(resp, req)
}

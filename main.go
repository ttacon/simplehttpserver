package main

import (
	"flag"
	"log"
	"net/http"
	"os"
	"strconv"
)

var (
	dir     = flag.String("b", "", "base directory to serve from")
	port    = flag.String("p", "8080", "port to serve files on")
	verbose = flag.Bool("v", false, "verbose logging")
)

func main() {
	flag.Parse()

	if _, err := strconv.Atoi(*port); err != nil {
		log.Fatal("port provided must be a valid integer")
	}

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
	addr := ":" + *port

	log.Println("Listening on", addr)
	if err := http.ListenAndServe(addr, nil); err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

type LoggerHandler struct {
	fs http.Handler
}

func (l LoggerHandler) ServeHTTP(resp http.ResponseWriter, req *http.Request) {
	args := []interface{}{req.Method, req.RequestURI}
	if *verbose {
		args = append(args, req.RemoteAddr)
	}

	log.Println(args...)
	l.fs.ServeHTTP(resp, req)
}

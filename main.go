package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
)

var (
	port = flag.String("port", os.Getenv("PORT"), "http port")
)

func init() {
    if *port == "" {
		*port = "8080"
	}
}

func main() {
	flag.Parse()

	http.Handle("/", indexHandler())
	http.Handle("/weclome", welcomeHandler())
    fmt.Printf("listening on port %v\n", *port)

    http.ListenAndServe(":"+ *port, nil)
}

func indexHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("YOREDDIT"))
	})
}

func welcomeHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Weclome back!"))
	})
}

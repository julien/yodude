package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"net/http"
	"os"
	"regexp"
	"strings"
	"time"
)

type yo struct {
	From     string    `json:"from"`
	Received time.Time `json:"received"`
}

var (
	port       = flag.String("port", os.Getenv("PORT"), "http port")
	userRegexp = regexp.MustCompile(`username=`)
	yos        []yo
)

func init() {
	if *port == "" {
		*port = "8080"
	}
}

func main() {

	flag.Parse()

	yos = make([]yo, 0)

	http.Handle("/welcome", welcomeHandler())
	http.Handle("/yo", yoHandler())
	http.Handle("/yos", yosHandler())
	fmt.Printf("listening on port %v\n", *port)

	http.ListenAndServe(":"+*port, nil)
}

func welcomeHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Return a template here
		w.Write([]byte("Weclome to YOREDDIT"))
	})
}

func yoHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		s := strings.Replace(r.URL.RawQuery, "username=", "", -1)
		t := time.Now()

		if s == "" {
			http.Error(w, "No username received", http.StatusBadRequest)
			return
		}

		y := yo{From: s, Received: t}
		yos = append(yos, y)

		enc := json.NewEncoder(w)
		w.Header().Set("Content-Type", "application/json")
		enc.Encode(yos)

	})
}

func yosHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		s := strings.Replace(r.URL.RawQuery, "username=", "", -1)

		out := make([]yo, 0)

		if s != "" {
			for i := 0; i < len(yos); i++ {
				if yos[i].From == s {
					out = append(out, yos[i])
				}
			}
		} else {
			out = yos
		}

		enc := json.NewEncoder(w)
		w.Header().Set("Content-Type", "application/json")
		enc.Encode(yos)
	})
}

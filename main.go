package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"net/http"
	"net/url"
	"os"
)

type yo struct {
	Username string `json:"username"`
	Location []string `json:"location"`
	URL      string `json:"url"`
	UserIp   string `json:"user_ip"`
}

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
	http.Handle("/static/", staticHandler())
	http.Handle("/yo", yoHandler())
	fmt.Printf("listening on port %v\n", *port)

	http.ListenAndServe(":"+*port, nil)
}

func indexHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "index.html")
	})
}

func staticHandler() http.Handler {
    return http.HandlerFunc(func (w http.ResponseWriter, r *http.Request) {
        http.ServeFile(w, r, r.URL.Path[1:])
    })
}

func yoHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		m, _ := url.ParseQuery(r.URL.RawQuery)
		if m == nil || m["username"] == nil {
			http.Error(w, "No username received", http.StatusBadRequest)
			return
		}

		// y := yo{
		// 	Username: m["username"][0],
		// 	Location: m["location"],
		// 	URL:      m["url"][0],
		// 	UserIp:   m["user_ip"][0],
		// }

		enc := json.NewEncoder(w)
		w.Header().Set("Content-Type", "application/json")
		enc.Encode([]byte("{\"message\": \"yo received\"}"))

	})
}

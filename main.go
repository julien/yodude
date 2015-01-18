package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"net/http"
	"net/url"
	"os"

	"github.com/melvinmt/firebase"
)

type yo struct {
	Location string   `json:"location"`
	URL      string   `json:"url"`
	Username string   `json:"username"`
	UserIP   string   `json:"user_ip"`
}

type response struct {
	Message string `json:"message"`
}

const (
	fbToken = "yHzyVMPbrQJilp8aDLpJUhkr6o4H6Xn23ADblY0D"
	fbURL   = "https://scorching-fire-3007.firebaseio.com/yos/"
)

var (
	port  = flag.String("port", os.Getenv("PORT"), "http port")
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
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
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

		_, err := createYO(m)
		var res response
		if err != nil {
			res.Message = "YO received but not saved"
		} else {
			res.Message = "YO received and saved"
		}

		enc := json.NewEncoder(w)
		w.Header().Set("Content-Type", "application/json")
		enc.Encode(&res)

	})
}

func createYO(m map[string][]string) (yo, error) {

	var b bool
	y := yo{}

	if b = m["location"] != nil && len(m["location"]) > 0; b {
		y.Location = m["location"][0]
	}
	if b = m["username"] != nil && len(m["username"]) > 0; b {
		y.Username = m["username"][0]
	}
	if b = m["url"] != nil && len(m["url"]) > 0; b {
		y.URL = m["url"][0]
	}
	if b = m["user_ip"] != nil && len(m["user_ip"]) > 0; b {
		y.UserIP = m["user_ip"][0]
	}

    ref := firebase.NewReference(fbURL)

	if err := ref.Push(y); err != nil {
        fmt.Println("Firebase error", err)

		return y, err
	}

	return y, nil
}



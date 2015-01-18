package main

import (
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func TestMain(m *testing.M) {

	// var t *testing.T
	// if *port != "" {
	// 	t.Errorf("got %v want 8080", *port)
	// }

	os.Exit(m.Run())
}

func TestIndexHandler(t *testing.T) {

	http.NewRequest("GET", "/", nil)
	w := httptest.NewRecorder()

	if w.Code != http.StatusOK {
		t.Errorf("got %v want 200", w.Code)
	}

}

func TestStaticHandler(t *testing.T) {

	h := staticHandler()
	r, _ := http.NewRequest("GET", "/static/style.css", nil)
	w := httptest.NewRecorder()

	h.ServeHTTP(w, r)

	if w.Code != http.StatusOK {
		t.Errorf("got %v want 200", w.Code)
	}
}

func TestYoHandler(t *testing.T) {
	h := yoHandler()
    r, _ := http.NewRequest("GET", "/yo?username=SOMEONE&location=45.234,3.8983&user_ip=0.0.0.0&url=http://localhost", nil)
	w := httptest.NewRecorder()

	h.ServeHTTP(w, r)

	if w.Code != http.StatusOK {
		t.Errorf("got %v want 200", w.Code)
	}
}

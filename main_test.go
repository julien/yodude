package main

import (
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func TestMain(m *testing.M) {

	var t *testing.T
    if *port != ":" {
		t.Errorf("got %v want 8080", *port)
	}

	os.Exit(m.Run())
}

func TestHomeHandler(t *testing.T) {

	h := indexHandler()
	r, _ := http.NewRequest("GET", "/", nil)
	w := httptest.NewRecorder()

	h.ServeHTTP(w, r)

	if w.Code != http.StatusOK {
		t.Errorf("got %v want 200", w.Code)
	}
}

func TestWelcomeHandler(t *testing.T) {

	h := welcomeHandler()
	r, _ := http.NewRequest("GET", "/welcome", nil)
	w := httptest.NewRecorder()

	h.ServeHTTP(w, r)

	if w.Code != http.StatusOK {
		t.Errorf("got %v want 200", w.Code)
	}
}

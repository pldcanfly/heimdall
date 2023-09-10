package server

import (
	"net/http"
)

func handleGetRoot(w http.ResponseWriter, r *http.Request) error {
	executePageTemplate(w, "index", nil)
	return nil
}

func handleGetTest(w http.ResponseWriter, r *http.Request) error {
	executeComponentTemplate(w, "button", nil)
	return nil
}

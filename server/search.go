package server

import (
	"encoding/json"
	"github.com/kefkius/nomen/node"
	"net/http"
)

type FilterStruct struct {
	Regexp     string `json:"regexp"`
	NumResults int    `json:"numResults"`
}

// FilterNames searches for names that match a regexp.
func FilterNames(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	var filter FilterStruct
	err := json.NewDecoder(r.Body).Decode(&filter)
	if err != nil {
		w.WriteHeader(500)
		return
	}
	result, err := node.FilterNames(filter.Regexp, filter.NumResults)
	if err != nil {
		w.WriteHeader(500)
		return
	}
	if err := json.NewEncoder(w).Encode(result); err != nil {
		w.WriteHeader(500)
	}
	return
}

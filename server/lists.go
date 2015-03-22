package server

import (
	"encoding/json"
	"github.com/kefkius/nomen/node"
	"net/http"
)

// ListNames retrieves a list of names that the node owns.
func ListNames(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	result, err := node.ListNames()
	if err != nil {
		w.WriteHeader(500)
		return
	}
	if result != nil {
		if err := json.NewEncoder(w).Encode(result); err != nil {
			w.WriteHeader(500)
			return
		}
	} else {
		w.WriteHeader(500)
		return
	}
}

// ListIds retrieves a list of identities that the node owns.
func ListIds(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	result, err := node.ListIds()
	if err != nil {
		w.WriteHeader(500)
		return
	}
	if result != nil {
		if err := json.NewEncoder(w).Encode(result); err != nil {
			w.WriteHeader(500)
			return
		}
	} else {
		w.WriteHeader(500)
		return
	}
}

// ListDomains retrieves a list of domains that the node owns.
func ListDomains(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	result, err := node.ListDomains()
	if err != nil {
		w.WriteHeader(500)
		return
	}
	if result != nil {
		if err := json.NewEncoder(w).Encode(result); err != nil {
			w.WriteHeader(500)
			return
		}
	} else {
		w.WriteHeader(500)
		return
	}
}

// ListExpiredNames retrieves expired names
func ListExpiredNames(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	result, err := node.ListExpiredNames()
	if err != nil {
		w.WriteHeader(500)
		return
	}
	if result != nil {
		if err := json.NewEncoder(w).Encode(result); err != nil {
			w.WriteHeader(500)
			return
		}
	} else {
		w.WriteHeader(500)
		return
	}
}

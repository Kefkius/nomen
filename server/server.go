package server

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/kefkius/nmcjson"
	"github.com/kefkius/nomen/node"
	"log"
	"net/http"
)

// Init starts up the server.
func Init(confUser, confPass, confServer string) {
	// Initialize nmcjson
	nmcjson.Init()
	// Initialize node
	node.Init(confUser, confPass, confServer)

	r := mux.NewRouter().StrictSlash(true)
	// Register HTML routes
	//r.PathPrefix("/").Handler(http.FileServer(http.Dir("public/")))
	// Register API routes
	api := r.PathPrefix("/api").Subrouter()
	for _, route := range apiRoutes {
		api.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			HandlerFunc(route.HandlerFunc)
	}
	r.PathPrefix("/").Handler(http.FileServer(http.Dir("public/")))

	//api.HandleFunc("/names", ListNames)
	//http.Handle("/", r)

	log.Fatal(http.ListenAndServe(":8000", r))
}

// GetBalance retrieves balance.
func GetBalance(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	result, err := node.GetBalance()
	if err != nil {
		w.WriteHeader(500)
		return
	}
	balanceMap := map[string]float64{
		"balance": result,
	}
	if err := json.NewEncoder(w).Encode(balanceMap); err != nil {
		w.WriteHeader(500)
		return
	}
	return
}

// ShowId retrieves an ID
func ShowId(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	vars := mux.Vars(r)
	identifier := vars["identifier"]
	result, err := node.ShowId(identifier)
	if err != nil {
		w.WriteHeader(500)
		return
	}
	if result.Name != "" {
		if err := json.NewEncoder(w).Encode(result); err != nil {
			w.WriteHeader(500)
			return
		}
	} else {
		w.WriteHeader(500)
		return
	}
}

// ShowDomain retrieves an domain
func ShowDomain(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	vars := mux.Vars(r)
	identifier := vars["identifier"]
	result, err := node.ShowDomain(identifier)
	if err != nil {
		w.WriteHeader(500)
		return
	}
	if result.Name != "" {
		if err := json.NewEncoder(w).Encode(result); err != nil {
			w.WriteHeader(500)
			return
		}
	} else {
		w.WriteHeader(500)
		return
	}
}

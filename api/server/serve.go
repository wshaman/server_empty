package server

import (
	"github.com/gorilla/mux"
)

func NewServer(prefix string) (*mux.Router, error) {
	r := mux.NewRouter()
	apiV1 := r.PathPrefix(prefix).Subrouter()
	for _, v:= range apiRoutes {
		apiV1.HandleFunc(v.Pattern, v.HandlerFunc).Methods(v.Method)
	}
	return apiV1, nil
}

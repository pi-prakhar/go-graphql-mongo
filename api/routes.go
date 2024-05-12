package api

import "github.com/gorilla/mux"

func Router() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/playground", gqlPlaygroundHandler())
	r.HandleFunc("/api/query", queryHandler())
	return r

}

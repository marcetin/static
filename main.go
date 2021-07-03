package main

import (
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	path := "/var/www/"
	r := mux.NewRouter()
	sites, err := ioutil.ReadDir(path)
	if err != nil {
		log.Fatal(err)
	}
	for _, f := range sites {
		n := f.Name()
		r.Host(n).Subrouter()
		r.StrictSlash(true)
		r.PathPrefix("/").Handler(http.StripPrefix("/", http.FileServer(http.Dir(path+n))))
	}
	log.Fatal(http.ListenAndServe(":80", handlers.CORS(handlers.AllowedOrigins([]string{"*"}))(r)))
}

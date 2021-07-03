package main

import (
	"fmt"
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
		fmt.Println("domain:", n)
		s := r.Host(n).Subrouter()
		s.StrictSlash(true)
		s.PathPrefix("/").Handler(http.StripPrefix("/", http.FileServer(http.Dir(path+n))))
		if n == "json.okno.rs" {
			s.Headers("Content-Type", "application/json")
		}
	}
	log.Fatal(http.ListenAndServe(":80", handlers.CORS(handlers.AllowedOrigins([]string{"*"}))(r)))
}

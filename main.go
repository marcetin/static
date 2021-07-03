package main

import (
	"crypto/tls"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/user"
	"path/filepath"

	"golang.org/x/crypto/acme/autocert"
)

func main() {
	// setup a simple handler which sends a HTHS header for six months (!)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Strict-Transport-Security", "max-age=15768000 ; includeSubDomains")
		fmt.Fprintf(w, "Hello, HTTPS world!")
	})

	// look for the domains to be served from command line args
	flag.Parse()
	domains := flag.Args()
	if len(domains) == 0 {
		log.Fatalf("fatal; specify domains as arguments")
	}

	// create the autocert.Manager with domains and path to the cache
	certManager := autocert.Manager{
		Prompt:     autocert.AcceptTOS,
		HostPolicy: autocert.HostWhitelist(domains...),
	}

	// optionally use a cache dir
	dir := cacheDir()
	if dir != "" {
		certManager.Cache = autocert.DirCache(dir)
	}

	// create the server itself
	server := &http.Server{
		Addr: ":https",
		TLSConfig: &tls.Config{
			GetCertificate: certManager.GetCertificate,
		},
	}

	log.Printf("Serving http/https for domains: %+v", domains)
	go func() {
		// serve HTTP, which will redirect automatically to HTTPS
		h := certManager.HTTPHandler(nil)
		log.Fatal(http.ListenAndServe(":http", h))
	}()

	// serve HTTPS!
	log.Fatal(server.ListenAndServeTLS("", ""))
}

// cacheDir makes a consistent cache directory inside /tmp. Returns "" on error.
func cacheDir() (dir string) {
	if u, _ := user.Current(); u != nil {
		dir = filepath.Join(os.TempDir(), "cache-golang-autocert-"+u.Username)
		if err := os.MkdirAll(dir, 0700); err == nil {
			return dir
		}
	}
	return ""
}

//package main
//
//import (
//	"crypto/tls"
//	"github.com/gorilla/handlers"
//	"github.com/gorilla/mux"
//	"golang.org/x/crypto/acme/autocert"
//	"log"
//	"net/http"
//	"time"
//)
//
//var (
//	path = "~/www"
//)
//
//func main() {
//	r := mux.NewRouter()
//	tlsman := autocert.Manager{
//		Prompt: autocert.AcceptTOS,
//		//HostPolicy: autocert.HostWhitelist("ws.okno.rs", "wss.okno.rs", "ns.okno.rs"),
//		Cache: autocert.DirCache(path),
//	}
//	www := &http.Server{
//		Handler:      handler(r),
//		Addr:         ":80",
//		WriteTimeout: 15 * time.Second,
//		ReadTimeout:  15 * time.Second,
//	}
//	wwwtls := www
//	wwwtls.Addr = ":443"
//	wwwtls.TLSConfig = &tls.Config{
//		GetCertificate: tlsman.GetCertificate,
//	}
//	//r.PathPrefix("/").Handler(http.StripPrefix("/", http.FileServer(http.Dir("/home/gorun/okno/templates/"))))
//	//log.Fatal(http.ListenAndServe(":80", handlers.CORS(handlers.AllowedOrigins([]string{"*"}))(r)))
//
//	//log.Fatal(http.ListenAndServeTLS("","",":80", handlers.CORS(handlers.AllowedOrigins([]string{"*"}))(r)))
//	//log.Fatal(www.ListenAndServe())
//
//	log.Fatal(wwwtls.ListenAndServeTLS("", ""))
//}
//
//func handler(r *mux.Router) http.Handler {
//	r.Host("{domain}").Subrouter().PathPrefix("/").Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
//		http.StripPrefix("/", http.FileServer(http.Dir(path+mux.Vars(r)["domain"])))
//	}))
//	//return handlers.CORS()(handlers.CompressHandler(interceptHandler(r, defaultErrorHandler)))
//	return handlers.CORS(handlers.AllowedOrigins([]string{"*"}))(r)
//}

//func defaultErrorHandler(w http.ResponseWriter, status int) {
//	//t := template.Must(template.ParseFiles("errors/error.html"))
//	//w.Header().Set("Content-Type", "text/html")
//	//t.Execute(w, map[string]interface{}{"status": status})
//	w.Header().Set("Content-Type", "text/html")
//	//tpl.TemplateHandler(cfg.Path).ExecuteTemplate(w, "error_gohtml", map[string]interface{}{"status": status})
//}
//func interceptHandler(next http.Handler, errH errorHandler) http.Handler {
//	if errH == nil {
//		errH = defaultErrorHandler
//	}
//	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
//		next.ServeHTTP(&interceptResponseWriter{w, errH}, r)
//	})
//}

//type errorHandler func(http.ResponseWriter, int)
//
//type interceptResponseWriter struct {
//	http.ResponseWriter
//	errH func(http.ResponseWriter, int)
//}

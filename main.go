package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

var (
	path = "/var/www/"
)

func main() {
	r := mux.NewRouter()
	//tlsman := autocert.Manager{
	//	Prompt:     autocert.AcceptTOS,
	//	//HostPolicy: autocert.HostWhitelist("ws.okno.rs", "wss.okno.rs", "ns.okno.rs"),
	//	Cache:      autocert.DirCache(path),
	//}
	//www := &http.Server{
	//	Handler:      handler(r),
	//	Addr:         ":80",
	//	WriteTimeout: 15 * time.Second,
	//	ReadTimeout:  15 * time.Second,
	//}
	//wwwtls := www
	//wwwtls.Addr = ":443"
	//wwwtls.TLSConfig =&tls.Config{
	//	GetCertificate: tlsman.GetCertificate,
	//}
	//r.PathPrefix("/").Handler(http.StripPrefix("/", http.FileServer(http.Dir("/home/gorun/okno/templates/"))))
	//log.Fatal(http.ListenAndServe(":80", handlers.CORS(handlers.AllowedOrigins([]string{"*"}))(r)))

	r.Host("{sub}.{domain}.{tld}").PathPrefix("/").Handler(h(r))
	log.Fatal(http.ListenAndServe(":80", nil))
	//log.Fatal(http.ListenAndServeTLS("","",":80", handlers.CORS(handlers.AllowedOrigins([]string{"*"}))(r)))
	//log.Fatal(www.ListenAndServe())
	//log.Fatal(wwwtls.ListenAndServeTLS("", ""))
}

func h(r *mux.Router) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		//v := mux.Vars(r)
		path = path + mux.Vars(r)["domain"] + "." + mux.Vars(r)["tld"]
		if mux.Vars(r)["sub"] != "" {
			path = path + mux.Vars(r)["sub"] + "." + mux.Vars(r)["domain"] + "." + mux.Vars(r)["tld"]
			fmt.Println("p221:", path)
		}
		fmt.Println("p1:", path)

		http.StripPrefix("/", http.FileServer(http.Dir(path)))
	})
	//return handlers.CORS()(handlers.CompressHandler(interceptHandler(r, defaultErrorHandler)))
	//return handlers.CORS(handlers.AllowedOrigins([]string{"*"}))(r)
	//return http.StripPrefix("/", http.FileServer(http.Dir(path)))
}

func defaultErrorHandler(w http.ResponseWriter, status int) {
	//t := template.Must(template.ParseFiles("errors/error.html"))
	//w.Header().Set("Content-Type", "text/html")
	//t.Execute(w, map[string]interface{}{"status": status})
	w.Header().Set("Content-Type", "text/html")
	//tpl.TemplateHandler(cfg.Path).ExecuteTemplate(w, "error_gohtml", map[string]interface{}{"status": status})
}
func interceptHandler(next http.Handler, errH errorHandler) http.Handler {
	if errH == nil {
		errH = defaultErrorHandler
	}
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		next.ServeHTTP(&interceptResponseWriter{w, errH}, r)
	})
}

type errorHandler func(http.ResponseWriter, int)

type interceptResponseWriter struct {
	http.ResponseWriter
	errH func(http.ResponseWriter, int)
}

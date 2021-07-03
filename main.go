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

//	//tlsman := autocert.Manager{
//	//	Prompt:     autocert.AcceptTOS,
//	//	//HostPolicy: autocert.HostWhitelist("ws.okno.rs", "wss.okno.rs", "ns.okno.rs"),
//	//	Cache:      autocert.DirCache(path),
//	//}
//	//www := &http.Server{
//	//	Handler:      handler(r),
//	//	Addr:         ":80",
//	//	WriteTimeout: 15 * time.Second,
//	//	ReadTimeout:  15 * time.Second,
//	//}
//	//wwwtls := www
//	//wwwtls.Addr = ":443"
//	//wwwtls.TLSConfig =&tls.Config{
//	//	GetCertificate: tlsman.GetCertificate,
//	//}
//	//r.PathPrefix("/").Handler(http.StripPrefix("/", http.FileServer(http.Dir("/home/gorun/okno/templates/"))))
//	//log.Fatal(http.ListenAndServe(":80", handlers.CORS(handlers.AllowedOrigins([]string{"*"}))(r)))
//
//	//r.Host("{sub}.{domain}.{tld}").PathPrefix("/").Handler(http.HandlerFunc(h))
//	//log.Fatal(http.ListenAndServe(":80", nil))
//	//log.Fatal(http.ListenAndServeTLS("","",":80", handlers.CORS(handlers.AllowedOrigins([]string{"*"}))(r)))
//	//log.Fatal(www.ListenAndServe())
//	//log.Fatal(wwwtls.ListenAndServeTLS("", ""))
//}
//
//func h(w http.ResponseWriter, r *http.Request) {
//	//return (w http.ResponseWriter, r *http.Request){
//	//return func(w http.ResponseWriter, r *http.Request) {
//
//		//path = path + mux.Vars(r)["domain"] + "." + mux.Vars(r)["tld"]
//		//if mux.Vars(r)["sub"] != "" {
//		//	path = path + mux.Vars(r)["sub"] + "." + mux.Vars(r)["domain"] + "." + mux.Vars(r)["tld"]
//		//	fmt.Println("p221:", path)
//		//}
//		//fmt.Println("p1:", path)
//		//
//		//http.StripPrefix("/", http.FileServer(http.Dir(path)))
//		//})
//		//return handlers.CORS()(handlers.CompressHandler(interceptHandler(r, defaultErrorHandler)))
//		//return handlers.CORS(handlers.AllowedOrigins([]string{"*"}))(r)
//		//return http.StripPrefix("/", http.FileServer(http.Dir(path)))
//	//})
//}
//
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
//
//type errorHandler func(http.ResponseWriter, int)
//
//type interceptResponseWriter struct {
//	http.ResponseWriter
//	errH func(http.ResponseWriter, int)
//}

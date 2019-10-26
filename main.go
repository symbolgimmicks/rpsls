package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/symbolgimmicks/rpsls/gameservicerouter"
)

func handleRequests() {
	router := mux.NewRouter().StrictSlash(true)

	//https://www.alexedwards.net/blog/serving-static-sites-with-go
	// JJB - Let the home page be the test page?  Sure, for now.
	// that approach didn't work with gorilla (need to learn the basics later though...)
	// So instead trying this approach...
	//https://github.com/gorilla/mux#static-files
	var dir string
	flag.StringVar(&dir, "dir", ".", "the directory to serve files from. Defaults to the current dir")
	flag.Parse()
	router.PathPrefix("/Sandbox/").Handler(http.StripPrefix("/Sandbox/", http.FileServer(http.Dir(dir))))

	router.HandleFunc("/choices/{id:[0-9]+}", gameservicerouter.OnGetSingleChoice)
	router.HandleFunc("/choice", gameservicerouter.OnGetRandomChoice).Methods("GET")
	router.HandleFunc("/choices", gameservicerouter.OnGetChoices).Methods("GET")
	router.HandleFunc("/play", gameservicerouter.OnPlay).Methods("POST")

	//https: //www.thepolyglotdeveloper.com/2017/10/handling-cors-golang-web-application/
	//https://golang.org/pkg/net/http/
	//https://gist.github.com/denji/12b3a568f092ab951456
	//https://stackoverflow.com/questions/50625283/how-to-install-openssl-in-windows-10
	//https://golangcode.com/basic-https-server-with-certificate/
	log.Fatal(http.ListenAndServeTLS(":4077", "rpsls-server.crt", "rpsls-server.key", handlers.CORS(handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"}), handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS"}), handlers.AllowedOrigins([]string{"*"}))(router)))
}

func main() {
	handleRequests()
}

package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/symbolgimmicks/rpsls/gameservicerouter"
)

func onHome(w http.ResponseWriter, r *http.Request) {
	// Setup the body
	fmt.Fprintf(w, "JJB")
}

func handleRequests() {
	router := mux.NewRouter().StrictSlash(true)

	// router.HandleFunc("/", onHome) // Not needed... or can A UI live here...?

	router.HandleFunc("/choices/{id:[0-9]+}", gameservicerouter.OnGetChoices)

	router.HandleFunc("/choice", gameservicerouter.OnGetRandomChoice).Methods("GET")
	router.HandleFunc("/choices", gameservicerouter.OnGetChoices).Methods("GET")

	router.HandleFunc("/play", gameservicerouter.OnPlay).Methods("POST")

	//https: //www.thepolyglotdeveloper.com/2017/10/handling-cors-golang-web-application/
	// JJB - I do not understand CORS. I just understand it was in my way of testing the server and host on the same machine (I think...)
	// Yes.  That's M*A*S*H*
	log.Fatal(http.ListenAndServe(":4077", handlers.CORS(handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"}), handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS"}), handlers.AllowedOrigins([]string{"*"}))(router)))

	// JJB - Learn this [https://github.com/gorilla/mux#handling-cors-requests] because it explains how to do CORS without that handlers lib.
}

func main() {
	handleRequests()
}

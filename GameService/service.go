package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"runtime"
	"strconv"

	handlers "github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

// Choice - Game options
type Choice struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// PlayerChoice - for deseraliazing player post data.
type PlayerChoice struct {
	Player int `json:"player"`
}

// Game - object - in case of multiple
type Game struct {
	ID    int
	Title string
}

// Result - When playing, what happened?
type Result struct {
	Results  string `json:"results"`
	Player   int    `json:"player"`
	Computer int    `json:"computer"`
}

// Choices - user can only select these.
var Choices = []Choice{
	{1, "Rock"},
	{2, "Paper"},
	{3, "Scissors"},
	{4, "Lizard"},
	{5, "Spock"},
}

func consoleLogEndpointHit(funcName string) {
	//https://stackoverflow.com/questions/10742749/get-name-of-function-using-reflection-in-golang/41672632
	// Show the name of where we are in the io - intended for debug but how?
	fmt.Println("Endpoint Hit: " + funcName)
}

func onHome(w http.ResponseWriter, r *http.Request) {
	pc, _, _, _ := runtime.Caller(0)
	consoleLogEndpointHit(runtime.FuncForPC(pc).Name())

	// Setup the body
	fmt.Fprintf(w, "Welcome to the HomePage!")
}

func onGetChoices(w http.ResponseWriter, r *http.Request) {
	pc, _, _, _ := runtime.Caller(0)
	consoleLogEndpointHit(runtime.FuncForPC(pc).Name())

	json.NewEncoder(w).Encode(Choices)
}

func onGetSingleChoice(w http.ResponseWriter, r *http.Request) {
	pc, _, _, _ := runtime.Caller(0)
	consoleLogEndpointHit(runtime.FuncForPC(pc).Name())

	vars := mux.Vars(r)
	sKey := vars["id"]
	iKey, err := strconv.ParseInt(sKey, 10, 32)

	if err != nil {
		fmt.Fprintf(w, "Key: "+sKey)
		fmt.Fprintf(w, "<!-- "+err.Error()+" -->")
	}
	for _, item := range Choices {
		if item.ID == int(iKey) {
			json.NewEncoder(w).Encode(item)
		}
	}
}

func onPlay(w http.ResponseWriter, r *http.Request) {
	pc, _, _, _ := runtime.Caller(0)
	consoleLogEndpointHit(runtime.FuncForPC(pc).Name())

	reqBody, _ := ioutil.ReadAll(r.Body)
	var playerChoice PlayerChoice
	json.Unmarshal(reqBody, &playerChoice)

	// TODO: who wins...
	results := Result{"win", playerChoice.Player, 2}
	json.NewEncoder(w).Encode(results)
}

func onGetRandomChoice(w http.ResponseWriter, r *http.Request) {
	pc, _, _, _ := runtime.Caller(0)
	fmt.Println("Endpoint Hit: " + runtime.FuncForPC(pc).Name())

	var response Choice
	response = Choices[3]

	json.NewEncoder(w).Encode(response)
}

func handleRequests() {
	pc, _, _, _ := runtime.Caller(0)
	consoleLogEndpointHit(runtime.FuncForPC(pc).Name())

	router := mux.NewRouter().StrictSlash(true)

	// router.HandleFunc("/", onHome) // Not needed... or can A UI live here...?

	router.HandleFunc("/choices/{id:[0-9]+}", onGetSingleChoice)

	router.HandleFunc("/choice", onGetRandomChoice).Methods("GET")
	router.HandleFunc("/choices", onGetChoices).Methods("GET")

	router.HandleFunc("/play", onPlay).Methods("POST")

	//https: //www.thepolyglotdeveloper.com/2017/10/handling-cors-golang-web-application/
	// JJB - I do not understand CORS. I just understand it was in my way of testing the server and host on the same machine (I think...)
	// Yes.  That's M*A*S*H*
	log.Fatal(http.ListenAndServe(":4077", handlers.CORS(handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"}), handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS"}), handlers.AllowedOrigins([]string{"*"}))(router)))

}

func main() {
	handleRequests()
}

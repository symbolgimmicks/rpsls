package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"runtime"
	"strconv"
	"time"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

//https://stackoverflow.com/questions/17156371/how-to-get-json-response-from-http-get
// JJB - default client does not have a timeout so you probably don't want to use it on purpose.
var myClient = &http.Client{Timeout: 10 * time.Second}

// RandomNumber - expected response from [https://codechallenge.boohma.com/random]
type RandomNumber struct {
	RandomNumber int `json:"random_number"`
}

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

// Choices - users can only select these.
// JJB - Are these secure here?  Can they change?
var Choices = []Choice{
	{1, "Rock"},
	{2, "Paper"},
	{3, "Scissors"},
	{4, "Lizard"},
	{5, "Spock"},
}

// Is lumping 3rd party services into packages a practical thing in Go?
func rollRandomNumber() int {
	pc, _, _, _ := runtime.Caller(0)
	consoleLogEndpointHit(runtime.FuncForPC(pc).Name())

	data := &RandomNumber{}

	getResponse, err := myClient.Get("https://codechallenge.boohma.com/random")
	if err != nil {
		fmt.Println("err: ", err.Error())
		return -1
	}
	defer getResponse.Body.Close()

	//https://stackoverflow.com/questions/26744873/converting-map-to-struct

	//go get github.com/mitchellh/mapstructure - is this required?
	err2 := json.NewDecoder(getResponse.Body).Decode(&data)
	if err2 != nil {
		fmt.Println("err2: ", err2.Error())
		return -2
	}

	return data.RandomNumber
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
	fmt.Fprintf(w, "JJB")
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

	var roll int = rollRandomNumber()

	fmt.Println("data: ", roll)

	var computerChoice Choice = rollToChoice(roll)

	// TODO: figure out who actually wins...

	results := Result{"win", playerChoice.Player, computerChoice.ID}
	json.NewEncoder(w).Encode(results)
}

func rollToChoice(roll int) Choice {
	switch {
	case roll > 80:
		return Choices[0]
	case roll > 60:
		return Choices[1]
	case roll > 40:
		return Choices[2]
	case roll > 20:
		return Choices[3]
	default:
		return Choices[4]
	}
}
func onGetRandomChoice(w http.ResponseWriter, r *http.Request) {
	pc, _, _, _ := runtime.Caller(0)
	consoleLogEndpointHit(runtime.FuncForPC(pc).Name())

	var roll int = rollRandomNumber()

	fmt.Println("data: ", roll)

	var response Choice = rollToChoice(roll)

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

	// JJB - Learn this [https://github.com/gorilla/mux#handling-cors-requests] because it explains how to do CORS without that handlers lib.
}

func main() {
	handleRequests()
}

package gameservicerouter

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/symbolgimmicks/rpsls/choice"
)

// PlayerChoice - for deseraliazing player post data.
type PlayerChoice struct {
	Player int `json:"player"`
}

// Result - When playing, what happened?
type Result struct {
	Results  string `json:"results"`
	Player   int    `json:"player"`
	Computer int    `json:"computer"`
}

// OnGetChoices - returns all Choices
func OnGetChoices(w http.ResponseWriter, r *http.Request) {
	var results = choice.ValidChoices()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(&results)
}

// OnGetSingleChoice - returns the value of a given choice ID
func OnGetSingleChoice(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	sKey := vars["id"]
	iKey, err := strconv.ParseInt(sKey, 10, 32)
	var found bool = false
	if err != nil {
		fmt.Fprintf(w, "Key: "+sKey)
		fmt.Fprintf(w, "<!-- "+err.Error()+" -->")
	}

	w.Header().Set("Content-Type", "application/json")

	for _, item := range choice.Choices {
		if item.ID == int(iKey) {
			json.NewEncoder(w).Encode(item)
			found = true
			break
		}
	}
	if found == false {
		json.NewEncoder(w).Encode(choice.Choices[0])
	}
}

// OnPlay - plays a round.
func OnPlay(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)
	var playerChoice PlayerChoice
	json.Unmarshal(reqBody, &playerChoice)
	var computerChoice choice.Choice = choice.GenerateRandom()

	var roundResult int = choice.Evaluate(choice.Choices[playerChoice.Player], computerChoice)
	var strResult string
	if roundResult == -1 {
		strResult = "win"
	} else if roundResult == 0 {
		strResult = "tie"
	} else {
		strResult = "lose"
	}

	results := Result{strResult, playerChoice.Player, computerChoice.ID}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(results)
}

// OnGetRandomChoice - Picks a choice
func OnGetRandomChoice(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(choice.GenerateRandom())
}

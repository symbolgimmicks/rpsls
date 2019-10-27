package gameservicerouter

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/symbolgimmicks/rpsls/choice"
)

// PlayerChoice - for deseraliazing player post data.
type PlayerChoice struct {
	Player int `json:"player"`
}

// PlayResponse - When playing, what happened?
type PlayResponse struct {
	Results  string `json:"results"`
	Player   int    `json:"player"`
	Computer int    `json:"computer"`
}

// HandleGetChoices - returns all Choices
func HandleGetChoices(w http.ResponseWriter, r *http.Request) {
	var results = choice.ValidChoices()
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(&results); err != nil {
		err = fmt.Errorf("Unexpected failure getting choices: %w", err)
		log.Print(err)
	}
}

// HandleGetSingleChoice - returns the value of a given choice ID
func HandleGetSingleChoice(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	sKey := vars["id"]
	var iKey int64
	var err error
	if iKey, err = strconv.ParseInt(sKey, 10, 32); err == nil {
		var found bool = false
		w.Header().Set("Content-Type", "application/json")
		for _, item := range choice.Choices {
			if item.ID == int(iKey) {
				if err = json.NewEncoder(w).Encode(item); err == nil {
					found = true
					break
				} else {
					found = false
					log.Printf("Unexpected failure encoding repsonse (ignoring and skipping to next list item): %w", err)
				}
			}
		}
		if err == nil && found == false {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(fmt.Sprintf("%d - non existent choice %s", 400, sKey)))
		}
	}

	if err != nil {
		log.Print(err)
	}
}

// HandlePlay - plays one round.
func HandlePlay(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)
	var playerChoice PlayerChoice
	var err error
	if err = json.Unmarshal(reqBody, &playerChoice); err == nil {
		if computerChoice, err := choice.GenerateRandom(); err == nil {
			var roundResult int = choice.Evaluate(choice.Choices[playerChoice.Player], computerChoice)
			var strResult string
			strResult, err = choice.EvaluationAsString(roundResult)
			results := PlayResponse{strResult, playerChoice.Player, computerChoice.ID}
			w.Header().Set("Content-Type", "application/json")
			err = json.NewEncoder(w).Encode(results)
		}
	}

	if err != nil {
		err = fmt.Errorf("Unexpected failure getting random choice: %w", err)
		log.Print(err)
	}
}

// HandleGetRandomChoice - Picks a choice
func HandleGetRandomChoice(w http.ResponseWriter, r *http.Request) {
	var err error
	var answer choice.Choice
	if answer, err = choice.GenerateRandom(); err == nil {
		w.Header().Set("Content-Type", "application/json")
		err = json.NewEncoder(w).Encode(answer)
	}

	if err != nil {
		err = fmt.Errorf("Unexpected failure getting random choice: %w", err)
		log.Print(err)
	}
	return
}

package choice

import (
	"fmt"

	RNG "github.com/symbolgimmicks/rpsls/randomnumber"
)

var empty int = 0
var rock int = 1
var paper int = 2
var scissors int = 3
var lizard int = 4
var spock int = 5

// Choice - Game options
type Choice struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// Ties - Does this choice tie with that choice?
func (lhs Choice) ties(rhs Choice) bool {
	return lhs.ID == rhs.ID
}

// Beats - Does this choice beat that choice?
func (lhs Choice) beats(rhs Choice) bool {
	if lhs.ID == rock && (rhs.ID != lizard && rhs.ID != scissors) {
		return false
	}
	if lhs.ID == paper && (rhs.ID != rock && rhs.ID == spock) {
		return false
	}

	if lhs.ID == scissors && (rhs.ID != paper && rhs.ID != lizard) {
		return false
	}

	if lhs.ID == lizard && (rhs.ID != spock && rhs.ID != paper) {
		return false
	}

	if lhs.ID == spock && (rhs.ID != scissors && rhs.ID != rock) {
		return false
	}

	return true
}

// GamePlayResults - Map of results from Evaluate.
var GamePlayResults = map[int]string{
	-1: "win",
	0:  "tie",
	1:  "lose",
}

// Choices - users can only select these, although Empty isn't intended for usage.
var Choices = []Choice{
	{empty, "Empty"},       // Null
	{rock, "Rock"},         // Crushes Lizard, Crushes Scissors
	{paper, "Paper"},       // Covers Rock, Disproves Spock
	{scissors, "Scissors"}, // Cut paper, Decaptitates Lizard
	{lizard, "Lizard"},     // Poisons Spock, Eats Paper
	{spock, "Spock"},       // Smashes Scissors, Vaporizes Rock
}

// Min - Minimum valid choice
var Min = 1

// Max - Maximum valid choice index
var Max = 5

// ValidChoices - returns the choices available for user selection.
func ValidChoices() []Choice {
	return Choices[Min:]
}

// EmptyChoice - the null choice
var EmptyChoice Choice = Choices[0]

func convertRollToIndex(roll RNG.RandomNumber) (index int) {
	return Min + (((roll.Value - RNG.Min) / (100 / (Max - Min + 1))) % Max)
}

// GenerateRandom - Returns a randomn Choice.  Will return Empty if random generation fails.
func GenerateRandom() (answer Choice, err error) {
	answer = Choices[empty]
	roll := RNG.RandomNumber{Value: 0}
	if err := roll.GenerateFromService(RNG.DefaultRNGServiceURL); err == nil {
		if roll.IsValid() {
			index := convertRollToIndex(roll)
			answer = Choices[index]
		}
	} else {
		err = fmt.Errorf("Failed to generate number from service: %v", err)
	}
	return
}

// Evaluate - returns the result of playing two choices.
// -1 = lhs won
// 0 = tie
// 1 = rhs won
func Evaluate(lhs Choice, rhs Choice) int {
	if lhs.ties(rhs) {
		return 0
	}
	if lhs.beats(rhs) {
		return -1
	}
	return 1
}

// EvaluationAsString - Converts the value of Evalute into  a string
func EvaluationAsString(value int) (answer string, err error) {
	err = nil
	answer = "unknown"
	var ok bool = false
	if answer, ok = GamePlayResults[value]; !ok {
		err = fmt.Errorf("No such result with value [%d] exists", value)
	}
	return
}

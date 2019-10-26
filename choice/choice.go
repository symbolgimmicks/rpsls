package choice

import (
	rand "github.com/symbolgimmicks/rpsls/randomnumber"
)

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

// Choices - users can only select these, although Empty isn't intended for usage.
var Choices = []Choice{
	{0, "Empty"},           // Null
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

// GenerateRandom - Returns a randomn Choice.  Will return Empty if random generation fails.
func GenerateRandom() Choice {
	var roll int = rand.Generate()
	var index int = Min + (((roll - rand.Min) / (100 / (Max - Min + 1))) % Max)
	return Choices[index]
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

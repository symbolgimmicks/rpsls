package choice

import rand "github.com/symbolgimmicks/rpsls/randomnumber"

// Choice - Game options
type Choice struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// Choices - users can only select these, although Empty isn't intended for usage.
var Choices = []Choice{
	{0, "Empty"},
	{1, "Rock"},
	{2, "Paper"},
	{3, "Scissors"},
	{4, "Lizard"},
	{5, "Spock"},
}

// Min - Minimum valid choice
var Min = 1

// Max - Maximum valid choice index
var Max = 5

// GenerateRandom - Returns a randomn Choice.  Will return Empty if random generation fails.
func GenerateRandom() Choice {
	var index int = Min + ((rand.Generate()-rand.Min)/(rand.Max-rand.Min+1)/(Max-Min+1))%Max
	return Choices[index]
}

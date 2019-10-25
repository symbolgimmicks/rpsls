package choice

// Choice - Game options
type Choice struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
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

func IntToChoice(roll int) Choice {
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

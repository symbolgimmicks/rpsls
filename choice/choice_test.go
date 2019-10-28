package choice_test

import (
	"errors"
	"fmt"
	"log"
	"net/http/httptest"
	"strconv"

	"github.com/DATA-DOG/godog"
	"github.com/DATA-DOG/godog/gherkin"
	"github.com/symbolgimmicks/rpsls/choice"
)

type validateFeature struct {
	resp             *httptest.ResponseRecorder
	lastValidChoices []choice.Choice
	lastGameResult   int
	activeChoice     choice.Choice
}

func (a *validateFeature) reset() {
	a.lastGameResult = -9999
	a.resp = httptest.NewRecorder()
}

func (a *validateFeature) thereAreTheseChoices(testData *gherkin.DataTable) (err error) {
	err = nil
	for _, row := range testData.Rows {
		if id, err := strconv.Atoi(row.Cells[0].Value); err == nil {
			if name := row.Cells[1].Value; choice.Choices[id].Name != name {
				err = errors.New(fmt.Sprintf("Choice not found: %s", name))
			}
		}
	}
	return
}

func (a *validateFeature) callValidChoices() (err error) {
	err = nil
	a.lastValidChoices = choice.ValidChoices()
	return
}

//func (a *validateFeature) theResponseShouldMatchJson(body *gherkin.DocString) (err error) {
func (a *validateFeature) theResponseShouldMatchJson(body *gherkin.DocString) (err error) {
	return errors.New("Not implemented")
}

func (a *validateFeature) theResponseShouldMatchString(content string) (err error) {
	if actual := fmt.Sprintf("%v", a.lastValidChoices); content != actual {
		err = errors.New(fmt.Sprintf("%v != %v", content, actual))
	}

	return
}

func getChoiceByName(name string) (answer choice.Choice, err error) {
	err = nil
	answer = choice.EmptyChoice

	for _, next := range choice.ValidChoices() {
		log.Printf("Next item: %v", next)
		if answer == choice.EmptyChoice && next.Name == name {
			answer = next
			return
		}
	}
	return
}

func (a *validateFeature) lhsPlaysRhs(lhsName string, rhsName string) (err error) {
	var lhs choice.Choice
	var rhs choice.Choice
	if lhs, err = choice.NewByString(lhsName); err != nil {
		err = fmt.Errorf("LHS choice not found (%v)", err)
	} else if rhs, err = choice.NewByString(rhsName); err != nil {
		err = fmt.Errorf("RHS choice not found (%v)", err)
	} else if a.lastGameResult, err = lhs.Play(rhs); err != nil {
		err = fmt.Errorf("Unexpected failure during game play (%v)", err)
	} else {
		log.Printf("%v vs %v => %d", lhs, rhs, a.lastGameResult)
	}

	return
}

func (a *validateFeature) lastPlayResult(expected string) (err error) {
	err = nil
	var actual string

	if actual, err = choice.EvaluationAsString(a.lastGameResult); err == nil {
		if actual != expected {
			err = errors.New(fmt.Sprintf("Expected [%s]. Actual: [%s (%d)]", expected, actual, a.lastGameResult))
		}
	} else {
		err = fmt.Errorf("Unexpected failure converting Evaluation result to string (%v)", err)
	}
	return
}

func (a *validateFeature) setActiveChoice(choice string) (err error) {
	if a.activeChoice, err = getChoiceByName(choice); err != nil {
		err = fmt.Errorf("Failed to set active choice (%v)", err)
	}
	return
}

func (a *validateFeature) playActiveChoiceAgainst(targetName string) (err error) {
	err = nil
	if target, err := choice.NewByString(targetName); err != nil {
		a.lastGameResult = -9999
		err = fmt.Errorf("Unexpected failure during gameplay (%v)", err)
	} else if a.lastGameResult, err = a.activeChoice.Play(target); err != nil {
		a.lastGameResult = -9999
		err = fmt.Errorf("Unexpected failure during gameplay (%v)", err)
	} else {
		log.Printf("[%v] PLAYS [%v]!!", a.activeChoice, target)
		log.Printf("LAST GAME RESULT [%d]", a.lastGameResult)
	}
	return
}

func FeatureContext(s *godog.Suite) {
	validateApi := &validateFeature{}

	s.BeforeScenario(func(interface{}) {
		validateApi.reset()
	})

	s.Step(`^there are these choices:$`, validateApi.thereAreTheseChoices)
	s.Step(`^I call ValidChoices$`, validateApi.callValidChoices)
	s.Step(`^the response should match the string: \"([^"]*)\"$`, validateApi.theResponseShouldMatchString)
	s.Step(`^\"([^"]*)\" plays against \"([^"]*)\"$`, validateApi.lhsPlaysRhs)
	s.Step(`^the play result is \"([^"]*)\"$`, validateApi.lastPlayResult)
	s.Step(`^I set the active choice to \"([^"]*)\"$`, validateApi.setActiveChoice)
	s.Step(`^I play \"([^"]*)\"$`, validateApi.playActiveChoiceAgainst)
}

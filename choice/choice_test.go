package choice_test

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http/httptest"
	"reflect"
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

func (a *validateFeature) theResponseShouldMatchJson(body *gherkin.DocString) (err error) {
	var expected, actual interface{}

	// re-encode expected response
	// if err = json.Unmarshal([]byte(body.Content), &expected); err != nil {
	// 	err = fmt.Errorf("Failed to unmarshal the expected JSON data: %v", err)
	// } else
	if err = json.Unmarshal(a.resp.Body.Bytes(), &actual); err != nil {
		// re-encode actual response too
		err = fmt.Errorf("Failed to unmarshal the actual JSON data: %v", err)
	}

	log.Printf("expected JSON: %v\r\nactual JSON: %v\r\n", expected, actual)

	if !reflect.DeepEqual(expected, actual) {
		// the matching may be adapted per different requirements.
		err = errors.New(fmt.Sprintf("expected JSON does not match actual, %v vs. %v", expected, actual))
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
	err = nil

	var lhsChoice = choice.EmptyChoice
	lhsChoice, err = getChoiceByName(lhsName)

	if lhsChoice == choice.EmptyChoice {
		err = errors.New(fmt.Sprintf("LHS choice [%s] not found", lhsName))
		return
	}

	var rhsChoice = choice.EmptyChoice
	rhsChoice, err = getChoiceByName(rhsName)
	if rhsChoice == choice.EmptyChoice {
		err = errors.New(fmt.Sprintf("RHS choice [%s] not found", rhsName))
		return
	}

	a.lastGameResult = choice.Evaluate(lhsChoice, rhsChoice)

	log.Printf("%v vs %v => %d", lhsChoice, rhsChoice, a.lastGameResult)

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
	var target choice.Choice

	if target, err = getChoiceByName(targetName); err == nil {
		a.lastGameResult = choice.Evaluate(a.activeChoice, target)
	} else {
		a.lastGameResult = -9999
	}
	log.Printf("[%v] PLAYS [%v]!!", a.activeChoice, target)

	log.Printf("LAST GAME RESULT [%d]", a.lastGameResult)
	return
}

func FeatureContext(s *godog.Suite) {
	validateApi := &validateFeature{}

	s.BeforeScenario(func(interface{}) {
		validateApi.reset()
	})

	s.Step(`^there are these choices:$`, validateApi.thereAreTheseChoices)
	s.Step(`^I call ValidChoices$`, validateApi.callValidChoices)
	s.Step(`^the response should match:$`, validateApi.theResponseShouldMatchJson)
	s.Step(`^\"([^"]*)\" plays against \"([^"]*)\"$`, validateApi.lhsPlaysRhs)
	s.Step(`^the play result is \"([^"]*)\"$`, validateApi.lastPlayResult)
	s.Step(`^I set the active choice to \"([^"]*)\"$`, validateApi.setActiveChoice)
	s.Step(`^I play \"([^"]*)\"$`, validateApi.playActiveChoiceAgainst)
}

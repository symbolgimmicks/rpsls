package main_test

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"reflect"
	"strconv"
	"time"

	"github.com/DATA-DOG/godog"
	"github.com/DATA-DOG/godog/gherkin"
	"github.com/symbolgimmicks/rpsls/choice"
	RNG "github.com/symbolgimmicks/rpsls/randomnumber"
)

type choiceFeatures struct {
	resp             *httptest.ResponseRecorder
	lastValidChoices []choice.Choice
	lastGameResult   int
	activeChoice     choice.Choice
}

func (a *choiceFeatures) reset() {
	a.lastGameResult = -9999
	a.resp = httptest.NewRecorder()
}

func (a *choiceFeatures) thereAreTheseChoices(testData *gherkin.DataTable) (err error) {
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

func (a *choiceFeatures) callValidChoices() (err error) {
	err = nil
	a.lastValidChoices = choice.ValidChoices()
	return
}

func (a *choiceFeatures) theResponseShouldMatchString(content string) (err error) {
	if actual := fmt.Sprintf("%v", a.lastValidChoices); content != actual {
		err = errors.New(fmt.Sprintf("%v != %v", content, actual))
	}

	return
}

func getChoiceByName(name string) (answer choice.Choice, err error) {
	err = nil
	answer = choice.EmptyChoice
	for _, next := range choice.ValidChoices() {
		if answer == choice.EmptyChoice && next.Name == name {
			answer = next
			return
		}
	}
	return
}

func (a *choiceFeatures) lhsPlaysRhs(lhsName string, rhsName string) (err error) {
	var lhs choice.Choice
	var rhs choice.Choice
	if lhs, err = choice.NewByString(lhsName); err != nil {
		err = fmt.Errorf("LHS choice not found (%v)", err)
	} else if rhs, err = choice.NewByString(rhsName); err != nil {
		err = fmt.Errorf("RHS choice not found (%v)", err)
	} else if a.lastGameResult, err = lhs.Play(rhs); err != nil {
		err = fmt.Errorf("Unexpected failure during game play (%v)", err)
	}

	return
}

func (a *choiceFeatures) lastPlayResult(expected string) (err error) {
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

func (a *choiceFeatures) setActiveChoice(choice string) (err error) {
	if a.activeChoice, err = getChoiceByName(choice); err != nil {
		err = fmt.Errorf("Failed to set active choice (%v)", err)
	}
	return
}

func (a *choiceFeatures) playActiveChoiceAgainst(targetName string) (err error) {
	err = nil
	if target, err := choice.NewByString(targetName); err != nil {
		a.lastGameResult = -9999
		err = fmt.Errorf("Unexpected failure during gameplay (%v)", err)
	} else if a.lastGameResult, err = a.activeChoice.Play(target); err != nil {
		a.lastGameResult = -9999
		err = fmt.Errorf("Unexpected failure during gameplay (%v)", err)
	}
	return
}

// GameServiceRouter
type gameServiceFeatures struct {
	resp                   *httptest.ResponseRecorder
	targetUrl              string
	jsonData               []byte
	lastChoiceReceived     choice.Choice
	lastChoiceListReceived []choice.Choice
	myClient               *http.Client
	nextId                 int
	lastGameResult         []byte
}

func (a *gameServiceFeatures) reset() {
	a.resp = httptest.NewRecorder()
	a.lastChoiceReceived, _ = choice.NewByID(0)
	a.targetUrl = ""
	a.nextId = 0

	// https://stackoverflow.com/questions/29164375/correct-way-to-initialize-empty-slice
	a.lastGameResult = nil
	a.lastChoiceListReceived = nil

	// JJB - Apparently, signing is a problem so learn this later
	// but hack it for now?
	//https://github.com/andygrunwald/go-jira/issues/52
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	a.myClient = &http.Client{Timeout: 15 * time.Second, Transport: tr}
}

func (a *gameServiceFeatures) setUrl(url string) (err error) {
	err = nil
	a.targetUrl = url
	return
}

func (a *gameServiceFeatures) setNextID(ID int) (err error) {
	err = nil
	a.nextId = ID
	return
}

func (a *gameServiceFeatures) sendGet(endpoint string, v interface{}) (resp *http.Response, err error) {
	var url string = a.targetUrl + endpoint
	fmt.Println("Sending request to [" + url + "]")
	if resp, err := a.myClient.Get(url); err != nil {
	} else {
		defer func() {
			switch t := recover().(type) {
			case string:
				err = fmt.Errorf(t)
			case error:
				err = t
			}
		}()

		defer resp.Body.Close()
		if err = json.NewDecoder(resp.Body).Decode(&v); err != nil {
			fmt.Println(fmt.Errorf("Failed to decode response (%v)", err))
		} else {
			fmt.Println(fmt.Sprintf("Received: [%v]", v))
		}
	}

	return
}

type choiceResultCollection struct {
	ids   map[string]string
	names map[string]string
}

func (a *gameServiceFeatures) sendGetToEndpoint(endpoint string) (err error) {
	err = nil
	switch endpoint {
	case "/choice":
		if resp, err := a.sendGet(endpoint, &a.lastChoiceReceived); err != nil {
			fmt.Printf(fmt.Sprintf("Failed to get a response (%v): %v", err.Error(), resp))
		}
	case "/choices/":
		if resp, err := a.sendGet(endpoint+strconv.Itoa(a.nextId), &a.lastChoiceReceived); err != nil {
			fmt.Printf(fmt.Sprintf("Failed to get a response (%v): %v", err.Error(), resp))
		}
	}

	return
}

func (a *gameServiceFeatures) sendPostToEndpoint(endpoint string) (err error) {
	//https://stackoverflow.com/questions/24455147/how-do-i-send-a-json-string-in-a-post-request-in-go
	var url string = a.targetUrl + endpoint
	var payload = []byte(fmt.Sprintf("{\"player\":%d}", a.lastChoiceReceived.ID))
	// json.NewEncoder(w).Encode(payload)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(payload))
	req.Header.Set("X-Custom-Header", "myvalue")
	req.Header.Set("Content-Type", "application/json")
	if resp, err := a.myClient.Do(req); err != nil {
		err = fmt.Errorf("Unexpected failure posting to play (%v)", err)
	} else {
		defer resp.Body.Close()
		if a.lastGameResult, err = ioutil.ReadAll(resp.Body); err != nil {
			err = fmt.Errorf("Unexpected failure posting to play (%v).  Response: %v", err, a.lastGameResult)
		}
	}

	return
}

func (a *gameServiceFeatures) aGameResultWasReturned() (err error) {
	err = nil
	if a.lastGameResult == nil {
		err = errors.New("No game result has been received")
	}
	return
}

func (a *gameServiceFeatures) anyChoiceExcept(except string) (err error) {
	var filter choice.Choice
	if filter, err = choice.NewByString(except); err != nil {
		err = fmt.Errorf("Unexpected failure testing game result (%v)", err)
	} else if reflect.DeepEqual(a.lastChoiceReceived, filter) {
		err = fmt.Errorf("[%v] == [%v]", a.lastChoiceReceived, filter)
	}
	return
}

func (a *gameServiceFeatures) choiceHasID() (err error) {
	if a.lastChoiceReceived.ID != a.nextId {
		err = errors.New(fmt.Sprintf("[%d] != [%d]", a.lastChoiceReceived.ID, a.nextId))
	}
	return
}

func (a *gameServiceFeatures) choiceHasName(name string) (err error) {
	if a.lastChoiceReceived.Name != name {
		err = errors.New(fmt.Sprintf("[%s] != [%s]", a.lastChoiceReceived.Name, name))
	}
	return
}

// Random Number Generator
type rngFeature struct {
	num               RNG.RandomNumber
	lastIsValidResult bool
	lastGetError      error
}

func (a *rngFeature) minIsDefinedAs(arg1 int) (err error) {
	err = nil
	if RNG.Min != arg1 {
		err = errors.New(fmt.Sprintf("Expected RNG.Min = %d.  Actual Result %d = %d", arg1, RNG.Min, arg1))
	}
	return
}

func (a *rngFeature) maxIsDefinedAs(value int) (err error) {
	err = nil
	if RNG.Max != value {
		err = errors.New(fmt.Sprintf("Expected RNG.Max = %d.  Actual Result %d = %d", value, RNG.Max, value))
	}
	return
}

func (a *rngFeature) initializeARandomNumberTo(value int) (err error) {
	err = nil
	a.num = RNG.RandomNumber{Value: value}
	return
}

func (a *rngFeature) initializeARandomNumberToURL(url string) (err error) {
	err = nil
	a.num = RNG.RandomNumber{Value: 0}
	err = a.num.GenerateFromService(url)
	if err != nil {
		err = errors.New(fmt.Sprintf("Failed to generate number from [%s]: %v", url, err))
	}
	a.lastGetError = err
	return
}

func (a *rngFeature) initializeARandomNumberToDefaultURL() (err error) {
	err = a.initializeARandomNumberToURL(RNG.DefaultRNGServiceURL)
	return
}

func (a *rngFeature) isValidSucceeds() (err error) {
	err = nil
	if actual := a.num.IsValid(); actual != true {
		err = errors.New(fmt.Sprintf("isValid failed. Actual Value [%v], GET result [%v]", a.num.Value, a.lastGetError))
	}
	return
}

func (a *rngFeature) isValidFails() (err error) {
	err = nil
	if actual := a.num.IsValid(); actual != false {
		err = errors.New("isValid succeeded")
	}
	return
}

func (a *rngFeature) reset() {
	a.num = RNG.RandomNumber{Value: 0}
}

func FeatureContext(s *godog.Suite) {
	cAPI := &choiceFeatures{}
	gAPI := &gameServiceFeatures{}
	rAPI := &rngFeature{}
	s.BeforeScenario(func(interface{}) {
		cAPI.reset()
		gAPI.reset()
		rAPI.reset()
	})

	s.Step(`^there are these choices:$`, cAPI.thereAreTheseChoices)
	s.Step(`^I call ValidChoices$`, cAPI.callValidChoices)
	s.Step(`^the response should match the string: \"([^"]*)\"$`, cAPI.theResponseShouldMatchString)
	s.Step(`^\"([^"]*)\" plays against \"([^"]*)\"$`, cAPI.lhsPlaysRhs)
	s.Step(`^the play result is \"([^"]*)\"$`, cAPI.lastPlayResult)
	s.Step(`^I set the active choice to \"([^"]*)\"$`, cAPI.setActiveChoice)
	s.Step(`^I play \"([^"]*)\"$`, cAPI.playActiveChoiceAgainst)

	s.Step(`^I set the target URL to \"([^"]*)\"$`, gAPI.setUrl)
	s.Step(`^I set the next ID to (\d+)$`, gAPI.setNextID)
	s.Step(`^I send a GET request to the \"([^"]*)\" endpoint$`, gAPI.sendGetToEndpoint)
	s.Step(`^I send a GET request to the \"([^"]*)\" endpoint with id set to (\d+)$`, gAPI.sendGetToEndpoint)
	s.Step(`^I send a POST request to the \"([^"]*)\" endpoint$`, gAPI.sendPostToEndpoint)
	s.Step(`^a choice other than \"([^"]*)\" is returned$`, gAPI.anyChoiceExcept)
	s.Step(`^the choice ids match$`, gAPI.choiceHasID)
	s.Step(`^the choice name is "([^"]*)"$`, gAPI.choiceHasName)
	s.Step("^a game result is returned$", gAPI.aGameResultWasReturned)

	s.Step(`^Min is defined as (\d+)$`, rAPI.minIsDefinedAs)
	s.Step(`^Max is defined as (\d+)$`, rAPI.maxIsDefinedAs)
	s.Step(`^I initialize a RandomNumber to (\d+)$`, rAPI.initializeARandomNumberTo)
	s.Step(`^I send a GET request to initialize a RandomNumber using the Default RNG endpoint$`, rAPI.initializeARandomNumberToDefaultURL)
	s.Step(`^I send a GET request to initialize a RandomNumber using the endpoint ([^"]*)$`, rAPI.initializeARandomNumberToURL)
	s.Step(`^isValid succeeds$`, rAPI.isValidSucceeds)
	s.Step(`^isValid fails$`, rAPI.isValidFails)
}

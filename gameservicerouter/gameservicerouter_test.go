package gameservicerouter_test

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

	"github.com/symbolgimmicks/rpsls/choice"

	"github.com/DATA-DOG/godog"
)

type validateFeature struct {
	resp                   *httptest.ResponseRecorder
	targetUrl              string
	jsonData               []byte
	lastChoiceReceived     choice.Choice
	lastChoiceListReceived []choice.Choice
	myClient               *http.Client
	nextId                 int
	lastGameResult         []byte
}

func (a *validateFeature) reset() {
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

func (a *validateFeature) setUrl(url string) (err error) {
	err = nil
	a.targetUrl = url
	return
}

func (a *validateFeature) setNextID(ID int) (err error) {
	err = nil
	a.nextId = ID
	return
}

func (a *validateFeature) sendGet(endpoint string, v interface{}) (resp *http.Response, err error) {
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

func (a *validateFeature) sendGetToEndpoint(endpoint string) (err error) {
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

func (a *validateFeature) sendPostToEndpoint(endpoint string) (err error) {
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

func (a *validateFeature) aGameResultWasReturned() (err error) {
	err = nil
	if a.lastGameResult == nil {
		err = errors.New("No game result has been received")
	}
	return
}

func (a *validateFeature) anyChoiceExcept(except string) (err error) {
	var filter choice.Choice
	if filter, err = choice.NewByString(except); err != nil {
		err = fmt.Errorf("Unexpected failure testing game result (%v)", err)
	} else if reflect.DeepEqual(a.lastChoiceReceived, filter) {
		err = fmt.Errorf("[%v] == [%v]", a.lastChoiceReceived, filter)
	} else {
		fmt.Printf("[%v] != [%v]", a.lastChoiceReceived, filter)
	}
	return
}

func (a *validateFeature) choiceHasID() (err error) {
	if a.lastChoiceReceived.ID != a.nextId {
		err = errors.New(fmt.Sprintf("[%d] != [%d]", a.lastChoiceReceived.ID, a.nextId))
	} else {
		fmt.Println(fmt.Sprintf("[%d] == [%d]", a.lastChoiceReceived.ID, a.nextId))
	}
	return
}

func (a *validateFeature) choiceHasName(name string) (err error) {
	if a.lastChoiceReceived.Name != name {
		err = errors.New(fmt.Sprintf("[%s] != [%s]", a.lastChoiceReceived.Name, name))
	} else {
		fmt.Println(fmt.Sprintf("[%s] == [%s]", a.lastChoiceReceived.Name, name))
	}
	return
}

func FeatureContext(s *godog.Suite) {
	validateApi := &validateFeature{}

	s.BeforeScenario(func(interface{}) {
		validateApi.reset()
	})

	/*
			When I set the next ID to 0
		         And I send a GET request to the "/choices/" endpoint with id set to 0
		         Then the choice with id 0 is returned
		         And the choice name is "empty"
	*/
	s.Step(`^I set the target URL to \"([^"]*)\"$`, validateApi.setUrl)
	s.Step(`^I set the next ID to (\d+)$`, validateApi.setNextID)
	s.Step(`^I send a GET request to the \"([^"]*)\" endpoint$`, validateApi.sendGetToEndpoint)
	s.Step(`^I send a GET request to the \"([^"]*)\" endpoint with id set to (\d+)$`, validateApi.sendGetToEndpoint)
	s.Step(`^I send a POST request to the \"([^"]*)\" endpoint$`, validateApi.sendPostToEndpoint)
	s.Step(`^a choice other than \"([^"]*)\" is returned$`, validateApi.anyChoiceExcept)
	s.Step(`^the choice ids match$`, validateApi.choiceHasID)
	s.Step(`^the choice name is "([^"]*)"$`, validateApi.choiceHasName)
	s.Step("^a game result is returned$", validateApi.aGameResultWasReturned)
}

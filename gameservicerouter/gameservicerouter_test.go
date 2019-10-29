package gameservicerouter_test

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"reflect"
	"strconv"
	"strings"
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
	lastThingReceived      interface{}
}

func (a *validateFeature) reset() {
	a.resp = httptest.NewRecorder()
	a.lastChoiceReceived, _ = choice.NewByID(0)
	a.targetUrl = ""

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

func (a *validateFeature) sendGet(endpoint string) (resp *http.Response, err error) {
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
		if err = json.NewDecoder(resp.Body).Decode(&a.lastThingReceived); err != nil {
			fmt.Println(fmt.Errorf("Failed to decode response (%v)", err))
		} else {
			fmt.Println(fmt.Sprintf("Received: [%v]", a.lastThingReceived))
		}
	}

	return
}

func (a *validateFeature) sendGetForRandomChoice(targetUrl string) (resp *http.Response, err error) {
	fmt.Println("Sending request to [" + targetUrl + "]")
	a.lastChoiceReceived = choice.EmptyChoice
	if resp, err := a.myClient.Get(targetUrl); err != nil {
		fmt.Println(fmt.Errorf("Failed to get a response (%v)", err))
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
		if err = json.NewDecoder(resp.Body).Decode(&a.lastChoiceReceived); err != nil {
			fmt.Println(fmt.Errorf("Failed to decode response (%v)", err))
		} else {
			fmt.Println(fmt.Sprintf("Received: [%v]", a.lastChoiceReceived))
		}
	}

	return
}

func (a *validateFeature) sendGetToEndpoint(endpoint string) (err error) {
	err = nil
	switch {
	case endpoint == "/choice":
		{
			if resp, err := a.sendGet(endpoint); err != nil {
				fmt.Println(fmt.Sprintf("Receieved [%v]", a.lastThingReceived))
			} else {
				fmt.Printf(fmt.Sprintf("Failed to get a response (%v): %v", err.Error(), resp))
			}
			break
		}
	case strings.HasPrefix(endpoint, "/choices/"):
		{
			var parts []string = strings.Split(endpoint, "/")
			if _, err = strconv.Atoi(parts[len(parts)-1]); err == nil {
				_, err = a.sendGet(endpoint)
			}
			break
		}
	default:
		break
	}

	return
}

func (a *validateFeature) sendGetToEndpointDirectIndex(endpoint string, id int) (err error) {
	if resp, err := a.sendGet(endpoint + strconv.Itoa(id)); err != nil {
		fmt.Printf(fmt.Sprintf("Failed to get a response (%v): %v", err.Error(), resp))
	}
	return
}

func (a *validateFeature) sendPostToEndpoint(except string) (err error) {
	err = nil
	return
}

func (a *validateFeature) thereAreTheseChoices(except string) (err error) {
	err = nil
	return
}

func (a *validateFeature) theseChoicesAreReturned(except string) (err error) {
	err = nil
	return
}

func (a *validateFeature) prepareJSONDataForPlay(data string) (err error) {
	err = nil
	return
}

func (a *validateFeature) aGameResultWasReturned(except string) (err error) {
	err = nil
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

func FeatureContext(s *godog.Suite) {
	validateApi := &validateFeature{}

	s.BeforeScenario(func(interface{}) {
		validateApi.reset()
	})

	s.Step(`^I set the target URL to \"([^"]*)\"$`, validateApi.setUrl)
	s.Step(`^I send a GET request to the \"([^"]*)\" endpoint$`, validateApi.sendGetToEndpoint)
	s.Step(`^I send a GET request to the \"([^"]*)\" endpoint with id set to (\d+)$`, validateApi.sendGetToEndpointDirectIndex)
	s.Step(`^I send a POST request to the \"([^"]*)\" endpoint$`, validateApi.sendPostToEndpoint)
	s.Step(`^a choice other than \"([^"]*)\" is returned$`, validateApi.anyChoiceExcept)
	s.Step(`^the following choices are returned:$`, validateApi.theseChoicesAreReturned)
	s.Step(`^I prepare the following JSON data for play:$`, validateApi.prepareJSONDataForPlay)
	s.Step(`^a game result is returned$`, validateApi.aGameResultWasReturned)
}

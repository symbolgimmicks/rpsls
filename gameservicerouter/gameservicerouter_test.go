package gameservicerouter_test

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"time"

	"github.com/symbolgimmicks/rpsls/choice"

	"github.com/DATA-DOG/godog"
)

var myClient = &http.Client{Timeout: 15 * time.Second}

type validateFeature struct {
	resp                   *httptest.ResponseRecorder
	targetUrl              string
	jsonData               []byte
	lastChoiceReceived     choice.Choice
	lastChoiceListReceived []choice.Choice
}

func (a *validateFeature) reset() {
	a.resp = httptest.NewRecorder()
	a.lastChoiceReceived = choice.EmptyChoice
	a.targetUrl = ""
}

func (a *validateFeature) setUrl(url string) (err error) {
	err = nil
	a.targetUrl = url
	return
}

func sendGet(targetUrl string, v interface{}) (resp *http.Response, err error) {
	if resp, err := myClient.Get(targetUrl); err != nil {
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
		err = json.NewDecoder(resp.Body).Decode(&v)
	}
	return
}

func (a *validateFeature) sendGetToEndpoint(endpoint string) (err error) {
	err = nil
	switch {
	case endpoint == "/choice":
		{
			_, err = sendGet(a.targetUrl+endpoint, &a.lastChoiceReceived)
			break
		}
	case strings.HasPrefix(endpoint, "/choices/"):
		{
			var parts []string = strings.Split(endpoint, "/")
			if _, err = strconv.Atoi(parts[len(parts)-1]); err == nil {
				_, err = sendGet(a.targetUrl+endpoint, &a.lastChoiceListReceived)
			}
			break
		}
	default:
		break
	}

	return
}

func (a *validateFeature) sendGetToEndpointDirectIndex(endpoint string, id int) (err error) {
	err = nil
	// url/choices/#
	var target string = a.targetUrl + endpoint + strconv.Itoa(id)
	_, err = sendGet(target, &a.lastChoiceListReceived)
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

func FeatureContext(s *godog.Suite) {
	validateApi := &validateFeature{}

	s.BeforeScenario(func(interface{}) {
		validateApi.reset()
	})

	s.Step(`^I set the target URL to \"([^"]*)\"$`, validateApi.setUrl)
	s.Step(`^I send a GET request to the \"([^"]*)\" endpoint$`, validateApi.sendGetToEndpoint)
	s.Step(`^I send a GET request to the \"([^"]*)\" endpoint with id set to (\d+)$`, validateApi.sendGetToEndpointDirectIndex)
	s.Step(`^I send a POST request to the \"([^"]*)\" endpoint$`, validateApi.sendPostToEndpoint)
	s.Step(`^a choice other than \"([^"]*)\" is returned$`, validateApi.thereAreTheseChoices)
	s.Step(`^the following choices are returned:$`, validateApi.theseChoicesAreReturned)
	s.Step(`^I prepare the following JSON data for play:$`, validateApi.prepareJSONDataForPlay)
	s.Step(`^a game result is returned$`, validateApi.aGameResultWasReturned)
}

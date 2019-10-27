package randomnumber

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

var myClient = &http.Client{Timeout: 10 * time.Second}

// DefaultRNGServiceURL - yup
var DefaultRNGServiceURL string = "https://codechallenge.boohma.com/random"

// Min - The minimum value Generate will give.  From specification for 3rd party service.
// unfortunately if this breaks we're in trouble...
var Min int = 1

// Max - The maximum value Generate will give.  From specification for 3rd party service.
var Max int = 100

// RandomNumber - expected response from [https://codechallenge.boohma.com/random]
type RandomNumber struct {
	Value int `json:"random_number"`
}

// IsValid - useable?
func (r RandomNumber) IsValid() (answer bool) {
	answer = !(r.Value > Max || r.Value < Min)
	return
}

// GenerateFromService - rolls using a service
func (r *RandomNumber) GenerateFromService(url string) (err error) {
	err = nil
	if getResponse, err := myClient.Get(url); err != nil {
		r.Value = -1
	} else {
		// https://github.com/DATA-DOG/godog/blob/master/examples/db/api_test.go
		// handle panic
		defer func() {
			switch t := recover().(type) {
			case string:
				err = fmt.Errorf(t)
			case error:
				err = t
			}
		}()

		defer getResponse.Body.Close()
		if err := json.NewDecoder(getResponse.Body).Decode(&r); err != nil {
			r.Value = -2
		}
	}
	return
}

// GenerateFromService - Rolls using a custom service
func GenerateFromService(url string) (data RandomNumber, err error) {
	data = RandomNumber{Value: Min}
	fmt.Println("get from: ", url)
	if url == "" {
		err = data.GenerateFromService(DefaultRNGServiceURL)
	} else {
		err = data.GenerateFromService(url)
	}
	return
}

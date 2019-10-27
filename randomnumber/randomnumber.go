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
func (r RandomNumber) IsValid() bool {
	if r.Value > Max || r.Value < Min {
		return false
	}
	return true
}

// GenerateFromService - rolls using a service
func (r *RandomNumber) GenerateFromService(url string) {
	getResponse, err := myClient.Get(url)
	if err != nil {
		fmt.Println("failed: ", err)
		r.Value = -1
	} else {
		defer getResponse.Body.Close()
		err2 := json.NewDecoder(getResponse.Body).Decode(&r)
		if err2 != nil {
			fmt.Println("failed: ", err2)
			r.Value = -2
		}
	}
}

// GenerateFromService - Rolls using a custom service
func GenerateFromService(url string) RandomNumber {
	data := RandomNumber{Value: Min}
	fmt.Println("get from: ", url)
	if url == "" {
		data.GenerateFromService(DefaultRNGServiceURL)
	} else {
		data.GenerateFromService(url)
	}
	return data
}

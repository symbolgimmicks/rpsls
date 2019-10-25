package randomnumber

//https://www.callicoder.com/golang-packages/
import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

//https://stackoverflow.com/questions/17156371/how-to-get-json-response-from-http-get
// JJB - default client does not have a timeout so you probably don't want to use it on purpose.
var myClient = &http.Client{Timeout: 10 * time.Second}

// RandomNumber - expected response from [https://codechallenge.boohma.com/random]
type randomNumber struct {
	RandomNumber int `json:"random_number"`
}

// Roll1To100 requests a number between 1 and 100 from
// the third-party endpoint "https://codechallenge.boohma.com/random"
func Roll1To100() int {

	data := &randomNumber{}

	getResponse, err := myClient.Get("https://codechallenge.boohma.com/random")
	if err != nil {
		fmt.Println("err: ", err.Error())
		return -1
	}
	defer getResponse.Body.Close()

	//https://stackoverflow.com/questions/26744873/converting-map-to-struct

	//go get github.com/mitchellh/mapstructure - is this required?
	err2 := json.NewDecoder(getResponse.Body).Decode(&data)
	if err2 != nil {
		fmt.Println("err2: ", err2.Error())
		return -2
	}

	return data.RandomNumber
}

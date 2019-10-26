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

// Min - The minimum value Generate will give.  From specification for 3rd party service.
var Min int = 1

// Max - The maximum value Generate will give.  From specification for 3rd party service.
var Max int = 100

// RandomNumber - expected response from [https://codechallenge.boohma.com/random]
type randomNumber struct {
	RandomNumber int `json:"random_number"`
}

// Generate - generate a number between Min and Max
func Generate() int {

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

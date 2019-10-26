# rpsls

rpsls (Rock Paper Scissors Lizard Spock) is a Golang REST API learning adventure

## Installation

1. Install & Configure Golang such
2. Validate your installation works
3. Go get gorilla/mut and gorilla/handler
``` bash go
get https://github.com/gorilla/mux#install
go get https://github.com/gorilla/handler#install
```
5. Clone https://github.com/symbolgimmicks/rpsls into %GOPATH$%/src/github.com/symbolgimmicks/rpsls
6. Using openssl create TSL keys

## Usage
The application attempts to host on port 4077 and is intended to provide the following REST API:


``` json
Choices
Get all the choices that are usable for the UI.
GET: /choices
Result: application/json
[
  {
    “id": integer [1-5],
    "name": string [12] (rock, paper, scissors, lizard, spock)
  }
]

Choice
Get a randomly generated choice
GET: /choice
Result: application/json
{
  "id": integer [1-5],
  "name" : string [12] (rock, paper, scissors, lizard, spock)
}
Play
 Play a round against a computer opponent
POST: /play
Data: application/json
{
  “player”: choice_id 
}
Result: application/json
{
  "results": string [12] (win, lose, tie),
  “player”: choice_id,
  “computer”:  choice_id
}
```
## Acknowledgment
https://github.com/gorilla/mux
https://github.com/gorilla/handler
https://gist.github.com/denji/12b3a568f092ab951456
https://github.com/gorilla/mux#static-files
https://golang.org/pkg/net/http/
https://golangcode.com/basic-https-server-with-certificate/
https://stackoverflow.com/questions/17156371/how-to-get-json-response-from-http-get
https://stackoverflow.com/questions/26744873/converting-map-to-struct
https://stackoverflow.com/questions/50625283/how-to-install-openssl-in-windows-10
https://www.alexedwards.net/blog/serving-static-sites-with-go
https://www.callicoder.com/golang-packages/
https://www.thepolyglotdeveloper.com/2017/10/handling-cors-golang-web-application/
	

## License

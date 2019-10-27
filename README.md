# rpsls

rpsls (Rock Paper Scissors Lizard Spock) is a Golang REST API learning adventure

## Installation

1. Install & Configure Golang
- https://golang.org/doc/install
- https://golang.org/wiki/SettingGOPATH
2. Validate your installation works
3. Go get rpsls
``` bash
go get https://github.com/symbolgimmicks/rpsls
```
4. Using openssl create TSL keys
- See https://golangcode.com/basic-https-server-with-certificate/ for an example on how to generate a self-signed key.

## Testing
- Utilizes the official Cucumber BDD framework for Golang https://github.com/DATA-DOG/godog
Test coverage
- [x] randomnumber testing
- [x] choice testing
- [ ] gameservicerouter testing (stubbed, 50% done)
- [ ] main testing

## Usage
The application attempts to host on port 4077 and is intended to provide the following REST API:

Choices
Get all the choices that are usable for the UI.
GET: /choices
Result: application/json
``` json
[
  {
    “id": integer [1-5],
    "name": string [12] (rock, paper, scissors, lizard, spock)
  }
]
```

Choice
Get a randomly generated choice
GET: /choice
Result: application/json
``` json
{
  "id": integer [1-5],
  "name" : string [12] (rock, paper, scissors, lizard, spock)
}
```

Play
 Play a round against a computer opponent
POST: /play
Data: application/json
``` json
{
  “player”: choice_id 
}
```

Result: application/json
``` json
{
  "results": string [12] (win, lose, tie),
  “player”: choice_id,
  “computer”:  choice_id
}
```
## Acknowledgment
- https://github.com/gorilla/mux
- https://github.com/gorilla/handler
- https://gist.github.com/denji/12b3a568f092ab951456
- https://github.com/gorilla/mux#static-files
- https://golang.org/pkg/net/http/
- https://golangcode.com/basic-https-server-with-certificate/
- https://stackoverflow.com/questions/17156371/how-to-get-json-response-from-http-get
- https://stackoverflow.com/questions/26744873/converting-map-to-struct
- https://stackoverflow.com/questions/50625283/how-to-install-openssl-in-windows-10
- https://www.alexedwards.net/blog/serving-static-sites-with-go
- https://www.callicoder.com/golang-packages/
- https://www.thepolyglotdeveloper.com/2017/10/handling-cors-golang-web-application/
- https://github.com/DATA-DOG/godog
- https://tutorialedge.net/golang/creating-restful-api-with-golang/
- https://www.tecmint.com/install-go-in-linux/
- https://golang.org/wiki/SettingGOPATH
- https://github.com/DATA-DOG/godog/tree/master/examples/api
	
## License

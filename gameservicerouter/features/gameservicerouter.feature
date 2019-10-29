Feature: Validate the Game Service Router
    In order to implement the game
    As a developer
    I need to create a game service

    Scenario: Handle /choice GET request to generate a random choice
        Given I set the target URL to "https://localhost:4077"
        #Try a few times, fun fun!
        When I send a GET request to the "/choice" endpoint
        Then a choice other than "empty" is returned
        When I send a GET request to the "/choice" endpoint
        Then a choice other than "empty" is returned
        When I send a GET request to the "/choice" endpoint
        Then a choice other than "empty" is returned
        When I send a GET request to the "/choice" endpoint
        Then a choice other than "empty" is returned
        When I send a GET request to the "/choice" endpoint
        Then a choice other than "empty" is returned

 
     Scenario: Handle /choices/id GET request
         Given I set the target URL to "https://localhost:4077"
         
         When I set the next ID to 0
         And I send a GET request to the "/choices/" endpoint
         Then the choice ids match
         And the choice name is "empty"

         When I set the next ID to 1
         And I send a GET request to the "/choices/" endpoint
         Then the choice ids match
         And the choice name is "rock"

         When I set the next ID to 2
         And I send a GET request to the "/choices/" endpoint
         Then the choice ids match
         And the choice name is "paper"

         When I set the next ID to 3
         And I send a GET request to the "/choices/" endpoint
         Then the choice ids match
         And the choice name is "scissors"

         When I set the next ID to 4
         And I send a GET request to the "/choices/" endpoint
         Then the choice ids match
         And the choice name is "lizard"

         When I set the next ID to 5
         And I send a GET request to the "/choices/" endpoint
         Then the choice ids match
         And the choice name is "spock"
   
    Scenario: Handle /play POST request
        Given I set the target URL to "https://localhost:4077"

        When I send a GET request to the "/choice" endpoint
        Then a choice other than "empty" is returned
        When I send a POST request to the "/play" endpoint
        Then a game result is returned

Feature: Validate the Game Service Router
    In order to implement the game
    As a developer
    I need to create a game service

    Scenario: Handle /choice GET request
        Given I set the target URL to "https://localhost:4077"
        When I send a GET request to the "/choice" endpoint
        Then a choice other than "empty" is returned

 
    # Scenario: Handle /choices/id GET request
    #     Given I set the target URL to "https://localhost:4077"
    #     When I send a GET request to the "/choices/" endpoint with id set to 0
    #     Then the following choices are returned:
    #     """
    #     """
   
    # Scenario: Handle /play POST request
    #     Given I set the target URL to "https://localhost:4077"
    #     When I prepare the following JSON data for play:
    #     """
    #     """
    #     And I send a POST request to the "/play" endpoint
    #     Then a game result is returned

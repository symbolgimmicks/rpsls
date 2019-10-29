Feature: Third-Party RNG BAT
    In order to implement the game
    As a developer
    I need to validate that the Random Number Generation works

    Scenario: Generate a valid, random number using the default 3rd party service that is known to exist
    When I send a GET request to initialize a RandomNumber using the Default RNG endpoint
    Then isValid succeeds

    Scenario: Generate a valid, random number using a 3rd party service that does not exist
    When I send a GET request to initialize a RandomNumber using the endpoint 'https://localhost/nopenopenope'
    Then isValid fails

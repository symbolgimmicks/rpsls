Feature: Validate Random Number Generation Validation
    In order to implement the game
    As a developer
    I need to validate that Random Number Generation works

    Scenario: Validate isValid works when numbers are in range of Min and Max
        Given Min is defined as 1
        And Max is defined as 100
        When I initialize a RandomNumber to 1
        Then isValid succeeds
        When I initialize a RandomNumber to 2
        Then isValid succeeds
        When I initialize a RandomNumber to 99
        Then isValid succeeds
        When I initialize a RandomNumber to 100
        Then isValid succeeds

    Scenario: Validate isValid works when the value is out of range of Min and Max
        Given Min is defined as 1
        And Max is defined as 100
        When I initialize a RandomNumber to 0
        Then isValid fails
        When I initialize a RandomNumber to 101
        Then isValid fails

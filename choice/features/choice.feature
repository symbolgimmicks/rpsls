Feature: Validate Random Number Generation Validation
    In order to implement the game
    As a developer
    I need to offer players a selection of game choices

    Scenario: Validate that ValidChoices returns only playable choices
        Given there are these choices:
            |id| name |
            |0 | empty|
            |1 | rock|
            |2 | paper|
            |3 | scissors|
            |4 | lizard|
            |5 | spock|
        When I call ValidChoices
        Then the response should match the string: "[{1 rock} {2 paper} {3 scissors} {4 lizard} {5 spock}]"

    Scenario: Validate Play
        When I set the active choice to "rock"
        And I play "scissors"
        Then the play result is "win"
        When I play "lizard"
        Then the play result is "win"
        When I play "spock"
        Then the play result is "lose"
        When I play "paper"
        Then the play result is "lose"
        When I play "rock"
        Then the play result is "tie"
        When I set the active choice to "paper"
        And I play "rock"
        Then the play result is "win"
        When I play "spock"
        Then the play result is "win"
        When I play "scissors"
        Then the play result is "lose"
        When I play "lizard"
        Then the play result is "lose"
        When I play "paper"
        Then the play result is "tie"
        When I set the active choice to "scissors"
        And I play "paper"
        Then the play result is "win"
        When I play "lizard"
        Then the play result is "win"
        When I play "spock"
        Then the play result is "lose"
        When I play "rock"
        Then the play result is "lose"
        When I play "scissors"
        Then the play result is "tie"
        When I set the active choice to "lizard"
        And I play "spock"
        Then the play result is "win"
        When I play "paper"
        Then the play result is "win"
        When I play "scissors"
        Then the play result is "lose"
        When I play "rock"
        Then the play result is "lose"
        When I play "lizard"
        Then the play result is "tie"
        When I set the active choice to "spock"
        And I play "rock"
        Then the play result is "win"
        When I play "scissors"
        Then the play result is "win"
        When I play "paper"
        Then the play result is "lose"
        When I play "lizard"
        Then the play result is "lose"
        When I play "spock"
        Then the play result is "tie"

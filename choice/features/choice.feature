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
        Then the response should match: 
        """
        {
            [
                {
                "id" = "1"
                "name = "rock"
                },
                {
                "id" = "2"
                "name = "paper"
                },
                {
                "id" = "3"
                "name = "scissors"
                },
                {
                "id" = "4"
                "name = "lizard"
                },
                {
                "id" = "5"
                "name = "spock"
                },
            ]
        }
        """
    
    Scenario: Validate ties
        When "rock" plays against "rock"
        Then the play result is "tie"
        When "paper" plays against "paper"
        Then the play result is "tie"
        When "scissors" plays against "scissors"
        Then the play result is "tie"
        When "lizard" plays against "lizard"
        Then the play result is "tie"
        When "spock" plays against "spock"
        Then the play result is "tie"

    Scenario: Validate Play
        When I set the active choice to "rock"
        And I play "paper"
        Then the play result is "lose"
        When I play "scissors"
        Then the play result is "win"
        When I play "lizard"
        Then the play result is "win"
        When I play "spock"
        Then the play result is "lose"

        When I set the active choice to "paper"
        And I play "rock"
        Then the play result is "win"
        When I play "scissors"
        Then the play result is "lose"
        When I play "lizard"
        Then the play result is "lose"
        When I play "spock"
        Then the play result is "win"

        When I set the active choice to "scissors"
        And I play "rock"
        Then the play result is "lose"
        When I play "paper"
        Then the play result is "win"
        When I play "lizard"
        Then the play result is "win"
        When I play "spock"
        Then the play result is "lose"

        When I set the active choice to "lizard"
        And I play "rock"
        Then the play result is "lose"
        When I play "paper"
        Then the play result is "win"
        When I play "scissors"
        Then the play result is "lose"
        When I play "spock"
        Then the play result is "win"

        When I set the active choice to "spock"
        And I play "rock"
        Then the play result is "win"
        When I play "paper"
        Then the play result is "lose"
        When I play "scissors"
        Then the play result is "win"
        When I play "lizard"
        Then the play result is "lose"

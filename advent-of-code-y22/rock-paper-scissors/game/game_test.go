package game

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPlay(t *testing.T) {
	var plays = []struct {
		Player1Move    Move
		Player2Move    Move
		Player1Outcome Outcome
		Player2Outcome Outcome
	}{
		{Player1Move: Rock, Player2Move: Rock, Player1Outcome: Draw, Player2Outcome: Draw},
		{Player1Move: Rock, Player2Move: Paper, Player1Outcome: Loss, Player2Outcome: Win},
		{Player1Move: Rock, Player2Move: Scissors, Player1Outcome: Win, Player2Outcome: Loss},
		{Player1Move: Paper, Player2Move: Paper, Player1Outcome: Draw, Player2Outcome: Draw},
		{Player1Move: Paper, Player2Move: Scissors, Player1Outcome: Loss, Player2Outcome: Win},
		{Player1Move: Paper, Player2Move: Rock, Player1Outcome: Win, Player2Outcome: Loss},
		{Player1Move: Scissors, Player2Move: Scissors, Player1Outcome: Draw, Player2Outcome: Draw},
		{Player1Move: Scissors, Player2Move: Rock, Player1Outcome: Loss, Player2Outcome: Win},
		{Player1Move: Scissors, Player2Move: Paper, Player1Outcome: Win, Player2Outcome: Loss},
	}
	for game, play := range plays {
		resultPlayer1, resultPlayer2 := Play(play.Player1Move, play.Player2Move)

		assert.Equal(t, play.Player1Outcome, resultPlayer1, "Player 1 outcome in game %d", game)
		assert.Equal(t, play.Player2Outcome, resultPlayer2, "Player 2 outcome in game %d", game)
	}
}

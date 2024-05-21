package gameutil

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/uszebr/thegamem/play/player"
)

func TestShufflePlayers(t *testing.T) {
	t.Run("empty slice", func(t *testing.T) {
		players := []*player.Player{}
		shuffledPlayers := ShufflePlayers(players)
		assert.Equal(t, players, shuffledPlayers, "ShufflePlayers should return the same slice for empty input")
	})

	t.Run("single player slice", func(t *testing.T) {
		player1 := &player.Player{}
		players := []*player.Player{player1}
		shuffledPlayers := ShufflePlayers(players)
		assert.Equal(t, players, shuffledPlayers, "ShufflePlayers should return the same slice for single element input")
	})

	t.Run("all players are in shuffled slice", func(t *testing.T) {
		player1 := &player.Player{}
		player2 := &player.Player{}
		player3 := &player.Player{}
		player4 := &player.Player{}
		player5 := &player.Player{}
		player6 := &player.Player{}
		players := []*player.Player{player1, player2, player3, player4, player5, player6}

		for i := 0; i < 1; i++ {
			shuffledPlayers := ShufflePlayers(players)
			assert.Equal(t, len(players), len(shuffledPlayers), "Shuffled players have the same size")
			assert.ElementsMatch(t, players, shuffledPlayers, "ShufflePlayers did not match original the slice in iteration %d", i)
		}
	})
}

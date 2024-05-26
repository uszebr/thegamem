package gameutil

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/uszebr/thegamem/play/model/modelfactory"
	"github.com/uszebr/thegamem/play/player"
)

var (
	green             = "alwaysgreen"
	red               = "alwaysred"
	factory           = modelfactory.GetModelFactory()
	playerInstance0   = player.New(factory.MustCreateModel(green))
	playerInstance1   = player.New(factory.MustCreateModel(green))
	playerInstance2   = player.New(factory.MustCreateModel(green))
	playerInstance3   = player.New(factory.MustCreateModel(green))
	playerInstance4   = player.New(factory.MustCreateModel(green))
	playerInstance5   = player.New(factory.MustCreateModel(green))
	playerInstance6   = player.New(factory.MustCreateModel(green))
	playerInstance7   = player.New(factory.MustCreateModel(green))
	playerInstance8   = player.New(factory.MustCreateModel(green))
	playerInstance9   = player.New(factory.MustCreateModel(green))
	playerInstance10  = player.New(factory.MustCreateModel(green))
	playerInstance11  = player.New(factory.MustCreateModel(green))
	playerInstancer0  = player.New(factory.MustCreateModel(red))
	playerInstancer1  = player.New(factory.MustCreateModel(red))
	playerInstancer2  = player.New(factory.MustCreateModel(red))
	playerInstancer3  = player.New(factory.MustCreateModel(red))
	playerInstancer4  = player.New(factory.MustCreateModel(red))
	playerInstancer5  = player.New(factory.MustCreateModel(red))
	playerInstancer6  = player.New(factory.MustCreateModel(red))
	playerInstancer7  = player.New(factory.MustCreateModel(red))
	playerInstancer8  = player.New(factory.MustCreateModel(red))
	playerInstancer9  = player.New(factory.MustCreateModel(red))
	playerInstancer10 = player.New(factory.MustCreateModel(red))
	playerInstancer11 = player.New(factory.MustCreateModel(red))
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

func TestGenerateByModelAndQuantity(t *testing.T) {
	t.Run("empty models map", func(t *testing.T) {
		players, err := GenerateByModelAndQuantity(map[string]int{})
		assert.EqualError(t, err, "generare players with empty model map")
		assert.Empty(t, players)
	})

	t.Run("invalid quantity", func(t *testing.T) {
		players, err := GenerateByModelAndQuantity(map[string]int{"alwaysgreen": 0})
		assert.EqualError(t, err, "wrong model: alwaysgreen quantity: 0")
		assert.Empty(t, players)
	})

	t.Run("invalid quantity", func(t *testing.T) {
		assert.Panics(t, func() {
			GenerateByModelAndQuantity(map[string]int{"fakemodel": 3})
		}, "The code did not panic")
	})

	//positive
	tests := []struct {
		models map[string]int
		len    int
	}{
		{models: map[string]int{"alwaysgreen": 1}, len: 1},
		{models: map[string]int{"alwaysred": 1}, len: 1},
		{models: map[string]int{"alwaysred": 1, "alwaysgreen": 1}, len: 2},
		{models: map[string]int{"alwaysred": 5, "alwaysgreen": 3}, len: 8},
		{models: map[string]int{"alwaysred": 5, "blindrevenge": 7, "alwaysgreen": 3}, len: 15},
		//to fail	{models: map[string]int{"alwaysred": 5, "alwaysgreen": 3}, len: 7},
	}
	for index, test := range tests {
		t.Run(fmt.Sprintf("valid input index: %v", index), func(t *testing.T) {
			players, err := GenerateByModelAndQuantity(test.models)
			assert.NoError(t, err)
			assert.Len(t, players, test.len)
		})
	}
}

func TestConvertPlayerFlatList(t *testing.T) {
	tests := []struct {
		name          string
		playerInput   [][]*player.Player
		playersOutput []*player.Player
	}{
		{name: "Player Convert empy",
			playerInput:   [][]*player.Player{},
			playersOutput: []*player.Player{},
		},
		{name: "Player Convert 4 player",
			playerInput: [][]*player.Player{
				{playerInstance0, playerInstance1},
				{playerInstance2},
				{playerInstance3},
			},
			playersOutput: []*player.Player{playerInstance0, playerInstance1, playerInstance2, playerInstance3},
		},
		{name: "Player Convert 6 player",
			playerInput: [][]*player.Player{
				{playerInstance0, playerInstance1},
				{playerInstance2, playerInstance3, playerInstance4},
				{playerInstance5, playerInstance6, playerInstance7, playerInstance8},
				{playerInstance9, playerInstance10, playerInstance11},
			},
			playersOutput: []*player.Player{playerInstance0, playerInstance1, playerInstance2, playerInstance3, playerInstance4, playerInstance5, playerInstance6, playerInstance7, playerInstance8, playerInstance9, playerInstance10, playerInstance11},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, test.playersOutput, ConvertPlayerFlatList(test.playerInput))
		})
	}
}

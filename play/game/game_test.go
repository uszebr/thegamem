package game

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/uszebr/thegamem/play/board"
	"github.com/uszebr/thegamem/play/board/pair"
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

func TestGameNew(t *testing.T) {
	tests := []struct {
		name           string
		col            int
		row            int
		interactions   int
		initialModels  []string
		rotation       int
		pairsCreator   board.PairsCreatorI
		shufflePlayers bool
	}{
		{name: "Game smoke 3 by 4", col: 3, row: 4, interactions: 10, initialModels: []string{"alwaysred"}, pairsCreator: pair.PairsNeighbour{}, shufflePlayers: false},
		{name: "Game smoke 5 by 10", col: 5, row: 10, interactions: 4, initialModels: []string{"alwaysred", "alwaysgreen"}, pairsCreator: pair.PairAll{}, shufflePlayers: false},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {

			game := New(test.col, test.row, test.interactions, test.initialModels, test.pairsCreator, test.rotation, test.shufflePlayers)
			assert.Equal(t, test.col, game.col)
			assert.Equal(t, test.row, game.row)
			assert.Equal(t, test.interactions, game.interactions)
			assert.Empty(t, game.boards)
			assert.Equal(t, test.initialModels, game.initialModels)
			assert.Equal(t, test.rotation, game.rotation)
			assert.NotNil(t, game.pairsCreatorI)
		})
	}
}

func TestAddBoard(t *testing.T) {
	tests := []struct {
		name           string
		col            int
		row            int
		interactions   int
		initialModels  []string
		rotation       int
		pairsCreator   board.PairsCreatorI
		roundsExpected int
		shufflePlayers bool
	}{
		// first two are impossible scenarious that need to be filtered on the form submit
		// ???might be make sense to sanitizing check those cases in game constructor???
		{name: "Add Board smoke 1 by 1", roundsExpected: 0, col: 1, row: 1, interactions: 10, initialModels: []string{"alwaysred"}, pairsCreator: pair.PairsNeighbour{}, rotation: 1, shufflePlayers: false},
		{name: "Add Board smoke 1 by 2", roundsExpected: 1, col: 1, row: 2, interactions: 10, initialModels: []string{"alwaysred"}, pairsCreator: pair.PairsNeighbour{}, rotation: 1, shufflePlayers: false},
		{name: "Add Board smoke 2 by 2", roundsExpected: 6, col: 2, row: 2, interactions: 10, initialModels: []string{"alwaysred"}, pairsCreator: pair.PairsNeighbour{}, rotation: 2, shufflePlayers: false},
		{name: "Add Board smoke 2 by 2 with rotation", roundsExpected: 6, col: 2, row: 2, interactions: 10, initialModels: []string{"alwaysred"}, pairsCreator: pair.PairAll{}, rotation: 1, shufflePlayers: false},
		{name: "Add Board smoke 2 by 2 with All pairs", roundsExpected: 6, col: 2, row: 2, interactions: 10, initialModels: []string{"alwaysred"}, pairsCreator: pair.PairAll{}, rotation: 1, shufflePlayers: false},
		{name: "Add Board smoke 5 by 10", roundsExpected: 1225, col: 5, row: 10, interactions: 4, initialModels: []string{"alwaysred", "alwaysgreen"}, pairsCreator: pair.PairAll{}, rotation: 1, shufflePlayers: false},
		{name: "Add Board smoke 3 by 4 with rotation", roundsExpected: 190, col: 5, row: 4, interactions: 4, initialModels: []string{"alwaysred", "alwaysgreen", "blindrevenge", "random"}, pairsCreator: pair.PairAll{}, rotation: 5, shufflePlayers: false},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {

			game := New(test.col, test.row, test.interactions, test.initialModels, test.pairsCreator, test.rotation, test.shufflePlayers)
			err := game.AddNewBoard()
			assert.Nil(t, err)
			assert.Len(t, game.boards, 1)
			firstBoard := game.boards[0]
			assert.Equal(t, test.roundsExpected, len(firstBoard.GetAllRounds()))
			if test.name == "Add Board smoke 1 by 1" || test.name == "Add Board smoke 1 by 2" {
				return
			}
			firstLoozers := firstBoard.GetLoosers()
			firstWinners := firstBoard.GetWinners()
			var firstWinnersPlayers []*player.Player
			for _, firstWinner := range firstWinners {
				firstWinnersPlayers = append(firstWinnersPlayers, firstWinner.Player)
			}
			var firstLoosersPlayers []*player.Player
			for _, firstLooser := range firstLoozers {
				firstLoosersPlayers = append(firstLoosersPlayers, firstLooser.Player)
			}
			assert.Len(t, firstLoozers, test.rotation)
			assert.Len(t, firstWinners, test.rotation)
			err = game.AddNewBoard()
			assert.Nil(t, err)
			assert.Len(t, game.boards, 2)
			secondBoard := game.boards[1]
			assert.Equal(t, test.roundsExpected, len(secondBoard.GetAllRounds()))
			secondLoozers := secondBoard.GetLoosers()
			secondWinners := secondBoard.GetWinners()

			assert.Len(t, secondLoozers, test.rotation)
			assert.Len(t, secondWinners, test.rotation)
			secondAllPlayer := secondBoard.GetPlayersOneSlice()
			//checking that winners are in the second Board; loosers are not in the second board
			for _, player := range firstWinnersPlayers {
				assert.Contains(t, secondAllPlayer, player)
			}
			for _, player := range firstLoosersPlayers {
				assert.NotContains(t, secondAllPlayer, player)
			}

		})
	}
}

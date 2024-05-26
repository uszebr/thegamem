package board

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/uszebr/thegamem/play/board/pair"
	"github.com/uszebr/thegamem/play/model/modelfactory"
	"github.com/uszebr/thegamem/play/player"
)

func TestNewNegative(t *testing.T) {
	tests := []struct {
		name         string
		cols         int
		rows         int
		interactions int
		players      []*player.Player
		errorText    string
	}{
		{name: "Zero Cols", cols: 0, rows: 10, interactions: 10, errorText: "Creating board issue parameters col"},
		{name: "Negative Cols", cols: -4, rows: 10, interactions: 10, errorText: "Creating board issue parameters col"},
		{name: "Zero Rows", cols: 3, rows: 0, interactions: 10, errorText: "Creating board issue parameters col"},
		{name: "Negative Rows", cols: 4, rows: -10, interactions: 10, errorText: "Creating board issue parameters col"},
		{name: "Zero Interactions", cols: 90, rows: 10, interactions: 0, errorText: "Creating board issue interactions:"},
		{name: "Negative Interactions", cols: 90, rows: 10, interactions: -8, errorText: "Creating board issue interactions:"},
		// blank players becouse checking only wrong quantity err
		{name: "Players Quantity", cols: 2, rows: 3, interactions: 8, players: []*player.Player{{}, {}, {}, {}, {}}, errorText: "quantity of players:"},

		//tofail {name: "Zero Cols", cols: 0, rows: 10, interactions: 10, errorText: "fakeerrmsg"},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			board, err := New(test.cols, test.rows, test.players, test.interactions, pair.PairsNeighbour{})
			assert.Error(t, err)
			assert.ErrorContains(t, err, test.errorText)
			assert.Empty(t, board)
		})
	}
}

func TestNewSmoke(t *testing.T) {
	green := "alwaysgreen"
	red := "alwaysred"
	factory := modelfactory.GetModelFactory()
	playerInstance0 := player.New(factory.MustCreateModel(green))
	playerInstance1 := player.New(factory.MustCreateModel(green))
	playerInstance2 := player.New(factory.MustCreateModel(green))
	playerInstance3 := player.New(factory.MustCreateModel(green))
	playerInstance4 := player.New(factory.MustCreateModel(green))
	playerInstance5 := player.New(factory.MustCreateModel(green))
	playerInstance6 := player.New(factory.MustCreateModel(green))
	playerInstance7 := player.New(factory.MustCreateModel(green))
	playerInstance8 := player.New(factory.MustCreateModel(green))
	playerInstance9 := player.New(factory.MustCreateModel(green))
	playerInstance10 := player.New(factory.MustCreateModel(green))
	playerInstance11 := player.New(factory.MustCreateModel(green))
	playerInstancer0 := player.New(factory.MustCreateModel(red))
	playerInstancer1 := player.New(factory.MustCreateModel(red))
	playerInstancer2 := player.New(factory.MustCreateModel(red))
	playerInstancer3 := player.New(factory.MustCreateModel(red))
	playerInstancer4 := player.New(factory.MustCreateModel(red))
	playerInstancer5 := player.New(factory.MustCreateModel(red))
	playerInstancer6 := player.New(factory.MustCreateModel(red))
	playerInstancer7 := player.New(factory.MustCreateModel(red))
	playerInstancer8 := player.New(factory.MustCreateModel(red))
	playerInstancer9 := player.New(factory.MustCreateModel(red))
	playerInstancer10 := player.New(factory.MustCreateModel(red))
	playerInstancer11 := player.New(factory.MustCreateModel(red))

	tests := []struct {
		name          string
		cols          int
		rows          int
		interactions  int
		players       []*player.Player
		roundScoreSum int
		pairs         int
		modelName     string
	}{
		//alwaysgreen
		{name: "Positive 4 players alwaysgreen ", modelName: green, pairs: 6, cols: 2, rows: 2, roundScoreSum: 50, players: []*player.Player{playerInstance0, playerInstance1, playerInstance2, playerInstance3}, interactions: 10},
		{name: "Positive 6 players alwaysgreen ", modelName: green, pairs: 15, cols: 2, rows: 3, roundScoreSum: 75, players: []*player.Player{playerInstance0, playerInstance1, playerInstance2, playerInstance3, playerInstance4, playerInstance5}, interactions: 15},
		{name: "Positive 12 players alwaysgreen ", modelName: green, pairs: 48, cols: 3, rows: 4, roundScoreSum: 25, players: []*player.Player{playerInstance0, playerInstance1, playerInstance2, playerInstance3, playerInstance4, playerInstance5, playerInstance6, playerInstance7, playerInstance8, playerInstance9, playerInstance10, playerInstance11}, interactions: 5},
		//alwaysred
		{name: "Positive 4 players alwaysred ", modelName: red, pairs: 6, cols: 2, rows: 2, roundScoreSum: 10, players: []*player.Player{playerInstancer0, playerInstancer1, playerInstancer2, playerInstancer3}, interactions: 10},
		{name: "Positive 6 players alwaysred ", modelName: red, pairs: 15, cols: 2, rows: 3, roundScoreSum: 15, players: []*player.Player{playerInstancer0, playerInstancer1, playerInstancer2, playerInstancer3, playerInstancer4, playerInstancer5}, interactions: 15},
		{name: "Positive 12 players alwaysred ", modelName: red, pairs: 48, cols: 3, rows: 4, roundScoreSum: 5, players: []*player.Player{playerInstancer0, playerInstancer1, playerInstancer2, playerInstancer3, playerInstancer4, playerInstancer5, playerInstancer6, playerInstancer7, playerInstancer8, playerInstancer9, playerInstancer10, playerInstancer11}, interactions: 5},

		//to fail {name: "Positive 12 players alwaysred ", modelName: red, pairs: 45, cols: 3, rows: 4, roundScoreSum: 5, players: []*player.Player{playerInstancer0, playerInstancer1, playerInstancer2, playerInstancer3, playerInstancer4, playerInstancer5, playerInstancer6, playerInstancer7, playerInstancer8, playerInstancer9, playerInstancer10, playerInstancer11}, interactions: 5},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			board, err := New(test.cols, test.rows, test.players, test.interactions, pair.PairsNeighbour{})

			assert.Nil(t, err)
			assert.Len(t, board.GetPlayersOneSlice(), test.cols*test.rows)
			assert.Len(t, board.pairs, test.pairs)
			assert.Len(t, board.rounds, test.pairs)
			for _, round := range board.rounds {
				assert.Equal(t, test.roundScoreSum, round.Left.RoundScoreSum)
				assert.Equal(t, test.roundScoreSum, round.Right.RoundScoreSum)
				assert.Equal(t, test.interactions, len(round.Right.Signals))
				assert.Equal(t, test.interactions, len(round.Right.Scores))
				assert.Equal(t, test.interactions, len(round.Right.ScoreSums))
				assert.Equal(t, test.interactions, len(round.Left.Signals))
				assert.Equal(t, test.interactions, len(round.Left.Scores))
				assert.Equal(t, test.interactions, len(round.Left.ScoreSums))

				assert.Equal(t, test.modelName, round.Left.GetPlayer().GetModelName())
				assert.Equal(t, test.modelName, round.Right.GetPlayer().GetModelName())
			}
		})
	}

}

func TestBoardScores(t *testing.T) {
	green := "alwaysgreen"
	red := "alwaysred"
	factory := modelfactory.GetModelFactory()
	playerInstance0 := player.New(factory.MustCreateModel(green))
	playerInstance1 := player.New(factory.MustCreateModel(green))
	playerInstance2 := player.New(factory.MustCreateModel(green))
	playerInstance3 := player.New(factory.MustCreateModel(green))
	playerInstance4 := player.New(factory.MustCreateModel(green))
	playerInstance5 := player.New(factory.MustCreateModel(green))
	playerInstance6 := player.New(factory.MustCreateModel(green))
	playerInstance7 := player.New(factory.MustCreateModel(green))
	playerInstance8 := player.New(factory.MustCreateModel(green))
	playerInstance9 := player.New(factory.MustCreateModel(green))
	playerInstance10 := player.New(factory.MustCreateModel(green))
	playerInstance11 := player.New(factory.MustCreateModel(green))
	playerInstancer0 := player.New(factory.MustCreateModel(red))
	playerInstancer1 := player.New(factory.MustCreateModel(red))
	playerInstancer2 := player.New(factory.MustCreateModel(red))
	playerInstancer3 := player.New(factory.MustCreateModel(red))
	playerInstancer4 := player.New(factory.MustCreateModel(red))
	playerInstancer5 := player.New(factory.MustCreateModel(red))
	playerInstancer6 := player.New(factory.MustCreateModel(red))
	playerInstancer7 := player.New(factory.MustCreateModel(red))
	playerInstancer8 := player.New(factory.MustCreateModel(red))
	playerInstancer9 := player.New(factory.MustCreateModel(red))
	playerInstancer10 := player.New(factory.MustCreateModel(red))
	playerInstancer11 := player.New(factory.MustCreateModel(red))

	tests := []struct {
		name          string
		cols          int
		rows          int
		interactions  int
		players       []*player.Player
		roundScoreSum int

		modelName  string
		boardScore int
	}{
		//alwaysgreen
		{name: "Scores 4 players alwaysgreen ", modelName: green, cols: 2, rows: 2, players: []*player.Player{playerInstance0, playerInstance1, playerInstance2, playerInstance3}, interactions: 10, boardScore: 150},
		{name: "Scores 6 players alwaysgreen ", modelName: green, cols: 2, rows: 3, players: []*player.Player{playerInstance0, playerInstance1, playerInstance2, playerInstance3, playerInstance4, playerInstance5}, interactions: 10, boardScore: 250},
		{name: "Scores 12 players alwaysgreen ", modelName: green, cols: 3, rows: 4, players: []*player.Player{playerInstance0, playerInstance1, playerInstance2, playerInstance3, playerInstance4, playerInstance5, playerInstance6, playerInstance7, playerInstance8, playerInstance9, playerInstance10, playerInstance11}, interactions: 10, boardScore: 400},

		//alwaysred
		{name: "Scores 4 players red ", modelName: red, cols: 2, rows: 2, players: []*player.Player{playerInstancer0, playerInstancer1, playerInstancer2, playerInstancer3}, interactions: 10, boardScore: 30},
		{name: "Scores 6 players red ", modelName: red, cols: 2, rows: 3, players: []*player.Player{playerInstancer0, playerInstancer1, playerInstancer2, playerInstancer3, playerInstancer4, playerInstancer5}, interactions: 10, boardScore: 50},
		{name: "Scores 12 players red ", modelName: red, cols: 3, rows: 4, players: []*player.Player{playerInstancer0, playerInstancer1, playerInstancer2, playerInstancer3, playerInstancer4, playerInstancer5, playerInstancer6, playerInstancer7, playerInstancer8, playerInstancer9, playerInstancer10, playerInstancer11}, interactions: 10, boardScore: 80},

		//to fail	{name: "Scores 12 players red ", modelName: red, cols: 3, rows: 4, players: []*player.Player{playerInstancer0, playerInstancer1, playerInstancer2, playerInstancer3, playerInstancer4, playerInstancer5, playerInstancer6, playerInstancer7, playerInstancer8, playerInstancer9, playerInstancer10, playerInstancer11}, interactions: 10, boardScore: 66},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			board, _ := New(test.cols, test.rows, test.players, test.interactions, pair.PairsNeighbour{})
			boardScore := board.GetBoardPlayerScores()
			playersFromScore := make([]*player.Player, 0, test.cols*test.rows)
			//checking each player score
			for player, score := range boardScore {
				playersFromScore = append(playersFromScore, player)
				assert.Equal(t, score, test.boardScore)
			}
			// checking all player have scores
			assert.ElementsMatch(t, test.players, playersFromScore)
		})
	}
}

// transforming map of scores to the list of ratings(sorted desc)
func TestRatings(t *testing.T) {
	green := "alwaysgreen"
	red := "alwaysred"
	factory := modelfactory.GetModelFactory()
	playerInstance0 := player.New(factory.MustCreateModel(green))
	playerInstance1 := player.New(factory.MustCreateModel(green))
	playerInstance2 := player.New(factory.MustCreateModel(green))
	playerInstance3 := player.New(factory.MustCreateModel(green))

	playerInstancer0 := player.New(factory.MustCreateModel(red))
	playerInstancer1 := player.New(factory.MustCreateModel(red))
	playerInstancer2 := player.New(factory.MustCreateModel(red))
	playerInstancer3 := player.New(factory.MustCreateModel(red))

	tests := []struct {
		name         string
		cols         int
		rows         int
		interactions int
		players      []*player.Player
		scores       []int
	}{

		{name: "Ratings 4 players mix", scores: []int{5, 5, 1, 1}, cols: 2, rows: 2, players: []*player.Player{playerInstance0, playerInstance1, playerInstancer0, playerInstancer1}, interactions: 1},
		{name: "Ratings 4 players green", scores: []int{15, 15, 15, 15}, cols: 2, rows: 2, players: []*player.Player{playerInstance0, playerInstance1, playerInstance2, playerInstance3}, interactions: 1},
		{name: "Ratings 4 players red", scores: []int{3, 3, 3, 3}, cols: 2, rows: 2, players: []*player.Player{playerInstancer0, playerInstancer1, playerInstancer2, playerInstancer3}, interactions: 1},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			board, _ := New(test.cols, test.rows, test.players, test.interactions, pair.PairsNeighbour{})
			ratings := board.GetRatings()
			scores := board.GetBoardPlayerScores()
			assert.Len(t, ratings, len(test.players))
			for index, rating := range ratings {
				assert.Contains(t, test.players, rating.player)
				assert.Equal(t, rating.score, scores[rating.player])
				assert.Equal(t, test.scores[index], rating.score)
			}
		})
	}

}

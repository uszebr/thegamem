package round

import (
	"fmt"

	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/uszebr/thegamem/play/model/modelfactory"
	"github.com/uszebr/thegamem/play/player"
	"github.com/uszebr/thegamem/play/signal"
)

var (
	factory = modelfactory.GetModelFactory()
)

func TestNewRound(t *testing.T) {
	playerLeft := player.New(factory.MustCreateModel("alwaysgreen"))
	playerRight := player.New(factory.MustCreateModel("alwaysgreen"))
	interactions := 10
	round := New(playerLeft, playerRight, interactions)

	assert.NotNil(t, round.uuid)
	assert.NotNil(t, round.Left)
	assert.NotNil(t, round.Right)
	assert.Equal(t, interactions, round.intractionsQuantity)
	assert.True(t, round.aproximateInteractions >= interactions-((interactions*percentToRandomize)/100))
	assert.True(t, round.aproximateInteractions <= interactions+((interactions*percentToRandomize)/100))
}

func TestRoundRounders(t *testing.T) {
	playerLeft := player.New(factory.MustCreateModel("alwaysgreen"))
	playerRight := player.New(factory.MustCreateModel("alwaysgreen"))
	interactions := 50
	round := New(playerLeft, playerRight, interactions)
	assert.Equal(t, len(round.Left.Scores), interactions, "Scores quantity are equal interactions left")
	assert.Equal(t, len(round.Left.ScoreSums), interactions, "Scores sums quantity are equal interactions left")
	assert.Equal(t, len(round.Left.Signals), interactions, "Signals quantity are equal interactions left")
	assert.Equal(t, len(round.Right.Scores), interactions, "Scores quantity are equal interactions right")
	assert.Equal(t, len(round.Right.ScoreSums), interactions, "Scores sums quantity are equal interactions right")
	assert.Equal(t, len(round.Right.Signals), interactions, "Signals quantity are equal interactions right")
	var sumActualLeft = 0
	var sumActualRight = 0
	for i := range interactions {
		sumActualLeft += round.Left.Scores[i]
		sumActualRight += round.Right.Scores[i]
		assert.Equal(t, round.Left.ScoreSums[i], sumActualLeft, fmt.Sprintf("Current scoresum left: %v", i))
		assert.Equal(t, round.Right.ScoreSums[i], sumActualRight, fmt.Sprintf("Current scoresum right: %v", i))
		//signals are green
		assert.Equal(t, round.Left.Signals[i], signal.Green, fmt.Sprintf("Current signal green left: %v", i))
		assert.Equal(t, round.Right.Signals[i], signal.Green, fmt.Sprintf("Current signal green right: %v", i))
	}
	assert.Equal(t, round.Left.RoundScoreSum, sumActualLeft, "Round score sum left")
	assert.Equal(t, round.Right.RoundScoreSum, sumActualRight, "Round score sum right")
}

func TestGetApproximateInteractionsQuantityPositive(t *testing.T) {
	attempts := 20
	tests := []struct {
		interactionQuantity int
		percentRandom       int
		resultMin           int
		resultMax           int
	}{
		{interactionQuantity: 100, percentRandom: 10, resultMin: 90, resultMax: 110},
		{interactionQuantity: 100, percentRandom: 3, resultMin: 97, resultMax: 103},
		{interactionQuantity: 10, percentRandom: 10, resultMin: 9, resultMax: 19},
		{interactionQuantity: 10, percentRandom: 0, resultMin: 10, resultMax: 10},
		{interactionQuantity: 1000, percentRandom: 0, resultMin: 1000, resultMax: 1000},

		//to fail	{interactionQuantity: 100, percentRandom: 10, resultMin: 110, resultMax: 110},
	}

	for index, test := range tests {
		for i := range attempts {
			t.Run(fmt.Sprintf("Test: %v attempt: %v", index, i), func(t *testing.T) {
				actualResult := getApproximateInteractionsQuantity(test.interactionQuantity, test.percentRandom)
				assert.True(t, actualResult <= test.resultMax && actualResult >= test.resultMin)
			})
		}
	}

}

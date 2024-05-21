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

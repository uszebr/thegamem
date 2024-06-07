package round

import (
	"fmt"
	"math/rand"

	"github.com/google/uuid"
	"github.com/uszebr/thegamem/play/player"
	"github.com/uszebr/thegamem/play/rounder"
)

const (
	percentToRandomize = 5
)

// here two rounders(players) set of interactions(echanging signals) occurs
type Round struct {
	uuid                   uuid.UUID
	Left                   *rounder.Rounder
	Right                  *rounder.Rounder
	intractionsQuantity    int
	aproximateInteractions int
}

// Round created already played all interactions done, scores/sums calculated
// Scores/sums/round sum are needed to quickly show all data in handler
func New(playerLeft *player.Player, playerRight *player.Player, interactions int) Round {
	roundUUID := uuid.New()
	round := Round{
		uuid:                   roundUUID,
		Left:                   rounder.New(rounder.LEFT, playerLeft, roundUUID),
		Right:                  rounder.New(rounder.RIGHT, playerRight, roundUUID),
		intractionsQuantity:    interactions,
		aproximateInteractions: getApproximateInteractionsQuantity(interactions, percentToRandomize),
	}
	round.play()
	return round
}

// All interactions, full round play
func (round *Round) play() {
	for range round.intractionsQuantity {
		signalLeft := round.Left.GetPlayer().PlayOne(round.Left.Signals, round.Right.Signals, round.aproximateInteractions)
		signalRight := round.Right.GetPlayer().PlayOne(round.Right.Signals, round.Left.Signals, round.aproximateInteractions)
		//appending scores, and Score sum
		scoreLeft := signalLeft.CalcScore(signalRight)
		round.Left.Scores = append(round.Left.Scores, scoreLeft)
		round.Left.RoundScoreSum += scoreLeft
		round.Left.ScoreSums = append(round.Left.ScoreSums, round.Left.RoundScoreSum)
		scoreRight := signalRight.CalcScore(signalLeft)
		round.Right.Scores = append(round.Right.Scores, scoreRight)
		round.Right.RoundScoreSum += scoreRight
		round.Right.ScoreSums = append(round.Right.ScoreSums, round.Right.RoundScoreSum)
		//appending signals
		round.Left.Signals = append(round.Left.Signals, signalLeft)
		round.Right.Signals = append(round.Right.Signals, signalRight)
	}
}

func (round *Round) GetUUID() string {
	return round.uuid.String()
}

// Randomizing interactions quantities to hide exact quantity of interactions(from model)
// to avoid "always red" in the last interaction of the Round
func getApproximateInteractionsQuantity(interactions int, percent int) int {
	if percent < 0 {
		panic(fmt.Sprintf("percent to randomize less then 0: %v", percent))
	}
	max := interactions + (interactions * percent / 100)
	min := interactions - (interactions * percent / 100)
	return rand.Intn(max-min+1) + min
}

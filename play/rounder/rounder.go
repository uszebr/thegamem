package rounder

import (
	"github.com/google/uuid"
	"github.com/uszebr/thegamem/play/player"
	"github.com/uszebr/thegamem/play/signal"
)

type Location string

// Player can be on the right or on the left
const (
	LEFT  Location = "left"
	RIGHT Location = "right"
)

// Rounder is a struct with Player and it's signals scores and score sums _
type Rounder struct {
	location  Location
	player    *player.Player
	roundUUID uuid.UUID
	Signals   []signal.Signal
	Scores    []int
	//sum of scrore for particular interaction for showing purposes(not to calculate this each request)
	ScoreSums []int
	// sum of all scores
	RoundScoreSum int
}

func New(location Location, player *player.Player, rounduuid uuid.UUID) *Rounder {
	return &Rounder{
		location:      location,
		player:        player,
		roundUUID:     rounduuid,
		Signals:       []signal.Signal{},
		Scores:        []int{},
		ScoreSums:     []int{},
		RoundScoreSum: 0,
	}
}

func (rounder Rounder) GetPlayer() *player.Player {
	return rounder.player
}

func (rounder Rounder) GetUUID() uuid.UUID {
	return rounder.roundUUID
}

func (rounder Rounder) GetLocation() string {
	return string(rounder.location)
}

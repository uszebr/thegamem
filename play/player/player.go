package player

import (
	"github.com/google/uuid"
	"github.com/uszebr/thegamem/play/player/playerutil"
	"github.com/uszebr/thegamem/play/signal"
)

type Player struct {
	name   string
	modeli ModelI
	uuid   uuid.UUID
}

func New(model ModelI) *Player {
	return &Player{
		name:   playerutil.GenerateRandomName(),
		modeli: model,
		uuid:   uuid.New(),
	}
}

// Playing one interaction with own and  opponent signal
func (player Player) PlayOne(myHistory []signal.Signal, opponentHistory []signal.Signal, aproximateInteractions int) signal.Signal {
	return player.modeli.CalculateSignal(myHistory, opponentHistory, aproximateInteractions)
}

func (player Player) GetModelName() string {
	return player.modeli.GetName()
}

type CalculateSignalI interface {
	CalculateSignal(myHistory []signal.Signal, opponentHistory []signal.Signal, aproximateInteractions int) signal.Signal
}

type NameInfoI interface {
	GetName() string
	GetDescription() string
	GetIcon() string
}

type ModelI interface {
	CalculateSignalI
	NameInfoI
}

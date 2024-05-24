package modrandom

import (
	"math/rand"

	"github.com/uszebr/thegamem/play/model"
	"github.com/uszebr/thegamem/play/signal"
)

const (
	modelName   = "random"
	description = "Always response with random signal"
	iconPath    = "/static/image/icon/jellyfish.svg"
)

type ModRandom struct{}

func (modRandom ModRandom) GetModel() model.Model {
	action := func(myHistory []signal.Signal, opponentHistory []signal.Signal, aproximateInteractions int) signal.Signal {
		var signalR signal.Signal
		if rand.Intn(2) == 0 {
			signalR = signal.Green
		} else {
			signalR = signal.Red
		}
		return signalR
	}
	return model.New(modelName, description, iconPath, action)
}

package modcopystartgreen

import (
	"github.com/uszebr/thegamem/play/model"
	"github.com/uszebr/thegamem/play/signal"
)

const (
	modelName   = "copystartgreen"
	description = "Copy opponent signal, starts with green"
	iconPath    = "/static/image/icon/whale.svg"
)

type ModCopyStrartGreen struct{}

func (copyStrartGreen ModCopyStrartGreen) GetModel() model.Model {
	action := func(myHistory []signal.Signal, opponentHistory []signal.Signal, aproximateInteractions int) signal.Signal {
		if len(opponentHistory) == 0 {
			return signal.Green
		}
		return opponentHistory[len(opponentHistory)-1]
	}
	return model.New(modelName, description, iconPath, action)
}

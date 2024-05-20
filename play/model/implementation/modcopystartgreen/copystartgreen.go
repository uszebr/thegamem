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
	action := func(myHistory []signal.Signal, oponentHistory []signal.Signal, aproximateInteractions int) signal.Signal {
		if len(oponentHistory) == 0 {
			return signal.Green
		}
		return oponentHistory[len(oponentHistory)-1]
	}
	return model.New(modelName, description, iconPath, action)
}

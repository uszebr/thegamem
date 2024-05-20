package modalwaysgreen

import (
	"github.com/uszebr/thegamem/play/model"
	"github.com/uszebr/thegamem/play/signal"
)

const (
	modelName   = "alwaysgreen"
	description = "Always returns green(cooperation) signal"
	iconPath    = "/static/image/icon/rabbit.svg"
)

type ModAlwaysGreen struct {
}

func (alwaysGreen ModAlwaysGreen) GetModel() model.Model {
	action := func(myHistory []signal.Signal, opponentHistory []signal.Signal, approximateInteractions int) signal.Signal {
		return signal.Green
	}
	return model.New(modelName, description, iconPath, action)
}

package modblindrevenge

import (
	"github.com/uszebr/thegamem/play/model"
	"github.com/uszebr/thegamem/play/signal"
)

const (
	modelName   = "blindrevenge"
	description = "Returns Green(Cooperation) signal until opponent send Red(Confrontation). Aftre first opponent Red returns Red until end of the round"
	iconPath    = "/static/image/icon/pig.svg"
)

type ModBlindRevenge struct{}

func (blindRevenge ModBlindRevenge) GetModel() model.Model {
	action := func(myHistory []signal.Signal, opponentHistory []signal.Signal, approximateInteractions int) signal.Signal {
		for _, signalOp := range opponentHistory {
			if signalOp == signal.Red {
				return signal.Red
			}
		}
		return signal.Green
	}
	return model.New(modelName, description, iconPath, action)
}

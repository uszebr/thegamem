package modalwaysred

import (
	"github.com/uszebr/thegamem/play/model"
	"github.com/uszebr/thegamem/play/signal"
)

const (
	modelName   = "alwaysred"
	description = "Always returns red(confrontation) signal"
	iconPath    = "/static/image/icon/crocodile.svg"
)

type ModAlwaysRed struct{}

func (alwaysRed ModAlwaysRed) GetModel() model.Model {
	action := func(myHistory []signal.Signal, opponentHistory []signal.Signal, approximateInteractions int) signal.Signal {
		return signal.Red
	}
	return model.New(modelName, description, iconPath, action)
}

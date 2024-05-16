package modcopystartred

import (
	"github.com/uszebr/thegamem/play/model"
	"github.com/uszebr/thegamem/play/signal"
)

const (
	modelName   = "copystartred"
	description = "Copy opponent signal, starts with red"
	iconPath    = "/static/image/icon/fox.svg"
)

type ModCopyStrartRed struct{}

func (copyStrartRed ModCopyStrartRed) GetModel() model.Model {
	action := func(myHistory []signal.Signal, oponentHistory []signal.Signal, aproximateInteractions int) signal.Signal {
		if len(oponentHistory) == 0 {
			return signal.Red
		}
		return oponentHistory[len(oponentHistory)-1]
	}
	return model.New(modelName, description, iconPath, action)
}

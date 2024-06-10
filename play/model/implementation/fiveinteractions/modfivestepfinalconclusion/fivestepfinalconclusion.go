package modfivestepfinalconclusion

import (
	"github.com/uszebr/thegamem/play/model"
	"github.com/uszebr/thegamem/play/model/modelutil"
	"github.com/uszebr/thegamem/play/signal"
)

const (
	modelName   = "fivestepfinalconclusion"
	description = "First 5 interactions replys with Green; If 3 of first 5 opponent replies are Red - switch to always Red; else always Green"
	iconPath    = "/static/image/icon/elk.svg"
)

type Modfivestepfinalconclusion struct{}

func (modfivestepfinalconclusion Modfivestepfinalconclusion) GetModel() model.Model {
	action := func(myHistory []signal.Signal, opponentHistory []signal.Signal, approximateInteractions int) signal.Signal {
		if len(opponentHistory) <= 5 {
			return signal.Green
		}
		greenQuantity := modelutil.CalcGreenInFirstNInteractions(opponentHistory, 5)
		if greenQuantity > 2 {
			return signal.Green
		}
		return signal.Red
	}
	return model.New(modelName, description, iconPath, action)
}

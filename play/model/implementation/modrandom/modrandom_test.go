package modrandom

import (
	"testing"

	"github.com/uszebr/thegamem/play/signal"
)

const (
	attemptsToGetRandom    = 30 // quantity to get "not the same" signal
	aproximateInteractions = 30
)

func TestModelRandom(t *testing.T) {
	tests := []struct {
		myHistory []signal.Signal
		opHistory []signal.Signal
	}{
		{myHistory: []signal.Signal{}, opHistory: []signal.Signal{}},
		// one signal in the histhory
		{myHistory: []signal.Signal{signal.Red}, opHistory: []signal.Signal{signal.Red}},
		{myHistory: []signal.Signal{signal.Red}, opHistory: []signal.Signal{signal.Green}},
		{myHistory: []signal.Signal{signal.Green}, opHistory: []signal.Signal{signal.Red}},
		{myHistory: []signal.Signal{signal.Green}, opHistory: []signal.Signal{signal.Green}},
		//two signal history
		{myHistory: []signal.Signal{signal.Red, signal.Red}, opHistory: []signal.Signal{signal.Red, signal.Red}},
		{myHistory: []signal.Signal{signal.Red, signal.Red}, opHistory: []signal.Signal{signal.Red, signal.Green}},
		{myHistory: []signal.Signal{signal.Green, signal.Green}, opHistory: []signal.Signal{signal.Green, signal.Red}},
		{myHistory: []signal.Signal{signal.Green, signal.Green}, opHistory: []signal.Signal{signal.Green, signal.Green}},
		//three
		{
			myHistory: []signal.Signal{signal.Red, signal.Red, signal.Green},
			opHistory: []signal.Signal{signal.Red, signal.Red, signal.Red},
		},
		{
			myHistory: []signal.Signal{signal.Red, signal.Red, signal.Red},
			opHistory: []signal.Signal{signal.Red, signal.Red, signal.Green},
		},
	}

	for index, test := range tests {
		model := ModRandom{}.GetModel()
		var quantityRed int = 0
		var quantityGreen int = 0
		for i := 0; i < attemptsToGetRandom; i++ {
			signalR := model.CalculateSignal(test.myHistory, test.opHistory, aproximateInteractions)
			if signalR == signal.Green {
				quantityGreen++
			}
			if signalR == signal.Red {
				quantityRed++
			}
		}

		if quantityGreen == 0 || quantityRed == 0 {
			t.Errorf("i: %d Both, quantity red: %v and quantity green: %v should be more then 0 for %v attempts; ", index, quantityRed, quantityGreen, attemptsToGetRandom)
		}
	}

}

package modcopystartred

import (
	"fmt"
	"testing"

	"github.com/uszebr/thegamem/play/signal"
)

const (
	aproximateInteractions = 30
)

func TestModelCopyStartGreen(t *testing.T) {
	tests := []struct {
		myHistory []signal.Signal
		opHistory []signal.Signal
		result    signal.Signal
	}{
		{myHistory: []signal.Signal{}, opHistory: []signal.Signal{}, result: signal.Red},
		// one signal in the histhory
		{myHistory: []signal.Signal{signal.Red}, opHistory: []signal.Signal{signal.Red}, result: signal.Red},
		{myHistory: []signal.Signal{signal.Red}, opHistory: []signal.Signal{signal.Green}, result: signal.Green},
		{myHistory: []signal.Signal{signal.Green}, opHistory: []signal.Signal{signal.Red}, result: signal.Red},
		{myHistory: []signal.Signal{signal.Green}, opHistory: []signal.Signal{signal.Green}, result: signal.Green},
		//two signal history
		{myHistory: []signal.Signal{signal.Red, signal.Red}, opHistory: []signal.Signal{signal.Red, signal.Red}, result: signal.Red},
		{myHistory: []signal.Signal{signal.Red, signal.Red}, opHistory: []signal.Signal{signal.Red, signal.Green}, result: signal.Green},
		{myHistory: []signal.Signal{signal.Green, signal.Green}, opHistory: []signal.Signal{signal.Green, signal.Red}, result: signal.Red},
		{myHistory: []signal.Signal{signal.Green, signal.Green}, opHistory: []signal.Signal{signal.Green, signal.Green}, result: signal.Green},
		//three
		{
			myHistory: []signal.Signal{signal.Red, signal.Red, signal.Green},
			opHistory: []signal.Signal{signal.Red, signal.Red, signal.Red},
			result:    signal.Red},
		{
			myHistory: []signal.Signal{signal.Red, signal.Red, signal.Red},
			opHistory: []signal.Signal{signal.Red, signal.Red, signal.Green},
			result:    signal.Green},

		//	fail to check
		//{myHistory: []signal.Signal{}, opHistory: []signal.Signal{}, result: signal.Green},
	}

	for index, testl := range tests {
		test := testl
		t.Run(fmt.Sprintf("Test: %v", index), func(t *testing.T) {
			model := ModCopyStrartRed{}.GetModel()
			actualResult := model.CalculateSignal(test.myHistory, test.opHistory, aproximateInteractions)

			if actualResult != test.result {
				t.Errorf("i: %d actual result: %v not as expected: %v; ", index, actualResult, test.result)
			}
		})
	}

}

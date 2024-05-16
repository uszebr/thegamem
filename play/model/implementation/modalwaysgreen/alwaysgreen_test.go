package modalwaysgreen

import (
	"testing"

	"github.com/uszebr/thegamem/play/signal"
)

const (
	aproximateInteractions = 30
)

func TestModelAlwaysGreen(t *testing.T) {
	tests := []struct {
		myHistory []signal.Signal
		opHistory []signal.Signal
		result    signal.Signal
	}{
		{myHistory: []signal.Signal{signal.Red}, opHistory: []signal.Signal{signal.Red}, result: signal.Green},
		{myHistory: []signal.Signal{signal.Red}, opHistory: []signal.Signal{signal.Green}, result: signal.Green},
		{myHistory: []signal.Signal{signal.Red}, opHistory: []signal.Signal{}, result: signal.Green},
		{myHistory: []signal.Signal{}, opHistory: []signal.Signal{}, result: signal.Green},

		//	{myHistory: []signal.Signal{signal.Red}, opHistory: []signal.Signal{}, result: signal.Red}, // fail to check
	}

	for index, test := range tests {
		model := ModAlwaysGreen{}.GetModel()
		actualResult := model.CalculateSignal(test.myHistory, test.opHistory, aproximateInteractions)

		if actualResult != test.result {
			t.Errorf("i: %d actual result: %v not as expected: %v; ", index, actualResult, test.result)
		}

	}

}

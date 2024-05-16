package modalwaysred

import (
	"testing"

	"github.com/uszebr/thegamem/play/signal"
)

const (
	aproximateInteractions = 30
)

func TestModelAlwaysRed(t *testing.T) {
	tests := []struct {
		myHistory []signal.Signal
		opHistory []signal.Signal
		result    signal.Signal
	}{
		{myHistory: []signal.Signal{signal.Red}, opHistory: []signal.Signal{signal.Red}, result: signal.Red},
		{myHistory: []signal.Signal{signal.Red}, opHistory: []signal.Signal{signal.Green}, result: signal.Red},
		{myHistory: []signal.Signal{}, opHistory: []signal.Signal{}, result: signal.Red},

		//	{myHistory: []signal.Signal{signal.Red}, opHistory: []signal.Signal{}, result: signal.Green}, // fail to check
	}

	for index, test := range tests {
		model := ModAlwaysRed{}.GetModel()
		actualResult := model.CalculateSignal(test.myHistory, test.opHistory, aproximateInteractions)

		if actualResult != test.result {
			t.Errorf("i: %d actual result: %v not as expected: %v; ", index, actualResult, test.result)
		}
	}

}

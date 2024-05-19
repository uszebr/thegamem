package signal

import (
	"testing"
)

func TestCalcScore(t *testing.T) {
	tests := []struct {
		mySignal Signal
		opSignal Signal
		result   int
	}{
		{mySignal: Green, opSignal: Green, result: 5},
		{mySignal: Green, opSignal: Red, result: -2},
		{mySignal: Red, opSignal: Green, result: 2},
		{mySignal: Red, opSignal: Red, result: 1},

		//	{mySignal: Green, opSignal: Green, result: 2}, // to fail
	}
	for index, test := range tests {
		actualResult := test.mySignal.CalcScore(test.opSignal)
		if actualResult != test.result {
			t.Errorf("i: %d actual result: %d not as expected: %d", index, actualResult, test.result)
		}
	}

	// Unknown signal
	expectLogFatal(t, func() {
		_ = Signal("unknown").CalcScore(Green)
	})

}

func expectLogFatal(t *testing.T, f func()) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Expected log.Fatal to terminate the program")
		}
	}()
	f()
}

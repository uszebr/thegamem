package modelutil

import "github.com/uszebr/thegamem/play/signal"

func CalcGreenInFirstNInteractions(opponentHistory []signal.Signal, firsNSignals int) int {
	if len(opponentHistory) < firsNSignals {
		panic("not enough opponent history signals")
	}
	result := 0
	for i := range firsNSignals {
		if opponentHistory[i] == signal.Green {
			result++
		}
	}
	return result
}

package modblindrevenge

import (
	"fmt"
	"testing"

	"github.com/uszebr/thegamem/play/signal"
)

const (
	aproximateInteractions = 30
)

func TestModelBlindRevenge(t *testing.T) {
	tests := []struct {
		myHistory []signal.Signal
		opHistory []signal.Signal
		result    signal.Signal
	}{
		{myHistory: []signal.Signal{}, opHistory: []signal.Signal{}, result: signal.Green},
		{myHistory: []signal.Signal{signal.Red}, opHistory: []signal.Signal{signal.Red}, result: signal.Red},
		{myHistory: []signal.Signal{signal.Green}, opHistory: []signal.Signal{signal.Red}, result: signal.Red},
		//two signal history
		{myHistory: []signal.Signal{signal.Red, signal.Red}, opHistory: []signal.Signal{signal.Red, signal.Red}, result: signal.Red},
		{myHistory: []signal.Signal{signal.Red, signal.Red}, opHistory: []signal.Signal{signal.Red, signal.Green}, result: signal.Red},
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
			result:    signal.Red},
		{
			myHistory: []signal.Signal{signal.Red, signal.Red, signal.Red},
			opHistory: []signal.Signal{signal.Red, signal.Green, signal.Green},
			result:    signal.Red},
		{
			myHistory: []signal.Signal{signal.Red, signal.Red, signal.Red},
			opHistory: []signal.Signal{signal.Green, signal.Green, signal.Green},
			result:    signal.Green},
		//four
		{
			myHistory: []signal.Signal{signal.Red, signal.Red, signal.Green, signal.Green},
			opHistory: []signal.Signal{signal.Red, signal.Red, signal.Red, signal.Red},
			result:    signal.Red},
		{
			myHistory: []signal.Signal{signal.Red, signal.Red, signal.Green, signal.Green},
			opHistory: []signal.Signal{signal.Green, signal.Red, signal.Red, signal.Red},
			result:    signal.Red},
		{
			myHistory: []signal.Signal{signal.Red, signal.Red, signal.Green, signal.Green},
			opHistory: []signal.Signal{signal.Green, signal.Red, signal.Red, signal.Green},
			result:    signal.Red},
		{
			myHistory: []signal.Signal{signal.Red, signal.Red, signal.Green, signal.Green},
			opHistory: []signal.Signal{signal.Green, signal.Red, signal.Green, signal.Green},
			result:    signal.Red},
		{
			myHistory: []signal.Signal{signal.Red, signal.Red, signal.Green, signal.Green},
			opHistory: []signal.Signal{signal.Red, signal.Green, signal.Green, signal.Green},
			result:    signal.Red},
		{
			myHistory: []signal.Signal{signal.Red, signal.Red, signal.Green, signal.Green},
			opHistory: []signal.Signal{signal.Green, signal.Green, signal.Red, signal.Green},
			result:    signal.Red},
		{
			myHistory: []signal.Signal{signal.Red, signal.Red, signal.Green, signal.Green},
			opHistory: []signal.Signal{signal.Green, signal.Green, signal.Green, signal.Red},
			result:    signal.Red},
		{
			myHistory: []signal.Signal{signal.Red, signal.Red, signal.Green, signal.Green},
			opHistory: []signal.Signal{signal.Green, signal.Green, signal.Green, signal.Green},
			result:    signal.Green},

		// to fail
		// {
		// 	myHistory: []signal.Signal{signal.Red, signal.Red, signal.Green, signal.Green},
		// 	opHistory: []signal.Signal{signal.Green, signal.Green, signal.Green, signal.Green},
		// 	result:    signal.Red},
	}

	for index, testl := range tests {
		test := testl
		t.Run(fmt.Sprintf("Test: %v", index), func(t *testing.T) {
			model := ModBlindRevenge{}.GetModel()
			actualResult := model.CalculateSignal(test.myHistory, test.opHistory, aproximateInteractions)
			if actualResult != test.result {
				t.Errorf("i: %d actual result: %v not as expected: %v; ", index, actualResult, test.result)
			}
		})
	}
}

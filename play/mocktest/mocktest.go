package mocktest

import (
	"github.com/stretchr/testify/mock"
	"github.com/uszebr/thegamem/play/signal"
)

// Mock for ModelI interface
type MockModel struct {
	mock.Mock
}

func (m *MockModel) CalculateSignal(myHistory []signal.Signal, opponentHistory []signal.Signal, aproximateInteractions int) signal.Signal {
	args := m.Called(myHistory, opponentHistory, aproximateInteractions)
	return args.Get(0).(signal.Signal)
}

func (m *MockModel) GetName() string {
	args := m.Called()
	return args.String(0)
}

func (m *MockModel) GetDescription() string {
	args := m.Called()
	return args.String(0)
}
func (m *MockModel) GetIcon() string {
	args := m.Called()
	return args.String(0)
}

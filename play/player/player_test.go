package player

import (
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/uszebr/thegamem/play/model/modelfactory"
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

// Test New function
func TestNew(t *testing.T) {
	mockModel := new(MockModel)
	playerInstance := New(mockModel)
	assert.NotNil(t, playerInstance, "player nil")
	assert.NotEmpty(t, playerInstance.name, "player name empty")
	assert.Equal(t, mockModel, playerInstance.modeli, "player model")
	assert.NotEqual(t, uuid.Nil, playerInstance.uuid, "player uuid nil")
}

// Test PlayOne method with real model
func TestPlayOneSmoke(t *testing.T) {
	factory := modelfactory.GetModelFactory()
	playerInstance := New(factory.MustCreateModel("alwaysgreen"))

	myHistory := []signal.Signal{signal.Green, signal.Green}       // Populate as needed
	opponentHistory := []signal.Signal{signal.Green, signal.Green} // Populate as needed
	aproximateInteractions := 5

	actualSignal := playerInstance.PlayOne(myHistory, opponentHistory, aproximateInteractions)
	expectedSignal := signal.Green
	if actualSignal != expectedSignal {
		t.Errorf("Play one signal: %v not as expected: %v", actualSignal, expectedSignal)
	}
}

// Test PlayOne method with mock model
func TestPlayOneMock(t *testing.T) {
	mockModel := new(MockModel)
	playerInstance := New(mockModel)

	myHistory := []signal.Signal{signal.Green, signal.Green}
	opponentHistory := []signal.Signal{signal.Red, signal.Red}
	aproximateInteractions := 5

	expectedSignal := signal.Green
	mockModel.On("CalculateSignal", myHistory, opponentHistory, aproximateInteractions).Return(expectedSignal)

	result := playerInstance.PlayOne(myHistory, opponentHistory, aproximateInteractions)

	assert.Equal(t, expectedSignal, result)
	mockModel.AssertExpectations(t)
}

func TestGetNameMock(t *testing.T) {
	mockModel := new(MockModel)
	playerInstance := New(mockModel)

	expectedModelName := "expectedmodel"
	mockModel.On("GetName").Return(expectedModelName)

	result := playerInstance.GetModelName()

	assert.Equal(t, expectedModelName, result, "model name")
	mockModel.AssertExpectations(t)
}

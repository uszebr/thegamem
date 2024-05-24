package model

import "github.com/uszebr/thegamem/play/signal"

type Model struct {
	name        string
	description string
	iconPath    string
	action      Action
}

type Action func(myHistory []signal.Signal, opponentHistory []signal.Signal, aproximateInteractions int) signal.Signal

func New(name string, description string, iconPath string, action Action) Model {
	return Model{name: name, description: description, iconPath: iconPath, action: action}
}

func (model Model) CalculateSignal(myHistory []signal.Signal, opponentHistory []signal.Signal, aproximateInteractions int) signal.Signal {
	return model.action(myHistory, opponentHistory, aproximateInteractions)
}

func (model Model) GetName() string {
	return model.name
}

func (model Model) GetDescription() string {
	return model.description
}

func (model Model) GetIcon() string {
	return model.iconPath
}

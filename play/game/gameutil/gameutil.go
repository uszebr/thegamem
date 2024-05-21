package gameutil

import (
	"fmt"
	"math/rand"

	"github.com/uszebr/thegamem/play/model/modelfactory"
	"github.com/uszebr/thegamem/play/player"
)

func ShufflePlayers(players []*player.Player) []*player.Player {
	copiedSlice := make([]*player.Player, len(players))
	copy(copiedSlice, players)
	for i := range players {
		j := rand.Intn(len(players)-i) + i
		copiedSlice[i], copiedSlice[j] = copiedSlice[j], copiedSlice[i]
	}
	return copiedSlice
}

// creating slice of players for each model in the map quantity of players
func GenerateByModelAndQuantity(models map[string]int) ([]*player.Player, error) {
	if len(models) == 0 {
		return []*player.Player{}, fmt.Errorf("generare players with empty model map")
	}

	players := make([]*player.Player, 0, len(models))
	factory := modelfactory.GetModelFactory()
	for modelName, quantity := range models {
		if quantity <= 0 {
			return []*player.Player{}, fmt.Errorf("wrong model: %v quantity: %v", modelName, quantity)
		}
		for range quantity {
			model := factory.MustCreateModel(modelName)
			player := player.New(model)
			players = append(players, player)
		}
	}
	return players, nil
}

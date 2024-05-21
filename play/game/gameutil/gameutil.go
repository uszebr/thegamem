package gameutil

import (
	"math/rand"

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

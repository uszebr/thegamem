package usergames

import (
	"sync"

	"github.com/uszebr/thegamem/play/game"
)

var (
	games *UserGames
	once  sync.Once
)

// Hold all games for all users
type UserGames struct {
	games map[string]*game.Game
}

func GetUserGames() *UserGames {
	once.Do(func() {
		games = new()
	})
	return games
}

func new() *UserGames {
	boards := UserGames{
		games: make(map[string]*game.Game),
	}
	return &boards
}

func (userGames UserGames) GetGameForUser(userId string) (*game.Game, bool) {
	value, ok := userGames.games[userId]
	return value, ok
}

func (userGames UserGames) AddGameForUser(userId string, game *game.Game) {
	userGames.games[userId] = game
}

package usergames

import (
	"fmt"
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
	userGames := UserGames{
		games: make(map[string]*game.Game),
	}
	return &userGames
}

func (userGames UserGames) GetGameForUser(userId string) (*game.Game, bool) {
	value, ok := userGames.games[userId]
	return value, ok
}

func (userGames UserGames) AddGameForUser(userId string, game *game.Game) {
	userGames.games[userId] = game
}

func (userGames UserGames) GetGameByUUID(uuid string) (*game.Game, error) {
	for _, game := range userGames.games {
		if game.GetUUID() == uuid {
			return game, nil
		}
	}
	return &game.Game{}, fmt.Errorf("Issue game for uuid: %v", uuid)
}

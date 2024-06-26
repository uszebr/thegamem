package urlservice

import (
	"fmt"

	"github.com/uszebr/thegamem/play/board"
	"github.com/uszebr/thegamem/play/game"
	"github.com/uszebr/thegamem/play/round"
)

func GetRoundUrl(game *game.Game, board *board.Board, round *round.Round) string {
	// todo add game part
	return fmt.Sprintf("/board/%v/round/%v", board.GetUUID(), round.GetUUID())
}

func GetBoardUrl(game *game.Game, board *board.Board) string {
	// todo add game part
	return fmt.Sprintf("/boards/%v", board.GetUUID())
}

func GetGameStatUrl(game *game.Game) string {
	return fmt.Sprintf("/stat/%v", game.GetUUID())
}

// todo refactor to const part url and add to thegamem.go
func GetModelsByBoardUrl(game *game.Game) string {
	return fmt.Sprintf("/modelquantitybyboards/%v", game.GetUUID())
}

// todo refactor to const part url and add to thegamem.go
func GetModelsDistributionUrl(game *game.Game) string {
	return fmt.Sprintf("/modeldistributionlastboard/%v", game.GetUUID())
}

// todo refactor to const part url and add to thegamem.go
func GetAllScoresByBoardUrl(game *game.Game) string {
	return fmt.Sprintf("/allscoresbyboards/%v", game.GetUUID())
}

func GetGameUrl(game *game.Game) string {
	//todo implement when move to full urls
	return "/game"
}

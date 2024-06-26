package game

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/uszebr/thegamem/play/board"
	"github.com/uszebr/thegamem/play/game/gameutil"
	"github.com/uszebr/thegamem/play/model/modelfactory"
	"github.com/uszebr/thegamem/play/player"
)

type Game struct {
	// initial data

	col          int
	row          int
	interactions int

	//Slice of boards for particular game. Next board added to the end
	boards         []*board.Board
	initialModels  []string
	rotation       int  // quantity of players winners/losers we need to remove/multiply after each board played and transfer to the next bord
	shufflePlayers bool //should we shuffle players between boards
	pairsCreatorI  board.PairsCreatorI
	uuid           uuid.UUID
}

// all parameters should be checked before.. after form submission
func New(col int, row int, interactions int, initialModels []string, pairCreator board.PairsCreatorI, rotation int, shufflePlayers bool) *Game {
	return &Game{row: row, col: col, interactions: interactions, initialModels: initialModels, pairsCreatorI: pairCreator, rotation: rotation, shufflePlayers: shufflePlayers, uuid: uuid.New()}
}

func (game *Game) AddNewBoard() error {
	var boardToAdd *board.Board
	if len(game.boards) == 0 {
		//initial board creation
		modelsWithQuantity := getInitialModelQuantities(game.col, game.row, game.initialModels)

		players, err := gameutil.GenerateByModelAndQuantity(modelsWithQuantity)
		if err != nil {
			return err
		}
		//initial players are always shuffled
		playersShuffled := gameutil.ShufflePlayers(players)
		boardToAdd, err = board.New(game.col, game.row, playersShuffled, game.interactions, game.pairsCreatorI, game.rotation)
		if err != nil {
			return err
		}

	} else {
		previousBoard := game.boards[len(game.boards)-1]
		playersWithCoordinatePrevious := previousBoard.GetPlayers()
		//	playersWithCoordinatesToAdd := board.MakeBasePlayers(game.col, game.row)
		//copy(playersWithCoordinatesToAdd, playersWithCoordinatePrevious) // copy of players matrix from previous board
		playersWithCoordinatesToAdd := make([][]*player.Player, len(playersWithCoordinatePrevious))
		for col, innerSlice := range playersWithCoordinatePrevious {
			newInnerSlice := make([]*player.Player, len(innerSlice))
			for row, levelTwo := range innerSlice {
				newInnerSlice[row] = levelTwo
			}
			playersWithCoordinatesToAdd[col] = newInnerSlice
		}
		winners := previousBoard.GetWinners()
		loosers := previousBoard.GetLoosers()

		factory := modelfactory.GetModelFactory()
		//itterating loosers and subtituting their position for winners copy
		for index, looser := range loosers {
			looserPosition, err := previousBoard.GetPositionForPlayer(looser.Player)
			if err != nil {
				return err
			}
			winnerCopy := player.New(factory.MustCreateModel(winners[index].Player.GetModelName()))
			playersWithCoordinatesToAdd[looserPosition.X][looserPosition.Y] = winnerCopy
		}
		playersWithCoordinatesToAddFlatList := gameutil.ConvertPlayerFlatList(playersWithCoordinatesToAdd)
		if game.shufflePlayers {
			playersWithCoordinatesToAddFlatList = gameutil.ShufflePlayers(playersWithCoordinatesToAddFlatList)
		}
		var err error
		boardToAdd, err = board.New(game.col, game.row, playersWithCoordinatesToAddFlatList, game.interactions, game.pairsCreatorI, game.rotation)
		if err != nil {
			return err
		}
	}
	game.boards = append(game.boards, boardToAdd)
	return nil
}

func (game *Game) GetBoards() []*board.Board {
	return game.boards
}

func (game *Game) GetBoardByUUID(uuid string) (*board.Board, error) {
	for _, b := range game.GetBoards() {
		if b.GetUUID() == uuid {
			return b, nil
		}
	}
	return &board.Board{}, fmt.Errorf("Issue Board for uuid: %v", uuid)
}

func (game *Game) GetCols() int {
	return game.col
}

func (game *Game) GetRows() int {
	return game.row
}

func (game *Game) GetInteractions() int {
	return game.interactions
}

func (game *Game) GetRotations() int {
	return game.rotation
}

func (game *Game) GetShuffle() bool {
	return game.shufflePlayers
}

func (game *Game) GetBoardsQuantity() int {
	return len(game.GetBoards())
}

func (game *Game) GetPairDescription() string {
	return game.pairsCreatorI.GetDescription()
}

func (game *Game) GetInitialModels() []string {
	return game.initialModels
}
func (game *Game) GetUUID() string {
	return game.uuid.String()
}

// calculating quantity of each model players(even distribution of all models)
func getInitialModelQuantities(cols int, rows int, models []string) map[string]int {
	var quantityOfEachModel int
	quantityOfEachModel = cols * rows / len(models)
	result := make(map[string]int)
	for _, model := range models {
		result[model] = quantityOfEachModel
	}
	return result
}

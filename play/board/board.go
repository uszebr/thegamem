package board

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/uszebr/thegamem/play/board/coordinate"
	"github.com/uszebr/thegamem/play/player"
	"github.com/uszebr/thegamem/play/round"
)

type PairsCreatorI interface {
	CreatePairs(boardCols, boardRows int) ([]coordinate.PositionPair, error)
}

type Board struct {
	cols         int // quantity of cols - X
	rows         int // quantity of rows - Y
	interactions int // quantity of interactions(provided by the game but might be changed in later versions)

	players [][]*player.Player
	rounds  []*round.Round
	pairs   []coordinate.PositionPair

	// model names/quantity
	models map[string]int
	uuid   uuid.UUID

	// how pairs of players are created for the rounds
	// pairsCreator PairsCreatorI
}

// if players need to be shuffled - do it before
func New(cols, rows int, initialPlayers []*player.Player, interactions int, pairsCreator PairsCreatorI) (*Board, error) {
	if interactions <= 0 {
		return &Board{}, fmt.Errorf("Creating board issue interactions: %v ", interactions)
	}
	if cols <= 0 || rows <= 0 {
		return &Board{}, fmt.Errorf("Creating board issue parameters col: %v row: %v ", cols, rows)
	}
	if len(initialPlayers) != cols*rows {
		return &Board{}, fmt.Errorf("Creating board parameters col: %v row: %v quantity of players: %v", cols, rows, len(initialPlayers))
	}

	var playersToSave = makeBasePlayers(cols, rows)
	indexIncoming := 0
	models := make(map[string]int)
	for i := 0; i < cols; i++ {
		for j := 0; j < rows; j++ {
			playerToInsert := initialPlayers[indexIncoming]
			//collecting models and their quantities
			currentModel := playerToInsert.GetModelName()
			_, ok := models[currentModel]
			if ok {
				models[currentModel]++
			} else {
				models[currentModel] = 1
			}
			playersToSave[i][j] = playerToInsert
			indexIncoming++
		}
	}
	// Creating Pairs with neighbors(might be extracted to constructor as interface if neede more pairs option(like all with all))
	pairs, err := pairsCreator.CreatePairs(cols, rows)
	if err != nil {
		return &Board{}, fmt.Errorf("Creating Pairs issue: %w", err)
	}
	board := &Board{
		uuid:         uuid.New(),
		cols:         cols,
		rows:         rows,
		players:      playersToSave,
		interactions: interactions,
		pairs:        pairs,
	}
	board.createRounds() //might be insert error/check from rounds if needed??

	return board, nil
}

// Getting players with dimension slices
func (board Board) GetPlayers() [][]*player.Player {
	return board.players
}

// Getting players as one level(no dimensions) slice
func (board Board) GetPlayersOneSlice() []*player.Player {
	multiDim := board.GetPlayers()
	result := make([]*player.Player, 0, board.rows*board.cols)
	for _, levelOne := range multiDim {
		for _, player := range levelOne {
			result = append(result, player)
		}
	}
	return result
}

func (board *Board) GetUUID() string {
	return board.uuid.String()
}

func (board Board) GetAllRounds() []*round.Round {
	return board.rounds
}

// getting only rounds that particular player participated Left or Right
// if nill return all rounds
func (board Board) GetRoundsForPlayer(player *player.Player) []*round.Round {
	if player == nil {
		return board.rounds
	}
	result := make([]*round.Round, 0)
	for _, round := range board.rounds {
		if round.Left.GetPlayer() == player || round.Right.GetPlayer() == player {
			result = append(result, round)
		}
	}
	return result
}

func (board Board) GetPlayerByPosition(position coordinate.Position) *player.Player {
	if position.X < 0 || position.Y < 0 {
		return nil
	}
	return board.GetPlayers()[position.X][position.Y]
}

func (board *Board) createRounds() {
	for _, pair := range board.pairs {
		round := round.New(board.GetPlayerByPosition(pair.Left), board.GetPlayerByPosition(pair.Right), board.interactions)
		board.rounds = append(board.rounds, &round)
	}
}

// returns two dimensional slice for players
func makeBasePlayers(cols, rows int) [][]*player.Player {
	players := make([][]*player.Player, cols)
	for i := range players {
		players[i] = make([]*player.Player, rows)
	}
	return players
}

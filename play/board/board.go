package board

import (
	"fmt"
	"sort"

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

	rotation int

	boardScores   map[*player.Player]int
	playerRatings []Rating
}

// this needed to show quickly on the page
// calculate once show always
// ?? todo might be substituted for method that calculates on a fly(less memory more processor)
type Rating struct {
	Player *player.Player
	Score  int
}

// if players need to be shuffled - do it before
func New(cols, rows int, initialPlayers []*player.Player, interactions int, pairsCreator PairsCreatorI, rotation int) (*Board, error) {
	if interactions <= 0 {
		return &Board{}, fmt.Errorf("Creating board issue interactions: %v ", interactions)
	}
	if cols <= 0 || rows <= 0 {
		return &Board{}, fmt.Errorf("Creating board issue parameters col: %v row: %v ", cols, rows)
	}
	if len(initialPlayers) != cols*rows {
		return &Board{}, fmt.Errorf("Creating board parameters col: %v row: %v quantity of players: %v", cols, rows, len(initialPlayers))
	}

	var playersToSave = MakeBasePlayers(cols, rows)
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
		boardScores:  make(map[*player.Player]int, cols*rows),
		//	playerRatings: make([]Rating, 0, cols*rows),
		rotation: rotation,
	}
	board.createRounds()         //might be insert error/check from rounds if needed??
	board.calculateBoardScores() // calculating scores for players
	board.calculateRating()      // calculating list of players rating for current board
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

func (board *Board) calculateBoardScores() {
	for _, round := range board.rounds {
		board.boardScores[round.Left.GetPlayer()] += round.Left.RoundScoreSum
		board.boardScores[round.Right.GetPlayer()] += round.Right.RoundScoreSum
	}
}

func (board *Board) GetBoardPlayerScores() map[*player.Player]int {
	return board.boardScores
}

func (board *Board) calculateRating() {
	ratings := make([]Rating, 0, board.cols*board.rows)
	for player, score := range board.boardScores {
		ratings = append(ratings, Rating{Player: player, Score: score})
	}
	sort.Slice(ratings, func(i, j int) bool {
		return ratings[i].Score > ratings[j].Score
	})
	board.playerRatings = ratings
}

func (board *Board) GetRatings() []Rating {
	return board.playerRatings
}

// returns two dimensional slice for players
func MakeBasePlayers(cols, rows int) [][]*player.Player {
	players := make([][]*player.Player, cols)
	for i := range players {
		players[i] = make([]*player.Player, rows)
	}
	return players
}

func (board *Board) GetWinners() []Rating {
	return board.playerRatings[:board.rotation]
}

func (board *Board) GetLoosers() []Rating {
	return board.playerRatings[len(board.playerRatings)-board.rotation:]
}

func (board *Board) GetPositionForPlayer(playerToFind *player.Player) (coordinate.Position, error) {
	for x, levelOne := range board.players {
		for y, player := range levelOne {
			if player == playerToFind {
				return coordinate.Position{X: x, Y: y}, nil
			}
		}
	}
	return coordinate.Position{}, fmt.Errorf("Player not found on the board")
}

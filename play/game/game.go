package game

import "github.com/uszebr/thegamem/play/board"

type Game struct {
	// initial data
	row          int
	col          int
	interactions int

	//Slice of boards for particular game. Next board added to the end
	boards []board.Board
}

func New(row int, col int, interactions int) *Game {
	return &Game{row: row, col: col, interactions: interactions}
}

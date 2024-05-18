package game

type Game struct {
	// initial data
	row          int
	col          int
	interactions int
}

func New(row int, col int, interactions int) *Game {
	return &Game{row: row, col: col, interactions: interactions}
}

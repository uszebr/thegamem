package signal

type Signal string

// Rules for Signal calculation might be extracted to the separate entity if needed many
const (
	green_green = 5 //4
	red_red     = 1
	red_green   = 2
	green_red   = -2
)

const (
	Green Signal = "cooperation"
	Red   Signal = "confrontation"
)

func (s Signal) CalcScore(opponentSignal Signal) int {
	if s == Green && opponentSignal == Green {
		return green_green
	}
	if s == Red && opponentSignal == Red {
		return red_red
	}
	if s == Red && opponentSignal == Green {
		return red_green
	}
	if s == Green && opponentSignal == Red {
		return green_red
	}
	panic("Unknown signal combination on CalcScore")
}

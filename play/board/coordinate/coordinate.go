package coordinate

import (
	"fmt"
)

// position/coordinates of player in the board
type Position struct {
	X int
	Y int
}

// to hold pairs of players with position
type PositionPair struct {
	Left  Position
	Right Position
}

func (position Position) GetPositinsAround(boardCols, boardRows int) ([]Position, error) {
	if position.Y >= boardRows || position.Y < 0 {
		return []Position{}, fmt.Errorf("players arround issue: y position is wrong: %v", position.Y)
	}
	if position.X >= boardCols || position.X < 0 {
		return []Position{}, fmt.Errorf("players arround issue: x position is wrong: %v", position.X)
	}

	xUp := coordinateTransform(position.X, boardCols, true)
	xDown := coordinateTransform(position.X, boardCols, false)
	yUp := coordinateTransform(position.Y, boardRows, true)
	yDown := coordinateTransform(position.Y, boardRows, false)
	resultMap := make(map[Position]bool)

	resultMap[Position{xUp, position.Y}] = true
	resultMap[Position{xDown, position.Y}] = true
	resultMap[Position{position.X, yDown}] = true
	resultMap[Position{position.X, yUp}] = true

	resultMap[Position{xUp, yUp}] = true
	resultMap[Position{xDown, yUp}] = true
	resultMap[Position{xUp, yDown}] = true
	resultMap[Position{xDown, yDown}] = true

	uniqueResult := make([]Position, 0, 8)
	for key := range resultMap {
		uniqueResult = append(uniqueResult, key)
	}
	return uniqueResult, nil
}

// Checking if position is in the slice of positions
func (position Position) In(positions []Position) bool {
	for _, p := range positions {
		if p.X == position.X && p.Y == position.Y {
			return true
		}
	}
	return false
}

// calculating coordinate transformation, bool - true increasing, false decreasing
// to simulate endless board
func coordinateTransform(coordinate int, boardSize int, direction bool) int {
	var result int
	if direction {
		result = coordinate + 1
		if result >= boardSize {
			result = 0
		}
	} else {
		result = coordinate - 1
		if result < 0 {
			result = boardSize - 1
		}
	}
	return result
}

// todo delete!! moved to pair package
// creating pairs of Positions for Rounds
// func CreatePairs(boardCols, boardRows int) ([]PositionPair, error) {
// 	result := make([]PositionPair, 0)
// 	calculatedPositions := make([]Position, 0, boardCols*boardRows)
// 	for x := range boardCols {
// 		for y := range boardRows {
// 			position := Position{x, y}
// 			calculatedPositions = append(calculatedPositions, position)
// 			partners, err := position.getPositinsAround(boardCols, boardRows)
// 			if err != nil {
// 				return []PositionPair{}, err
// 			}
// 			for _, partner := range partners {
// 				if !partner.In(calculatedPositions) {
// 					result = append(result, PositionPair{Left: position, Right: partner})
// 				}
// 			}
// 		}
// 	}
// 	return result, nil
// }

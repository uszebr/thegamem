package newgamehandler

import (
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/uszebr/thegamem/play/board"
	"github.com/uszebr/thegamem/play/board/pair"
)

type GameCreating struct {
	colsInput      int
	rowsInput      int
	interactions   int
	rotation       int
	pairCreator    board.PairsCreatorI
	shufflePlayers bool
	modelsQuantity int
	modelsInput    []string
}

// Checking form params for wrong values or values incompatibility
func newGameFormValuesChecker(c echo.Context) (GameCreating, []string) {
	issues := []string{}
	colsInputParam := c.FormValue("columnsinput")
	rowsInputParam := c.FormValue("rowsinput")
	interactionsParam := c.FormValue("interactions")
	rotationParam := c.FormValue("rotation")
	neighboursParam := c.FormValue("neighbours")
	shuffleParam := c.FormValue("shufflePlayers")
	formParams, err := c.FormParams()
	if err != nil {
		issues = append(issues, "Can not get Form Parameters")
	}
	modelsParam := formParams["modelname"]
	interactions, err := strconv.Atoi(interactionsParam)
	if err != nil {
		issues = append(issues, "Error interactions input: ")
	}
	if interactions <= 0 {
		issues = append(issues, "Should be more then 0 interactions")
	}

	// all check form values here..
	colsInput, err := strconv.Atoi(colsInputParam)
	if err != nil {
		issues = append(issues, "Error columns input: ")
	}
	if colsInput <= 2 {
		issues = append(issues, "Should be more then 2 columns")
	}
	rowsInput, err := strconv.Atoi(rowsInputParam)
	if err != nil {
		issues = append(issues, "Error rows input")
	}
	if rowsInput <= 2 {
		issues = append(issues, "Should be more then 2 rows")
	}
	rotation, err := strconv.Atoi(rotationParam)
	if err != nil {
		issues = append(issues, "Error rotation input")
	}
	if rotation < 0 {
		issues = append(issues, "Should be 0 or more player rotation")
	}
	// condition to duplicate winner players only once and avoid multiple duplications for some winners
	// should be less then half of all players rotation
	if rotation > rowsInput*colsInput/2 {
		issues = append(issues, "Should be rotations less then half of all players")
	}

	var pairCreator board.PairsCreatorI
	if neighboursParam == "neighbours" {
		pairCreator = pair.PairsNeighbour{}
	} else if neighboursParam == "allplayers" {
		pairCreator = pair.PairAll{}
	} else {
		issues = append(issues, "Player Pairs parameter issue")
	}

	var shufflePlayers bool
	if shuffleParam == "yes" {
		shufflePlayers = true
	} else if shuffleParam == "no" {
		shufflePlayers = false
	} else {
		issues = append(issues, "Player Shuffling parameter issue")
	}

	modelsInput := modelsParam
	modelsQuantity := len(modelsInput)

	if modelsQuantity < 2 {
		issues = append(issues, "Should be more then 2 models selected")
	} else {
		if (colsInput*rowsInput)%modelsQuantity != 0 {
			issues = append(issues, "the number of cells must be divided without a remainder by the number of models for equal distribution between models")
		}
	}

	return GameCreating{
		colsInput:      colsInput,
		rowsInput:      rowsInput,
		interactions:   interactions,
		rotation:       rotation,
		pairCreator:    pairCreator,
		shufflePlayers: shufflePlayers,
		modelsQuantity: modelsQuantity,
		modelsInput:    modelsInput,
	}, issues
}

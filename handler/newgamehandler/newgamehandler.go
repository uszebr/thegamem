package newgamehandler

import (
	"log/slog"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/uszebr/thegamem/handler/utilhandler"
	"github.com/uszebr/thegamem/internal/entity"
	"github.com/uszebr/thegamem/play/board"
	"github.com/uszebr/thegamem/play/board/coordinate"
	"github.com/uszebr/thegamem/play/game"
	"github.com/uszebr/thegamem/play/model/modelfactory"
	"github.com/uszebr/thegamem/play/usergames"
	"github.com/uszebr/thegamem/view/boardview"
	"github.com/uszebr/thegamem/view/component/cardview"
	"github.com/uszebr/thegamem/view/component/fullpageview"
	"github.com/uszebr/thegamem/view/gameview"
	"github.com/uszebr/thegamem/view/newgameview"
)

type NewGameHadler struct {
	usergames    *usergames.UserGames
	modelFactory *modelfactory.ModelFactory
}

func New(usergames *usergames.UserGames, modelFactory *modelfactory.ModelFactory) NewGameHadler {
	return NewGameHadler{usergames: usergames, modelFactory: modelFactory}
}

func (h *NewGameHadler) HandleShow(c echo.Context) error {

	ctx := c.Request().Context()
	user, ok := ctx.Value("user").(entity.UserAuth)
	if !ok {
		return utilhandler.Render(c, fullpageview.FullPageWithError("Access denied", "Access denied", "Need to be logged in to use boards"))
	}
	_, ok = h.usergames.GetGameForUser(user.UserId)
	return utilhandler.Render(c, newgameview.Show(user, h.modelFactory.GetAllModelNames(), ok, []string{}))
}

func (h *NewGameHadler) HandlePost(c echo.Context) error {

	ctx := c.Request().Context()
	user, ok := ctx.Value("user").(entity.UserAuth)
	if !ok {
		return utilhandler.Render(c, cardview.ShowDangerCart("Access denied", "Access denied"))
	}

	//	_, gameOk := h.usergames.GetGameForUser(user.UserId)
	//todo
	//getting form values
	// checking form values .. send form with errors
	//EDiT!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!
	newGameCreating, issues := newGameFormValuesChecker(c)

	if len(issues) > 0 {
		return utilhandler.Render(c, newgameview.GameCreateForm(h.modelFactory.GetAllModelNames(), issues))
	}

	gameFresh := game.New(
		newGameCreating.colsInput,
		newGameCreating.rowsInput,
		newGameCreating.interactions,
		newGameCreating.modelsInput,
		newGameCreating.pairCreator,
		newGameCreating.rotation,
		newGameCreating.shufflePlayers,
	)

	h.usergames.AddGameForUser(user.UserId, gameFresh)
	c.Response().Header().Set("HX-Redirect", "/game")
	return c.NoContent(http.StatusOK)
}

func (h *NewGameHadler) HandleExistingGame(c echo.Context) error {

	ctx := c.Request().Context()
	user, ok := ctx.Value("user").(entity.UserAuth)
	if !ok {
		return utilhandler.Render(c, fullpageview.FullPageWithError("Access denied", "Access denied", "Need to be logged in to use boards"))
	}
	game, ok := h.usergames.GetGameForUser(user.UserId)
	if !ok {
		return utilhandler.Render(c, fullpageview.FullPageWithError("No Game Found", "No Game", "No game found. Plz create new game to start."))
	}
	return utilhandler.Render(c, gameview.Show(user, h.modelFactory.GetAllModelNames(), game))
}

func (h *NewGameHadler) HandleAddBoardPost(c echo.Context) error {
	slog.Debug("Add Board Post")
	ctx := c.Request().Context()
	user, ok := ctx.Value("user").(entity.UserAuth)
	if !ok {
		return utilhandler.Render(c, cardview.ShowDangerCart("Access denied", "Need to be logged in to use boards"))
	}
	game, ok := h.usergames.GetGameForUser(user.UserId)
	if !ok {
		return utilhandler.Render(c, cardview.ShowDangerCart("No Game Found", "No game found. Plz create new game to start."))
	}
	err := game.AddNewBoard()
	if err != nil {
		return utilhandler.Render(c, cardview.ShowDangerCart("Game Error", "Creating new board issue"))
	}
	activeParam := c.FormValue("active")

	active, err := strconv.Atoi(activeParam)
	if err != nil {
		active = -1
	}
	return utilhandler.Render(c, gameview.BoardsPanel(game.GetBoards(), active))
}

func (h *NewGameHadler) HandleBoard(c echo.Context) error {
	slog.Debug("Particular Board Get")
	ctx := c.Request().Context()
	user, ok := ctx.Value("user").(entity.UserAuth)
	if !ok {
		return utilhandler.Render(c, fullpageview.FullPageWithError("Access denied", "Board issue: Access denied", "Need to be logged in to use boards"))
	}
	game, ok := h.usergames.GetGameForUser(user.UserId)
	if !ok {
		return utilhandler.Render(c, fullpageview.FullPageWithError("No Game Found", "Board issue: No Game", "No game found. Plz create new game to start."))
	}
	boardUrl := c.Param("id")

	for i, b := range game.GetBoards() {
		if b.GetUUID() == boardUrl {
			return utilhandler.Render(c, boardview.Show(b, i, game))
		}
	}
	return utilhandler.Render(c, fullpageview.FullPageWithError("No Board Found", "Board issue: No Board", "No board found: "+boardUrl))
}

func (h *NewGameHadler) HandleRoundsForPlayerPost(c echo.Context) error {
	slog.Debug("Rounds for player Post")
	ctx := c.Request().Context()
	user, ok := ctx.Value("user").(entity.UserAuth)
	if !ok {
		return utilhandler.Render(c, cardview.ShowDangerCart("Access denied", "Need to be logged in to use rounds for player"))
	}
	game, ok := h.usergames.GetGameForUser(user.UserId)
	if !ok {
		return utilhandler.Render(c, cardview.ShowDangerCart("No Game Found", "No game found. Plz create new game to start and play rounds."))
	}

	colPost := c.FormValue("col")
	rowPost := c.FormValue("row")
	boardPost := c.FormValue("board")

	col, err := strconv.Atoi(colPost)
	if err != nil {
		return utilhandler.Render(c, cardview.ShowDangerCart("Round for player issue", "Can not get rounds, column reques issue"))
	}
	row, err := strconv.Atoi(rowPost)
	if err != nil {
		return utilhandler.Render(c, cardview.ShowDangerCart("Round for player issue", "Can not get rounds, rows reques issue"))
	}

	var board *board.Board
	for _, b := range game.GetBoards() {
		if b.GetUUID() == boardPost {
			board = b
			break
		}
	}
	if board == nil {
		return utilhandler.Render(c, cardview.ShowDangerCart("Round for player issue", "Board not found: "+boardPost))
	}
	position := coordinate.Position{X: col, Y: row}
	player := board.GetPlayerByPosition(position)
	//logger.Debug("TEMP HandleRoundsForPlayer", "col", col, "row", row)
	return utilhandler.Render(c, boardview.PlayerRounds(board, player))
}

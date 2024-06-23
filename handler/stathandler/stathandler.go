package stathandler

import (
	"log/slog"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/uszebr/thegamem/handler/utilhandler"
	"github.com/uszebr/thegamem/internal/chart/modelsbyboardchart"
	"github.com/uszebr/thegamem/play/model/modelfactory"
	"github.com/uszebr/thegamem/play/usergames"
	"github.com/uszebr/thegamem/view/component/cardview"
	"github.com/uszebr/thegamem/view/component/fullpageview"
	"github.com/uszebr/thegamem/view/statview"
)

type StatHadler struct {
	usergames    *usergames.UserGames
	modelFactory *modelfactory.ModelFactory
}

func New(usergames *usergames.UserGames, modelFactory *modelfactory.ModelFactory) StatHadler {
	return StatHadler{usergames: usergames, modelFactory: modelFactory}
}

func (h *StatHadler) HandleStat(c echo.Context) error {
	slog.Debug("Handle stat")

	gameUrl := c.Param("gameId")
	game, err := h.usergames.GetGameByUUID(gameUrl)
	if err != nil {
		return utilhandler.Render(c, fullpageview.FullPageWithError("No Game Found", "Stat issue: No Game", "No game found or this game does not exist: "+gameUrl))
	}
	//todo cashing for game.. board quantity??

	return utilhandler.Render(c, statview.Show(game))
}

func (h *StatHadler) ModelsByBoardChart(c echo.Context) error {
	slog.Info("Chart SCript Post")
	gameUrl := c.Param("gameId")
	game, err := h.usergames.GetGameByUUID(gameUrl)
	if err != nil {
		return utilhandler.Render(c, cardview.ShowDangerCart("No Game Found", "No game found or this game does not exist: "+gameUrl))
	}
	// modelQuantities := []entity.ModelQuantity{
	// 	{Name: "alwaysgreen", Data: []int{14, 61, 75, 41, 46, 62, 44, 97, 48}},
	// }
	// readyString := `[ 10, 41, 35, 51, 49, 62, 69, 91, 148]`
	// modelsByBoardData := entity.ModelsByBoardChartData{
	// 	LineData: readyString,
	// }
	//return utilhandler.Render(c, modelsbyboardview.ChartScript(modelsByBoardData))
	return c.HTML(http.StatusOK, modelsbyboardchart.ModelsByBoardChart{}.GetChartScript(game))
}

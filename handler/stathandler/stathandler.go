package stathandler

import (
	"log/slog"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/uszebr/thegamem/handler/utilhandler"
	"github.com/uszebr/thegamem/internal/chart/allscoresbyboardchart"
	"github.com/uszebr/thegamem/internal/chart/modeldistributionchart"
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

	modelsByBoardChart     *modelsbyboardchart.ModelsByBoardChart
	modelDistributionChart *modeldistributionchart.ModelsDistributionChart
	allScoresByBoardChart  *allscoresbyboardchart.AllScoresByBoardChart
}

func New(usergames *usergames.UserGames, modelFactory *modelfactory.ModelFactory) StatHadler {
	return StatHadler{
		usergames:              usergames,
		modelFactory:           modelFactory,
		modelsByBoardChart:     modelsbyboardchart.New(),
		modelDistributionChart: modeldistributionchart.New(),
		allScoresByBoardChart:  allscoresbyboardchart.New()}
}

func (h *StatHadler) HandleStat(c echo.Context) error {
	slog.Debug("Handle stat")

	gameUrl := c.Param("gameId")
	game, err := h.usergames.GetGameByUUID(gameUrl)
	if err != nil {
		return utilhandler.Render(c, fullpageview.FullPageWithError("No Game Found", "Stat issue: No Game", "No game found or this game does not exist: "+gameUrl))
	}
	return utilhandler.Render(c, statview.Show(game))
}

func (h *StatHadler) ModelsByBoardChart(c echo.Context) error {
	gameUrl := c.Param("gameId")
	game, err := h.usergames.GetGameByUUID(gameUrl)
	if err != nil {
		return utilhandler.Render(c, cardview.ShowDangerCart("No Game Found", "No game found or this game does not exist: "+gameUrl))
	}
	return c.HTML(http.StatusOK, h.modelsByBoardChart.GetChartScript(game))
}

func (h *StatHadler) ModelDistributionLastBoardChart(c echo.Context) error {
	gameUrl := c.Param("gameId")
	game, err := h.usergames.GetGameByUUID(gameUrl)
	if err != nil {
		return utilhandler.Render(c, cardview.ShowDangerCart("No Game Found", "No game found or this game does not exist: "+gameUrl))
	}
	return c.HTML(http.StatusOK, h.modelDistributionChart.GetChartScript(game))
}

func (h *StatHadler) AllScoresByBoardChart(c echo.Context) error {
	gameUrl := c.Param("gameId")
	game, err := h.usergames.GetGameByUUID(gameUrl)
	if err != nil {
		return utilhandler.Render(c, cardview.ShowDangerCart("No Game Found", "No game found or this game does not exist: "+gameUrl))
	}
	return c.HTML(http.StatusOK, h.allScoresByBoardChart.GetChartScript(game))
}

package newgamehandler

import (
	"github.com/labstack/echo/v4"
	"github.com/uszebr/thegamem/handler/utilhandler"
	"github.com/uszebr/thegamem/internal/entity"
	"github.com/uszebr/thegamem/play/model/modelfactory"
	"github.com/uszebr/thegamem/play/usergames"
	"github.com/uszebr/thegamem/view/component/fullpageview"
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
	return utilhandler.Render(c, newgameview.Show(user, h.modelFactory.GetAllModelNames(), ok))
}

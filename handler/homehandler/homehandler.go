package homehandler

import (
	"github.com/labstack/echo/v4"
	"github.com/uszebr/thegamem/handler/utilhandler"
	"github.com/uszebr/thegamem/view/homeview"
)

type HomeHandler struct {
}

func (h *HomeHandler) HandleShow(c echo.Context) error {
	return utilhandler.Render(c, homeview.Show())
}

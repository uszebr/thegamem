package utilhandler

import (
	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
)

// Util to render templ file with Echo framework context
func Render(c echo.Context, component templ.Component) error {
	return component.Render(c.Request().Context(), c.Response())
}

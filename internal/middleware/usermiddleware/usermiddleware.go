package usermiddleware

import (
	"context"

	"github.com/labstack/echo/v4"
	"github.com/uszebr/thegamem/handler/utilhandler"
	"github.com/uszebr/thegamem/internal/authservice"
	"github.com/uszebr/thegamem/view/component/fullpageview"
)

type UserMiddleware struct {
	auth *authservice.AuthService
}

func New(auth *authservice.AuthService) UserMiddleware {
	return UserMiddleware{auth: auth}
}

func (um UserMiddleware) LoggedIn(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		user, err := um.auth.GetSignedInUser(c)
		if err == nil {
			ctx := context.WithValue(c.Request().Context(), "user", user)
			c.SetRequest(c.Request().WithContext(ctx))
			return next(c)
		}
		return utilhandler.Render(c, fullpageview.FullPageWithError("Access denied Game Theory", "Access denied", "Need to be logged in to use boards"))
	}
}

// todo refacor!!!!!!!!!!!!!!!!!!!!!!!!!!!
func (um UserMiddleware) GetUserForPublic(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		// here is redirect
		//return c.Redirect(http.StatusTemporaryRedirect, "/")
		user, err := um.auth.GetSignedInUser(c)

		//slog.Debug("Temp middleware: ", "user", user)
		if err == nil {
			ctx := context.WithValue(c.Request().Context(), "user", user)
			c.SetRequest(c.Request().WithContext(ctx))
		}
		return next(c)
	}
}

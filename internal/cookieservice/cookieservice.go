package cookieservice

import (
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

func DeleteCookies(c echo.Context, cookieNames []string) {
	for _, name := range cookieNames {
		cookie := &http.Cookie{
			Name:     name,
			Value:    "",
			Path:     "/",
			Expires:  time.Unix(0, 0),
			MaxAge:   -1,
			HttpOnly: true,
		}
		c.SetCookie(cookie)
	}
}

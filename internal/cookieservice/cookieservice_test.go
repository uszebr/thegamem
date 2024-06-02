package cookieservice

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestDeleteCookies(t *testing.T) {
	e := echo.New()

	// Create a request to pass to the Echo context
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	cookieNames := []string{"cookie0", "cookie1", "cookie2", "cookie3"}
	// Set initial cookies in the request
	for _, name := range cookieNames {
		cookie := &http.Cookie{
			Name:  name,
			Value: "initial_value",
		}
		req.AddCookie(cookie)
	}

	DeleteCookies(c, cookieNames)

	// Verify that the cookies are deleted
	for _, name := range cookieNames {
		// Echo's context stores the response cookies, so we need to get them from the recorder
		cookies := rec.Result().Cookies()

		var deletedCookie *http.Cookie
		for _, cookie := range cookies {
			if cookie.Name == name {
				deletedCookie = cookie
				break
			}
		}

		//checking that actual cookie is found
		assert.NotNil(t, deletedCookie, "expected cookie to be deleted but it was not found")

		assert.Equal(t, name, deletedCookie.Name)
		assert.Equal(t, "", deletedCookie.Value)
		assert.Equal(t, -1, deletedCookie.MaxAge)
		assert.True(t, deletedCookie.HttpOnly)
	}
}

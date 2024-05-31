package loginhandler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/uszebr/thegamem/handler/utilhandler"
	"github.com/uszebr/thegamem/internal/authservice"
	"github.com/uszebr/thegamem/internal/entity"
	"github.com/uszebr/thegamem/view/component/cardview"
	"github.com/uszebr/thegamem/view/loginview"
)

type LoginHandler struct {
	authservice *authservice.AuthService
}

func New(authservice *authservice.AuthService) LoginHandler {
	return LoginHandler{authservice: authservice}
}

func (h *LoginHandler) HandleShow(c echo.Context) error {
	return utilhandler.Render(c, loginview.Show([]error{}))
}

func (h *LoginHandler) HandlePost(c echo.Context) error {
	req := c.Request()
	//fmt.Println(req.Header.Get("HX-Request"))//true
	pass := req.PostFormValue("password")
	email := req.PostFormValue("email")
	//todo check values
	_, err := h.authservice.SignIn(c, email, pass)
	if err != nil {
		return utilhandler.Render(c, loginview.LoginForm([]error{err}))
	}

	return utilhandler.Render(c, loginview.SuccessLogin())
}

func (h *LoginHandler) LogoutPost(c echo.Context) error {

	ctx := c.Request().Context()
	if _, ok := ctx.Value("user").(entity.UserAuth); ok {
		return utilhandler.Render(c, cardview.ShowDangerCart("Logout Issue", "NOT logged in."))
	}

	acookie, err := c.Cookie("atoken")
	//no tokens found
	if err != nil || acookie == nil {
		return utilhandler.Render(c, cardview.ShowDangerCart("Logout Issue", "NOT logged in."))
	}

	atoken := acookie.Value
	if err := h.authservice.SignOut(c, atoken); err != nil {
		return utilhandler.Render(c, cardview.ShowDangerCart("Logout Issue", "Supa logout problem"))
	}
	c.Response().Header().Set("HX-Redirect", "/login")
	return c.NoContent(http.StatusOK)
	// return utilhandler.Render(c, loginview.SuccessLogout())
}

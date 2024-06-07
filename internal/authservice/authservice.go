package authservice

import (
	"context"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/uszebr/thegamem/internal/cookieservice"
	"github.com/uszebr/thegamem/internal/entity"
)

const (
	delayCookielifetime = 10
)

type AuthService struct {
	client AuthClientI //for now only supabase
	jwt    JWTServiceI // for now only jwtservice
}

func New(authClient AuthClientI, jwtService JWTServiceI) *AuthService {
	return &AuthService{client: authClient, jwt: jwtService}
}

type AuthClientI interface {
	RefreshUser(ctx context.Context, userToken string, refreshToken string) (*entity.AuthDetails, error)
	SignIn(ctx context.Context, userLogin string, userPass string) (*entity.AuthDetails, error)
	SignOut(ctx context.Context, userTocken string) error
}
type JWTServiceI interface {
	GetUserFromJWT(tokenExt string) (entity.UserAuth, error)
}

func (service *AuthService) SetAuthCookies(c echo.Context, details *entity.AuthDetails) error {
	if details == nil {
		return fmt.Errorf("Setting Cookies issue, details is missing")
	}
	if details.AccessToken == "" || details.RefreshToken == "" {
		return fmt.Errorf("Setting Cookies issue, Access or Refresh token is missing")
	}
	maxAge := 0
	if details.ExpiresIn-delayCookielifetime > 0 {
		maxAge = details.ExpiresIn - delayCookielifetime
	}
	accessCookie := &http.Cookie{
		Name:     "atoken",
		Path:     "/",
		MaxAge:   maxAge,
		Value:    details.AccessToken,
		HttpOnly: true,
	}
	c.SetCookie(accessCookie)
	refreshCookie := &http.Cookie{
		Name:     "rtoken",
		Path:     "/",
		Value:    details.RefreshToken,
		HttpOnly: true,
	}
	c.SetCookie(refreshCookie)
	return nil
}

func (service *AuthService) SignOut(c echo.Context, userTocken string) error {
	if err := service.client.SignOut(c.Request().Context(), userTocken); err != nil {
		return err
	}
	cookieservice.DeleteCookies(c, []string{"atoken", "rtoken"})
	return nil
}

func (service *AuthService) SignIn(c echo.Context, userLogin string, userPass string) (*entity.AuthDetails, error) {
	authDetails, err := service.client.SignIn(c.Request().Context(), userLogin, userPass)
	if err != nil {
		return authDetails, err
	}
	if err = service.SetAuthCookies(c, authDetails); err != nil {
		return &entity.AuthDetails{}, err
	}
	return authDetails, nil
}

func (service *AuthService) GetSignedInUser(c echo.Context) (entity.UserAuth, error) {
	acookie, err1 := c.Cookie("atoken")
	rcookie, err2 := c.Cookie("rtoken")
	//no tokens found
	if err1 != nil && err2 != nil {
		return entity.UserAuth{}, fmt.Errorf("No access and refresh cookies")
	}

	// access token present && not empty
	if err1 == nil {
		atoken := acookie.Value
		currentUser, err := service.jwt.GetUserFromJWT(atoken)
		if err == nil {
			return currentUser, nil
		}
	}
	// refresh token present && not empty
	if err2 == nil {
		rtoken := rcookie.Value
		refreshUserDetails, err := service.RefreshTokenAndSetCookie(c, rtoken)
		if err == nil {
			ruser, err := service.jwt.GetUserFromJWT(refreshUserDetails.AccessToken)
			if err == nil {
				return ruser, nil
			}
		}
	}
	cookieservice.DeleteCookies(c, []string{"atoken", "rtoken"})
	return entity.UserAuth{}, fmt.Errorf("Can not get SingedIn user: End")
}

func (service *AuthService) RefreshTokenAndSetCookie(c echo.Context, rtoken string) (*entity.AuthDetails, error) {
	detailsAfreRefresh, err := service.client.RefreshUser(c.Request().Context(), "", rtoken)
	//todo remove concrete implementation
	//supa.GetClient().Auth.RefreshUser(c.Request().Context(), "", rtoken)
	if err != nil {
		return &entity.AuthDetails{}, err
	}

	err = service.SetAuthCookies(c, detailsAfreRefresh)
	if err != nil {
		return &entity.AuthDetails{}, err
	}
	return detailsAfreRefresh, nil
}

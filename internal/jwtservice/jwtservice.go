package jwtservice

import (
	"fmt"
	"log/slog"

	"github.com/dgrijalva/jwt-go"
	"github.com/uszebr/thegamem/internal/entity"
	"github.com/uszebr/thegamem/internal/logger/loggerinit"
	"github.com/uszebr/thegamem/internal/logger/logutil"
)

// example of decoding https://jwt.io/
func GetUserFromJWT(tokenExt string) (entity.UserAuth, error) {
	token, _, err := new(jwt.Parser).ParseUnverified(tokenExt, jwt.MapClaims{})
	if err != nil {
		slog.Error("Error parsing JWT", logutil.Err(err))
		return entity.UserAuth{}, fmt.Errorf("Error parsing JWT: %v", err)
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		loggerinit.GetLogger().Error("Invalid JWT claims")
		return entity.UserAuth{}, fmt.Errorf("Invalid JWT claims")
	}

	// Extracting data
	email, emailOk := claims["email"].(string)
	sub, subOk := claims["sub"].(string)
	role, roleOk := claims["role"].(string)
	isAnonymous, isAnonymousOk := claims["is_anonymous"].(bool)
	// Check if all required fields are present
	if !emailOk || !subOk || !isAnonymousOk || !roleOk {
		return entity.UserAuth{}, fmt.Errorf("Missing required fields in JWT claims")
	}
	return entity.UserAuth{Email: email, Role: role, IsAnonymus: isAnonymous, UserId: sub}, nil
}

package jwtservice

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUserFromJWT(t *testing.T) {
	tokenToTest := `eyJhbGciOiJIUzI1NiIsImtpZCI6IlNsR2loZkN3U00zTG5vUzAiLCJ0eXAiOiJKV1QifQ.eyJhdWQiOiJhdXRoZW50aWNhdGVkIiwiZXhwIjoxNzEzNjg2MzI1LCJpYXQiOjE3MTM2ODI3MjUsImlzcyI6Imh0dHBzOi8vbXJ5Z29zdXphcHRnaHFjeXhicGkuc3VwYWJhc2UuY28vYXV0aC92MSIsInN1YiI6Ijk2ZGIwMjI4LTBmMWYtNDRlNy1iMjgxLTdjZWY5MDI3MDVmZCIsImVtYWlsIjoibXlAbXkubXkiLCJwaG9uZSI6IiIsImFwcF9tZXRhZGF0YSI6eyJwcm92aWRlciI6ImVtYWlsIiwicHJvdmlkZXJzIjpbImVtYWlsIl19LCJ1c2VyX21ldGFkYXRhIjp7fSwicm9sZSI6ImF1dGhlbnRpY2F0ZWQiLCJhYWwiOiJhYWwxIiwiYW1yIjpbeyJtZXRob2QiOiJwYXNzd29yZCIsInRpbWVzdGFtcCI6MTcxMzY4MjcyMn1dLCJzZXNzaW9uX2lkIjoiNTQ3ZTUyMmEtM2JmOC00MWMzLWFmOTYtY2FkNDA0OWNkZjcyIiwiaXNfYW5vbnltb3VzIjpmYWxzZX0.zLxCglfM6mGN62vThlBHtSTICoTIJYmB4XZV7SER7tk`
	jwtservice := JwtService{}
	user, err := jwtservice.GetUserFromJWT(tokenToTest)

	expectedEmail := "my@my.my"
	expecedId := "96db0228-0f1f-44e7-b281-7cef902705fd"
	expecedRole := "authenticated"
	assert.Nil(t, err)
	assert.Equal(t, expectedEmail, user.Email)
	assert.Equal(t, expecedId, user.UserId)
	assert.Equal(t, expecedRole, user.Role)
	assert.False(t, user.IsAnonymus)
}

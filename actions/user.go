package actions

import (
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gobuffalo/buffalo"
)

var signingKey = []byte("mysecretkey")

// UserLogin default implementation.
func UserLogin(c buffalo.Context) error {

	claims := jwt.StandardClaims{
		ExpiresAt: time.Now().Add(time.Duration(24*7) * time.Hour).Unix(),
		Id:        "2",
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(signingKey)
	if err != nil {
		return fmt.Errorf("could not sign token, %v", err)
	}
	return c.Render(200, r.JSON(map[string]string{"token": tokenString}))
}

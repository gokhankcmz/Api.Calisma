package Token

import (
	"Api.Calisma/src/Common/Models/ErrorModels"
	"Api.Calisma/src/Common/Models/RequestModels"
	"Api.Calisma/src/CustomerService/Constants"
	"fmt"
	"github.com/golang-jwt/jwt"
	"strings"
	"time"
)

type CustomClaims struct {
	ID    string `json:"ID"`
	Email string `json:"email"`
	jwt.StandardClaims
}

func CreateToken(tc *RequestModels.TokenCredentials) string {
	claims := CustomClaims{
		ID:    tc.ID,
		Email: tc.Email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(Constants.JWTExpTime).Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &claims)
	t, err := token.SignedString([]byte(Constants.JWTSecretKey))
	if err != nil {
		fmt.Println(err)
	}
	return t
}

func ValidateAndGetClaims(tokenString string) CustomClaims {
	tokenString = strings.Replace(tokenString, "Bearer ", "", 1)
	var claims CustomClaims
	_, err := jwt.ParseWithClaims(tokenString, &claims, func(t *jwt.Token) (interface{}, error) {
		return []byte(Constants.JWTSecretKey), nil
	})

	if err != nil {
		if err.Error() == "signature is invalid" {
			panic(ErrorModels.InvalidToken.SetPublicDetail("Invalid Signature."))
		}
	}

	if claims.ExpiresAt < time.Now().Unix() {
		panic(ErrorModels.InvalidToken.SetPublicDetail("Token Expired."))
	}
	return claims
}

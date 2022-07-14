package authorization

import (
	"crud_psql_7/model"
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"time"
)

// GenerateToken .
func GenerateToken(data *model.Login) (string, error) {
	claim := model.Claim{
		Email: data.Email,
		StandardClaims: jwt.StandardClaims{
			Audience:  "audicence",
			ExpiresAt: time.Now().Add(time.Hour * 2).Unix(),
			Id:        fmt.Sprintf("%s", data.Email),
			IssuedAt:  time.Now().Unix(),
			Issuer:    "chrisloarryn",
			NotBefore: 0,
			Subject:   "token",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claim)

	signedToken, err := token.SignedString(signKey)
	if err != nil {
		return "", err
	}

	return signedToken, nil
}

// ValidateToken .
func ValidateToken(t string) (model.Claim, error) {
	token, err := jwt.ParseWithClaims(t, &model.Claim{}, verifyFunction)
	if err != nil {
		return model.Claim{}, err
	}

	if !token.Valid {
		return model.Claim{}, errors.New("token is not valid")
	}

	claim, ok := token.Claims.(*model.Claim)
	if !ok {
		return model.Claim{}, errors.New("could not get claims")
	}

	return *claim, nil
}

func verifyFunction(t *jwt.Token) (interface{}, error) {
	return verifyKey, nil
}

package authorization

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"time"

	"class5/api/model"
)

// GenerateToken .
func GenerateToken(data *model.Login) (string, error) {
	claim := model.Claim{
		Email: data.Email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 2).Unix(),
			Issuer:    "go.micro.srv.user",
		},
	}
	//,			ExpiresAt: time.Now().Add(time.Hour * 2).Unix(),
	//			Issuer:    "EDteam",
	//		},

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
		return model.Claim{}, errors.New("token no válido")
	}

	claim, ok := token.Claims.(*model.Claim)
	if !ok {
		return model.Claim{}, errors.New("no se pudo obtener los claim")
	}

	return *claim, nil
}

func verifyFunction(t *jwt.Token) (interface{}, error) {
	return verifyKey, nil
}

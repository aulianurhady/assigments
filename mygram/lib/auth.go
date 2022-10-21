package lib

import (
	"errors"
	"strings"
)

func Auth(authorization, username, email string) error {
	tokenString := strings.Replace(authorization, "Bearer ", "", -1)

	claimUsername, claimEmail, err := ClaimJWT(tokenString)
	if err != nil {
		return errors.New("Cannot authorization")
	}

	if claimUsername != username || claimEmail != email {
		return errors.New("Cannot authorization")
	}

	return nil
}

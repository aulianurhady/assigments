package lib

import (
	"errors"
	"strings"
)

func Auth(authorization string) error {
	tokenString := strings.Replace(authorization, "Bearer ", "", -1)

	claimUserID, claimUsername, claimEmail, err := ClaimJWT(tokenString)
	if err != nil {
		return errors.New("Cannot authorization")
	}

	if claimUserID != 0 || claimUsername != "" || claimEmail != "" {
		return errors.New("Cannot authorization")
	}

	return nil
}

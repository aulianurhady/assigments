package lib

import (
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var (
	JWT_SIGNATURE_KEY  = []byte("my_secret_key")
	JWT_SIGNING_METHOD = jwt.SigningMethodHS256

	username string
	email    string
)

type Claims struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	jwt.StandardClaims
}

func BuildJWT(username, email string) (string, error) {
	expirationTime := time.Now().Add(5 * time.Minute)

	claims := &Claims{
		Username: username,
		Email:    email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(JWT_SIGNATURE_KEY)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func ClaimJWT(tokenString string) (string, string, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if method, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("Signing method invalid")
		} else if method != JWT_SIGNING_METHOD {
			return nil, errors.New("Signing method invalid")
		}

		return JWT_SIGNATURE_KEY, nil
	})

	if err != nil {
		return "", "", errors.New("Token cannot claim")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return "", "", errors.New("Token cannot claim")
	}

	setClaimData(claims["username"].(string), claims["email"].(string))

	return claims["username"].(string), claims["email"].(string), nil
}

func setClaimData(claimUsername, claimEmail string) {
	username, email = claimUsername, claimEmail
}

func GetUsernameFromClaim() string {
	return username
}

func GetEmailFromClaim() string {
	return email
}

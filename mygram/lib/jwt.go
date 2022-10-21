package lib

import (
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var (
	JWT_SIGNATURE_KEY  = []byte("my_secret_key")
	JWT_SIGNING_METHOD = jwt.SigningMethodHS256

	userID   int
	username string
	email    string
)

type Claims struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	UserID   int    `json:"user_id"`
	jwt.StandardClaims
}

func BuildJWT(id int, username, email string) (string, error) {
	expirationTime := time.Now().Add(5 * time.Minute)

	claims := &Claims{
		UserID:   id,
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

func ClaimJWT(tokenString string) (int, string, string, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if method, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("Signing method invalid")
		} else if method != JWT_SIGNING_METHOD {
			return nil, errors.New("Signing method invalid")
		}

		return JWT_SIGNATURE_KEY, nil
	})

	if err != nil {
		return 0, "", "", errors.New("Token cannot claim")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return 0, "", "", errors.New("Token cannot claim")
	}

	setClaimData(claims["user_id"].(int), claims["username"].(string), claims["email"].(string))

	return claims["user_id"].(int), claims["username"].(string), claims["email"].(string), nil
}

func setClaimData(claimUserID int, claimUsername, claimEmail string) {
	userID, username, email = claimUserID, claimUsername, claimEmail
}

func GetUserIDFromClaim() *int {
	return &userID
}

func GetUsernameFromClaim() *string {
	return &username
}

func GetEmailFromClaim() *string {
	return &email
}

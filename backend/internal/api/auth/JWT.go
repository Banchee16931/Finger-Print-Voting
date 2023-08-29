package auth

import (
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
)

type Claims struct {
	jwt.StandardClaims
	Username string `json:"username"`
}

const JWTIssuedAtOffset = -1 * time.Hour
const JWTValidityDuration = 12 * time.Hour

func (claims Claims) Valid() error {
	if claims.Username == "" {
		return fmt.Errorf("missing username")
	}

	if time.Now().Unix() > claims.ExpiresAt {
		return fmt.Errorf("token has expired")
	}

	return nil
}

func GenerateJWT(secret string, username string) (string, error) {
	secretKey := []byte(secret)

	claims := Claims{
		Username: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(JWTValidityDuration).Unix(),
			IssuedAt:  time.Now().Add(JWTIssuedAtOffset).Unix(),
			Subject:   username,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", fmt.Errorf("cannot sign token: %w", err)
	}

	return tokenString, nil
}

func GetClaims(secret string, encryotedToken string) (Claims, error) {
	token, err := jwt.ParseWithClaims(encryotedToken, &Claims{}, keyFunc(secret))
	if err != nil {
		return Claims{}, fmt.Errorf("%w: %s", ErrParse, err.Error())
	}

	if !token.Valid {
		return Claims{}, fmt.Errorf("invalid token")
	}

	claims, ok := token.Claims.(*Claims)
	if !ok {
		return Claims{}, fmt.Errorf("claims could not be retrieved")
	}

	return *claims, nil
}

func keyFunc(secret string) jwt.Keyfunc {
	secretKey := []byte(secret)
	return func(t *jwt.Token) (interface{}, error) {
		return secretKey, nil
	}
}

var ErrParse = errors.New("failed to parse token")

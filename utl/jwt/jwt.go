package jwt

import (
	"errors"

	"github.com/dgrijalva/jwt-go"
)

var (
	ErrSigningToken   = errors.New("Error signing token")
	ErrTokenInvalid   = errors.New("Token is invalid")
	ErrTokenExpired   = errors.New("Token is expired")
	ErrTokenMalformed = errors.New("Token is malformed")
	ErrTokenNotValid  = errors.New("Token is not valid yet")
)

func GenerateToken(claims jwt.Claims) (string, error) {
	signingKey := []byte("hey")

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(signingKey)
	if err != nil {
		return "", ErrSigningToken
	}

	return tokenString, nil
}

func ParseToken(tokenString string) (*jwt.Token, error) {
	signingKey := []byte("hey")

	parsedToken, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return signingKey, nil
	})

	if err != nil {
		if e, ok := err.(*jwt.ValidationError); ok {
			switch {
			case e.Errors&jwt.ValidationErrorMalformed != 0:
				return nil, ErrTokenMalformed
			case e.Errors&jwt.ValidationErrorExpired != 0:
				return nil, ErrTokenExpired
			case e.Errors&jwt.ValidationErrorNotValidYet != 0:
				return nil, ErrTokenNotValid
			case e.Inner != nil:
				return nil, e.Inner
			}
		}
		return nil, err
	}

	if !parsedToken.Valid {
		return nil, ErrTokenInvalid
	}

	return parsedToken, nil
}

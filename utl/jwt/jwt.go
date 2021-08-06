package jwt

import (
	"errors"

	"github.com/dgrijalva/jwt-go"
)

var (
	ErrSigningToken            = errors.New("Error signing token")
	ErrTokenInvalid            = errors.New("Token is invalid")
	ErrTokenExpired            = errors.New("Token is expired")
	ErrTokenMalformed          = errors.New("Token is malformed")
	ErrTokenNotValid           = errors.New("Token is not valid yet")
	ErrUnexpectedSigningMethod = errors.New("Unexpected signing method")
)

type TokenGenerator interface {
	GenerateToken(jwt.Claims) (string, error)
}

type TokenParser interface {
	ParseToken(string) (*jwt.Token, error)
}

type service struct {
	key    []byte
	method jwt.SigningMethod
}

func New(key string, method jwt.SigningMethod) service {
	return service{
		key:    []byte(key),
		method: method,
	}
}

func (s service) GenerateToken(claims jwt.Claims) (string, error) {
	token := jwt.NewWithClaims(s.method, claims)

	tokenString, err := token.SignedString(s.key)
	if err != nil {
		return "", ErrSigningToken
	}

	return tokenString, nil
}

func (s service) ParseToken(tokenString string) (*jwt.Token, error) {
	parsedToken, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if token.Method != s.method {
			return nil, ErrUnexpectedSigningMethod
		}

		return s.key, nil
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

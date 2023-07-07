package recipeservice

import (
	"crypto/sha256"
	"errors"
	"restipe/internal/model"
	"restipe/internal/storage"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type AuthService struct {
	storage   storage.Authorization
	signToken string
}

func NewAuthService(storage storage.Authorization, jwttoken string) *AuthService {
	return &AuthService{storage, jwttoken}
}

func (s *AuthService) SignupUser(user model.SignupUser) (int, error) {
	user.Password = hashPassword(user.Password)
	return s.storage.SignupUser(user)
}

type tokenClaimsWithId struct {
	jwt.StandardClaims
	UserId int `json:"user_id"`
}

func (s *AuthService) SigninUser(user model.SigninUser) (string, error) {
	user.Password = hashPassword(user.Password)
	id, err := s.storage.SigninUser(user)
	if err != nil {
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaimsWithId{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(24 * time.Hour).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		id,
	})

	return token.SignedString([]byte(s.signToken))
}

func (s *AuthService) Authorize(token string) (int, error) {
	parsedToken, err := jwt.ParseWithClaims(token, &tokenClaimsWithId{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}
		return []byte(s.signToken), nil
	})

	if err != nil {
		return 0, err
	}

	claims, ok := parsedToken.Claims.(*tokenClaimsWithId)
	if !ok {
		return 0, errors.New("token claims have wrong type")
	}

	return claims.UserId, nil
}

func hashPassword(password string) string {
	hash := sha256.New()
	hash.Write([]byte(password))

	return string(hash.Sum(nil))

}

package helper

import (
	"crypto/rand"
	"encoding/base32"
	"errors"
	"github.com/golang-jwt/jwt"
	"github.com/volkankocaali/e-commorce-go/pkg/config"
	"github.com/volkankocaali/e-commorce-go/pkg/schema"
	"golang.org/x/crypto/bcrypt"
	"time"
)

type Helper struct {
	cfg config.Config
}

func NewHelper(config *config.Config) *Helper {
	return &Helper{
		cfg: *config,
	}
}

type AutoCustomClaims struct {
	Id    uint   `json:"id"`
	Email string `json:"email"`
	Role  string `json:"role"`
	jwt.StandardClaims
}

func (helper *Helper) GenerateTokenClients(user schema.UserSchemaResponse) (string, error) {
	claims := &AutoCustomClaims{
		Id:    user.Id,
		Email: user.Email,
		Role:  "client",
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Minute * 1000).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte("ecommercegoapplication"))

	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func (helper *Helper) GeneratePasswordHash(password string) (string, error) {
	hashedPass, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", errors.New("internal server error")
	}

	return string(hashedPass), nil
}

func (helper *Helper) GenerateReferralCode() (string, error) {
	b := make([]byte, 16)
	_, err := rand.Read(b)

	if err != nil {
		return "", err
	}

	encoded := base32.NewEncoding("ABCDEFGHIJKLMNOPQRSTUVWXYZ234567").WithPadding(base32.NoPadding).EncodeToString(b)
	return encoded[:6], nil
}

func (helper *Helper) CompareHashAndPassword(a string, b string) error {
	err := bcrypt.CompareHashAndPassword([]byte(a), []byte(b))
	if err != nil {
		return err
	}
	return nil
}

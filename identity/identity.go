package identity

import (
	"strconv"
	"time"

	"github.com/birdbox/authnz/config"
	"github.com/birdbox/authnz/models"
	"github.com/golang-jwt/jwt/v5"
)

type IdentityClaims struct {
	SessionID string `json:"sid"`
	Name      string `json:"name"`
	jwt.RegisteredClaims
}

func (c *IdentityClaims) Sign(secret []byte) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	return token.SignedString(secret)
}

func Parse(tokenStr string, secret []byte) (*IdentityClaims, error) {
	token, err := jwt.ParseWithClaims(tokenStr, &IdentityClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})
	if err != nil {
		return nil, err
	}
	return token.Claims.(*IdentityClaims), nil
}

func NewIdentityClaims(sessionID string, user *models.User, cfg *config.Config) *IdentityClaims {
	maxAge := cfg.IdentityToken.MaxAge
	issuedAt := jwt.NewNumericDate(time.Now())
	expiresAt := jwt.NewNumericDate(time.Now().Add(time.Duration(maxAge) * time.Second))
	fullName := user.FirstName + " " + user.LastName

	return &IdentityClaims{
		sessionID,
		fullName,
		jwt.RegisteredClaims{
			ExpiresAt: expiresAt,
			IssuedAt:  issuedAt,
			NotBefore: issuedAt,
			Issuer:    cfg.Server.Host,
			Subject:   strconv.Itoa(user.ID),
			Audience:  []string{"somebody_else"},
		},
	}
}

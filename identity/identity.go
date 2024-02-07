package identity

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/matthiase/warden/config"
	"github.com/matthiase/warden/models"
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
			Subject:   user.ID,
			Audience:  []string{"somebody_else"},
		},
	}
}

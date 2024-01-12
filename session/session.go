package session

import (
	"time"

	"github.com/birdbox/authnz/config"
	"github.com/golang-jwt/jwt/v5"
)

var scope = "authentication"

type SessionClaims struct {
	Scope string `json:"scope"`
	jwt.RegisteredClaims
}

func (s *SessionClaims) Sign(secret []byte) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, s)
	return token.SignedString(secret)
}

func Parse(tokenStr string, secret []byte) (*SessionClaims, error) {
	token, err := jwt.ParseWithClaims(tokenStr, &SessionClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})
	if err != nil {
		return nil, err
	}
	return token.Claims.(*SessionClaims), nil
}

func NewSessionClaims(sessionID string, cfg *config.Config) *SessionClaims {
	maxAge := cfg.Session.MaxAge
	issuer := cfg.Server.Host
	issuedAt := jwt.NewNumericDate(time.Now())
	expiresAt := jwt.NewNumericDate(time.Now().Add(time.Duration(maxAge) * time.Second))

	return &SessionClaims{
		scope,
		jwt.RegisteredClaims{
			ExpiresAt: expiresAt,
			IssuedAt:  issuedAt,
			NotBefore: issuedAt,
			Issuer:    issuer,
			Subject:   sessionID,
			Audience:  []string{"somebody_else"},
		},
	}
}

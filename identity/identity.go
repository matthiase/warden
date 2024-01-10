package identity

import (
	"strconv"
	"time"

	"github.com/birdbox/authnz/config"
	"github.com/birdbox/authnz/session"
	"github.com/golang-jwt/jwt/v5"
)

var scope = "authentication"

type IdentityClaims struct {
	Scope     string `json:"scope"`
	SessionID string `json:"session_id"`
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

func NewIdentityClaims(userID int, session *session.SessionClaims, cfg *config.Config) *IdentityClaims {
	maxAge := cfg.AccessToken.MaxAge
	issuer := cfg.Server.Host
	issuedAt := jwt.NewNumericDate(time.Now())
	expiresAt := jwt.NewNumericDate(time.Now().Add(time.Duration(maxAge) * time.Second))

	return &IdentityClaims{
		scope,
		session.Subject,
		jwt.RegisteredClaims{
			ExpiresAt: expiresAt,
			IssuedAt:  issuedAt,
			NotBefore: issuedAt,
			Issuer:    issuer,
			Subject:   strconv.Itoa(userID),
			Audience:  []string{"somebody_else"},
		},
	}
}

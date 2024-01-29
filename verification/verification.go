package verification

import (
	"strconv"
	"time"

	"github.com/birdbox/authnz/config"
	"github.com/golang-jwt/jwt/v5"
)

type VerificationClaims struct {
	jwt.RegisteredClaims
}

func (c *VerificationClaims) Sign(secret []byte) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	return token.SignedString(secret)
}

func Parse(tokenStr string, secret []byte) (*VerificationClaims, error) {
	token, err := jwt.ParseWithClaims(tokenStr, &VerificationClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})
	if err != nil {
		return nil, err
	}
	return token.Claims.(*VerificationClaims), nil
}

func NewVerificationClaims(userID int, cfg *config.Config) *VerificationClaims {
	maxAge := cfg.VerificationToken.MaxAge
	issuer := cfg.Server.Host
	issuedAt := jwt.NewNumericDate(time.Now())
	expiresAt := jwt.NewNumericDate(time.Now().Add(time.Duration(maxAge) * time.Second))

	return &VerificationClaims{
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

package security

import (
	"context"
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type JWTIssuer struct {
	secret []byte
	issuer string
	now    func() time.Time
}

func NewJWTIssuer(secret, issuer string) *JWTIssuer {
	return &JWTIssuer{
		secret: []byte(secret),
		issuer: issuer,
		now:    func() time.Time { return time.Now().UTC() },
	}
}

func (j *JWTIssuer) Issue(ctx context.Context, subject string, claims map[string]interface{}, expiresIn time.Duration) (string, error) {
	now := j.now()

	std := jwt.MapClaims{
		"sub": subject,
		"iss": j.issuer,
		"iat": now.Unix(),
		"exp": now.Add(expiresIn).Unix(),
	}
	for k, v := range claims {
		std[k] = v
	}

	t := jwt.NewWithClaims(jwt.SigningMethodHS256, std)
	return t.SignedString(j.secret)
}

func (j *JWTIssuer) ParseAndVerify(ctx context.Context, token string) (map[string]interface{}, error) {
	parser := jwt.NewParser(
		jwt.WithValidMethods([]string{jwt.SigningMethodHS256.Name}),
		jwt.WithIssuer(j.issuer),
		jwt.WithLeeway(30*time.Second),
	)

	claims := jwt.MapClaims{}
	tok, err := parser.ParseWithClaims(token, claims, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return j.secret, nil
	})
	if err != nil {
		return nil, err
	}
	if !tok.Valid {
		return nil, errors.New("invalid token")
	}

	return map[string]interface{}(claims), nil
}

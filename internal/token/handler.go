package token

import (
	"errors"

	"github.com/bdreece/notable/pkg/server/config"
	"github.com/golang-jwt/jwt/v5"
)

var (
	ErrSign   = errors.New("failed to sign token")
	ErrVerify = errors.New("failed to verify token")
)

type Signer interface {
	Sign(*Claims) (string, error)
}

type Verifier interface {
	Verify(string) (*Claims, error)
}

type Handler struct {
	cfg *config.Token
}

// Sign implements Signer.
func (h *Handler) Sign(claims *Claims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString([]byte(h.cfg.Secret))
	if err != nil {
		return "", errors.Join(err, ErrSign)
	}

	return ss, nil
}

// Verify implements Verifier.
func (h *Handler) Verify(token string) (*Claims, error) {
	t, err := jwt.ParseWithClaims(token, &Claims{},
		func(t *jwt.Token) (interface{}, error) {
			return []byte(h.cfg.Secret), nil
		})

	if err != nil {
		return nil, errors.Join(err, ErrVerify)
	}

	return t.Claims.(*Claims), nil
}

func NewHandler(cfg *config.Token) *Handler {
	return &Handler{cfg}
}

var _ Signer = (*Handler)(nil)
var _ Verifier = (*Handler)(nil)

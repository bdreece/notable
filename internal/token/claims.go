package token

import "github.com/golang-jwt/jwt/v5"

type Claims struct {
    jwt.RegisteredClaims
}

var _ jwt.Claims = (*Claims)(nil)

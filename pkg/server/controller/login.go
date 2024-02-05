package controller

import (
	"fmt"
	"log/slog"
	"net/http"
	"time"

	"github.com/bdreece/notable/internal/token"
	"github.com/bdreece/notable/pkg/server/config"
	"github.com/bdreece/notable/pkg/server/database"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

type loginRequest struct {
	Email    string `form:"email"`
	Password string `form:"password"`
}

type loginResponse struct {
	AccessToken string    `json:"access_token"`
	Expiration  time.Time `json:"expiration"`
}

func Login(
	db database.Querier,
	signer token.Signer,
    cfg *config.Token,
	logger *slog.Logger,
) echo.HandlerFunc {
	return func(c echo.Context) error {
		r := new(loginRequest)
		if err := c.Bind(r); err != nil {
            logger.Error("failed to bind request", slog.String("error", err.Error()))
            return echo.NewHTTPError(http.StatusBadRequest)
		}

        logger.Info("querying user by email", slog.String("email", r.Email))
        u, err := db.FindUserByEmail(c.Request().Context(), r.Email)
        if err != nil {
            logger.Error("failed to find user", slog.String("err", err.Error()))
            return echo.NewHTTPError(http.StatusUnauthorized)
        }

        logger.Info("validating password attempt")
        err = bcrypt.CompareHashAndPassword(u.Hash, []byte(r.Password))
        if err != nil {
            logger.Error("failed to validate password", slog.String("err", err.Error()))
            return echo.NewHTTPError(http.StatusUnauthorized)
        }

        logger.Info("user authenticated, creating token")
        now := time.Now()
        t, err := signer.Sign(&token.Claims{
            RegisteredClaims: jwt.RegisteredClaims{
                ID: uuid.New().String(),
                Audience: jwt.ClaimStrings{cfg.Audience},
                Issuer: cfg.Issuer,
                Subject: u.ID.URN(),
                NotBefore: jwt.NewNumericDate(now),
                IssuedAt: jwt.NewNumericDate(now),
                ExpiresAt: jwt.NewNumericDate(now.Add(6 * time.Hour)),
            },
        })
        
        if err != nil {
            logger.Error("failed to create token", slog.String("err", err.Error()))
            return echo.NewHTTPError(http.StatusInternalServerError)
        }

        c.Response().Header().Add("Set-Cookie",
            fmt.Sprintf("notable-access-token=%s; HttpOnly; Path=/; SameSite=Strict; Expires=%s",
                t, now.Add(6 * time.Hour).Format(time.RFC1123Z)))

        return c.NoContent(http.StatusOK)
	}
}

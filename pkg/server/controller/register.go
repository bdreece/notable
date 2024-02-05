package controller

import (
	"log/slog"
	"net/http"

	"github.com/bdreece/notable/pkg/server/database"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

type registerRequest struct {
	FirstName       string `form:"firstName"`
	LastName        string `form:"lastName"`
	Email           string `form:"email"`
	Password        string `form:"password"`
	ConfirmPassword string `form:"confirmPassword"`
}

func Register(db database.Querier, logger *slog.Logger) echo.HandlerFunc {
    return func(c echo.Context) error {
        r := new(registerRequest)
        if err := c.Bind(r); err != nil || r.Password != r.ConfirmPassword {
            logger.Error("failed to bind request",
                slog.String("error", err.Error()))

            return echo.NewHTTPError(http.StatusBadRequest)
        }

        logger.Info("checking if user exists", slog.String("email", r.Email))
        count, err := db.UserExistsByEmail(c.Request().Context(), r.Email)
        if err != nil {
            logger.Error("failed to check if user exists",
                slog.String("error", err.Error()))
            return echo.NewHTTPError(http.StatusInternalServerError)
        } else if count > 0 {
            logger.Warn("user already exists")
            return echo.NewHTTPError(http.StatusConflict)
        }

        logger.Info("creating user...")
        hash, err := bcrypt.GenerateFromPassword([]byte(r.Password), bcrypt.DefaultCost)
        if err != nil {
            return echo.NewHTTPError(http.StatusInternalServerError)
        }

        u := database.InsertUserParams{
            FirstName: r.FirstName,
            LastName: r.LastName,
            EmailAddress: r.Email,
            Hash: hash,
        }

        if err = db.InsertUser(c.Request().Context(), u); err != nil {
            logger.Error("failed to create user", slog.String("error", err.Error()))
            return echo.NewHTTPError(http.StatusInternalServerError)
        }

        return c.NoContent(http.StatusCreated)
    }
}

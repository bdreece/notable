package controller

import (
	"log/slog"
	"net/http"

	"github.com/labstack/echo/v4"
)

func View(name string, data any, logger *slog.Logger) echo.HandlerFunc {
    return func(c echo.Context) error {
        if err := c.Render(200, name, data); err != nil {
            logger.Error("failed to render view",
                slog.String("error", err.Error()))

            return echo.NewHTTPError(http.StatusInternalServerError)
        }

        return nil
    }
}

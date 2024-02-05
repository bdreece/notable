package router

import (
	"log/slog"

	"github.com/bdreece/notable/internal/token"
	"github.com/bdreece/notable/pkg/server/config"
	"github.com/bdreece/notable/pkg/server/controller"
	"github.com/bdreece/notable/pkg/server/database"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func New(
    renderer echo.Renderer,
    db database.Querier,
    signer token.Signer,
    cfg *config.Config,
    logger *slog.Logger,
) *echo.Echo {
    e := echo.New()
    e.Renderer = renderer

    e.Use(middleware.Recover())
    e.Static("/", "web/ui/dist")

    e.GET("/", controller.View("home.gotmpl", struct{}{},
        logger.With(slog.String("route", "/"))))
    e.GET("/features", controller.View("features.gotmpl", struct{}{},
        logger.With(slog.String("route", "/features"))))

    e.GET("/login", controller.View("login.gotmpl", struct{}{},
        logger.With(slog.String("route", "/login"))))
    e.POST("/login", controller.Login(db, signer, &cfg.Token,
        logger.With(slog.String("route", "/login"))))

    e.GET("/register", controller.View("register.gotmpl", struct{}{},
        logger.With(slog.String("route", "/register"))))
    e.POST("/register", controller.Register(db,
        logger.With(slog.String("route", "/register"))))
    
    e.GET("/forgot-password", controller.View("forgot-password.gotmpl", struct{}{},
        logger.With(slog.String("route", "/forgot-password"))))

    return e
}

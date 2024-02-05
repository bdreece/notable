package server

import (
	"database/sql"
	"fmt"
	"log/slog"
	"net/http"
	"os"

	"github.com/bdreece/notable/internal/logging"
	"github.com/bdreece/notable/internal/renderer"
	"github.com/bdreece/notable/internal/token"
	"github.com/bdreece/notable/pkg/server/config"
	"github.com/bdreece/notable/pkg/server/database"
	"github.com/bdreece/notable/pkg/server/options"
	"github.com/bdreece/notable/pkg/server/router"
	"github.com/bdreece/notable/pkg/server/storage"
	"github.com/labstack/echo/v4"
	_ "github.com/lib/pq"
	"go.uber.org/dig"
	"google.golang.org/grpc"
)

type Builder struct {
	*dig.Container
}

func (b *Builder) Build() *App {
	b.provideConfig().
		provideLogger().
		provideRenderer().
		provideToken().
		provideDatabase().
		provideStorage().
		provideHttp().
		provideGrpc()

	ch := make(chan *App, 1)
	defer close(ch)

	if err := b.Invoke(func(
		http *http.Server,
		grpc *grpc.Server,
		logger *slog.Logger,
		opts *options.Options,
	) {
		ch <- &App{
			http:   servlet{http, opts.HttpPort},
			grpc:   servlet{grpc, opts.GrpcPort},
			logger: logger,
		}
	}); err != nil {
		b.quit(err)
	}

	return <-ch
}

func (b *Builder) provideConfig() *Builder {
	return b.provide(func(opts *options.Options) (*config.Config, error) {
		return config.Parse(opts.Config)
	}).
        provide(func(cfg *config.Config) *config.Database { return &cfg.Database }).
	    provide(func(cfg *config.Config) *config.Logging { return &cfg.Logging }).
	    provide(func(cfg *config.Config) *config.Storage { return &cfg.Storage }).
	    provide(func(cfg *config.Config) *config.TLS { return &cfg.TLS }).
	    provide(func(cfg *config.Config) *config.Token { return &cfg.Token })
}

func (b *Builder) provideLogger() *Builder {
	return b.provide(func(cfg *config.Logging) (*slog.Logger, error) {
		f, err := os.OpenFile(cfg.File, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
		if err != nil {
			return nil, err
		}

		opts := slog.HandlerOptions{
			Level: slog.Level(cfg.Level),
		}

		h := logging.NewHandler(
			slog.NewJSONHandler(f, &opts),
			slog.NewTextHandler(os.Stdout, &opts))

		return slog.New(h), nil
	})
}

func (b *Builder) provideRenderer() *Builder {
	return b.provide(renderer.New)
}

func (b *Builder) provideToken() *Builder {
	return b.
        provide(token.NewHandler).
		provide(func(h *token.Handler) token.Signer { return h }).
		provide(func(h *token.Handler) token.Verifier { return h })
}

func (b *Builder) provideDatabase() *Builder {
	return b.provide(func(cfg *config.Database) (database.Querier, error) {
        dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
            cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.DBName, cfg.SSLMode)

		db, err := sql.Open("postgres", dsn)
		if err != nil {
			return nil, err
		}

		return database.New(db), nil
	})
}

func (b *Builder) provideStorage() *Builder {
	return b.provide(storage.NewLocalProvider)
}

func (b *Builder) provideHttp() *Builder {
	return b.
        provide(router.New).
		provide(func(e *echo.Echo, opts *options.Options) *http.Server {
			return &http.Server{
				Addr:    fmt.Sprintf(":%d", opts.HttpPort),
				Handler: e,
			}
		})
}

func (b *Builder) provideGrpc() *Builder {
	return b.provide(func() *grpc.Server {
		return grpc.NewServer()
	})
}

func (b *Builder) provide(constructor any) *Builder {
	if err := b.Provide(constructor); err != nil {
		b.quit(err)
	}

	return b
}

func (b *Builder) quit(err error) {
	dig.Visualize(b.Container, os.Stdout,
		dig.VisualizeError(err))

	os.Exit(1)
}

func New(opts *options.Options) *Builder {
	c := dig.New()
	c.Provide(func() *options.Options { return opts })

	return &Builder{c}
}

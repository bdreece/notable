package server

import (
	"context"
	"log/slog"
)

type App struct {
    http servlet
    grpc servlet
    logger *slog.Logger
}

func (a *App) Run(ctx context.Context) error {
    ctx, cancel := context.WithCancel(ctx)
    defer cancel()

    errch := make(chan error, 2)
    defer close(errch)

    a.logger.Info("launching http server", slog.Int("port", a.http.port))
    go a.http.Start(ctx, errch)
    a.logger.Info("launching grpc server", slog.Int("port", a.grpc.port))
    go a.grpc.Start(ctx, errch)

    select {
    case err := <-errch:
        a.logger.Error("unexpected error occurred",
            slog.String("error", err.Error()))

        return err
    case <-ctx.Done():
        return nil
    }
}

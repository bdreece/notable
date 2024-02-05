package logging

import (
	"context"
	"log/slog"

	"golang.org/x/sync/errgroup"
)

type handler []slog.Handler

// Enabled implements slog.Handler.
func (handler handler) Enabled(c context.Context, l slog.Level) bool {
    for _, h := range handler {
        if h.Enabled(c, l) {
            return true
        }
    }

    return false
}

// Handle implements slog.Handler.
func (handler handler) Handle(c context.Context, r slog.Record) error {
    g, ctx := errgroup.WithContext(c)

    for _, h := range handler {
        h := h
        g.Go(func() error {
            return h.Handle(ctx, r)
        })
    }

    return g.Wait()
}

// WithAttrs implements slog.Handler.
func (oldHandler handler) WithAttrs(attrs []slog.Attr) slog.Handler {
    var newHandler handler

    for _, h := range oldHandler {
        newHandler = append(newHandler, h.WithAttrs(attrs))
    }

    return newHandler
}

// WithGroup implements slog.Handler.
func (oldHandler handler) WithGroup(name string) slog.Handler {
    var newHandler handler

    for _, h := range oldHandler {
        newHandler = append(newHandler, h.WithGroup(name))
    }

    return newHandler
}

func NewHandler(h ...slog.Handler) slog.Handler {
	return handler(h)
}

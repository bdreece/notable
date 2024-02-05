package main

import (
	"context"
	"os"
	"os/signal"

	"github.com/bdreece/notable/pkg/server"
	"github.com/bdreece/notable/pkg/server/options"
)

func main() {
    opts := options.Parse()
    app := server.New(opts).Build()
    
    ctx, cancel := context.WithCancel(context.Background())
    defer cancel()

    go func() {
        if err := app.Run(ctx); err != nil {
            panic(err)
        }
    }()

    quit := make(chan os.Signal, 1)
    signal.Notify(quit, os.Interrupt)
    <-quit
    cancel()
}

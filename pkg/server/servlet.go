package server

import (
	"context"
	"fmt"
	"net"
)

type servlet struct {
    server interface {
        Serve(net.Listener) error
    }
    port int
}

func (s *servlet) Start(ctx context.Context, errch chan<- error) {
    lst, err := net.Listen("tcp", fmt.Sprintf(":%d", s.port))
    if err != nil {
        errch<- err
        return
    }

    if err := s.server.Serve(lst); err != nil && err != net.ErrClosed {
        errch<- err
        return
    }
}


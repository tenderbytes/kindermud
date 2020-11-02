package mud

import (
	"context"
	"sync"

	"github.com/reiver/go-telnet"
)

type Connection struct {
	Context context.Context
}

type ServerOptions struct {
	Controllers       []*Controller
	DefaultController *Controller
	Addr              string
}

type Server struct {
	DefaultController *Controller
	Controllers       []*Controller
	Connections       []*Connection
	connMutex         *sync.Mutex
}

func NewServer(opts *ServerOptions) *Server {
	srv := &Server{
		DefaultController: opts.DefaultController,
		Controllers:       opts.Controllers[:],
		Connections:       make([]*Connection, 0),
		server: &telnet.Server{
			Addr:    opts.Addr,
			Handler: srv,
		},
	}

	return srv
}

func (srv *Server) ListenAndServe() error {
	return srv.server.ListenAndServe()
}

func (srv *Server) ServeTELNET(ctx Context, w Writer, r Reader) {
	conn := &Connection{
		Context: ctx,
	}
}

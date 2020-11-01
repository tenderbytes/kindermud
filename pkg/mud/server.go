package mud

import (
	"sync"

	"github.com/reiver/go-telnet"
)

type Connection struct {

}

type ServerOptions struct {
	Controllers       []*Controller
	DefaultController *Controller
}

type Server struct {
	DefaultController *Controller
	Controllers       []*Controller
	Connections       []*Connection
	connMutex         *sync.Mutex
}

func NewServer(opts *ServerOptions) *Server {
	return &Server{
		DefaultController: opts.DefaultController,
		Controllers:       opts.Controllers[:],
		Connections:       make([]*Connection, 0),
	}
}

func (srv *Server) ListenAndServe() error {
	handler := telnet.EchoHandler
return telnet.ListenAndServe(":5555", handler)
}

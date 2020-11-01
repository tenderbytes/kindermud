package mud

type ServerOptions struct {
	Controllers []*Controller
}

type Server struct {
}

func NewServer(opts *ServerOptions) *Server {

}

func (srv *Server) Listen() error {

}

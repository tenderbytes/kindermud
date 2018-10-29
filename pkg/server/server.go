package server

import (
	"context"
	"net"
	"net/http"

	gobagcontext "github.com/danielkrainas/gobag/context"
	"github.com/urfave/negroni"

	"github.com/tenderbytes/kindermud/pkg/api"
)

type Server interface {
	ListenAndServe() error
}

type CORSConfig struct {
	Origins []string
	Methods []string
	Headers []string
}

type Config struct {
	Debug bool
	Addr  string
	Host  string
	CORS  *CORSConfig
}

func New(ctx context.Context, apiMux api.Mux, config *Config) (srv Server, err error) {
	log := gobagcontext.GetLogger(ctx)
	n := negroni.New()

	n.Use(corsHandler(config.Debug, config.CORS))
	n.Use(contextHandler(ctx))
	n.UseFunc(loggingHandler)
	n.Use(&negroni.Recovery{
		Logger:     negroni.ALogger(log),
		PrintStack: true,
		StackAll:   true,
	})

	n.Use(aliveHandler("/"))
	n.UseFunc(trackErrorsHandler)
	n.UseHandler(apiMux)

	srv = &server{
		Context: ctx,
		config:  config,
		server: &http.Server{
			Addr:    config.Addr,
			Handler: n,
		},
	}

	return srv, nil
}

type server struct {
	context.Context
	config *Config
	server *http.Server
}

func (srv *server) ListenAndServe() error {
	config := srv.config
	ln, err := net.Listen("tcp", config.Addr)
	if err != nil {
		return err
	}

	gobagcontext.GetLogger(srv).Infof("listening on %v", ln.Addr())
	return srv.server.Serve(ln)
}

package manager

import (
	"context"
	"fmt"
	"os"

	"github.com/tenderbytes/kindermud/pkg/server"
)

type Manager interface {
	Start(ctx context.Context) error
}

func New(config *Config) (Manager, error) {
	return &manager{
		config:  config,
		factory: nil,
	}, nil
}

type manager struct {
	factory *Factory
	config  *Config
}

func (m *manager) init(ctx context.Context) (err error) {
	if m.factory != nil {
		return nil
	}

	fmt.Fprintf(os.Stderr, "====\n%+v\n====\n", m.config)
	m.factory = &Factory{
		SourceContext: ctx,
		Config:        m.config,
	}

	return m.factory.InitAll()
}

func (m *manager) Start(ctx context.Context) error {
	if err := m.init(ctx); err != nil {
		return err
	}

	factory := m.factory
	srv, err := server.New(factory.LoggingContext, factory.API, &server.Config{
		Debug: m.config.HTTP.Debug,
		Addr:  m.config.HTTP.Addr,
		Host:  m.config.HTTP.Host,
		CORS: &server.CORSConfig{
			Origins: m.config.HTTP.CORS.Origins,
			Methods: m.config.HTTP.CORS.Methods,
			Headers: m.config.HTTP.CORS.Headers,
		},
	})

	if err != nil {
		return err
	}

	return srv.ListenAndServe()
}

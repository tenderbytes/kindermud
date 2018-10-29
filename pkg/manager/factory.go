package manager

import (
	"context"
	"fmt"
	"time"

	cfg "github.com/danielkrainas/gobag/configuration"
	gobagcontext "github.com/danielkrainas/gobag/context"
	log "github.com/sirupsen/logrus"

	"github.com/tenderbytes/kindermud/pkg/api"
	"github.com/tenderbytes/kindermud/pkg/root"
	"github.com/tenderbytes/kindermud/pkg/root/realroot"
)

type Factory struct {
	Config         *Config
	SourceContext  context.Context
	LoggingContext context.Context
	API            api.Mux
}

func (f *Factory) InitAll() error {
	if err := f.InitLogging(); err != nil {
		return fmt.Errorf("error initializing logging: %v", err)
	}

	/*if err := f.InitStorage(); err != nil {
		return fmt.Errorf("error initializing storage: %v", err)
	}*/

	if err := f.InitRoot(); err != nil {
		return fmt.Errorf("error initializing root services: %v", err)
	}

	if err := f.InitAPI(); err != nil {
		return fmt.Errorf("error initializing api: %v", err)
	}

	return nil
}

func (f *Factory) InitLogging() error {
	if f.LoggingContext != nil {
		return nil
	}

	config := f.Config
	ctx := f.SourceContext
	log.SetLevel(logLevel(config.Log.Level))
	formatter := config.Log.Formatter
	if formatter == "" {
		formatter = "text"
	}

	switch formatter {
	case "json":
		log.SetFormatter(&log.JSONFormatter{
			TimestampFormat: time.RFC3339Nano,
		})

	case "text":
		log.SetFormatter(&log.TextFormatter{
			TimestampFormat: time.RFC3339Nano,
		})

	default:
		if config.Log.Formatter != "" {
			return fmt.Errorf("unsupported formatter: %q", config.Log.Formatter)
		}
	}

	if len(config.Log.Fields) > 0 {
		var fields []interface{}
		for k := range config.Log.Fields {
			fields = append(fields, k)
		}

		ctx = gobagcontext.WithValues(ctx, config.Log.Fields)
		ctx = gobagcontext.WithLogger(ctx, gobagcontext.GetLogger(ctx, fields...))
	}

	f.LoggingContext = gobagcontext.WithLogger(ctx, gobagcontext.GetLogger(ctx))
	log.Infof("using %q logging formatter", config.Log.Formatter)
	return nil
}

func (f *Factory) InitRoot() (err error) {
	if f.Root != nil {
		return nil
	}

	var root root.Service
	root, err = realroot.New(f.Storage, f.Labs)
	if err == nil && root != nil {
		gobagcontext.GetLogger(f.LoggingContext).Infof("using \"%s\" root services", root)
	}

	return err
}

func (f *Factory) InitAPI() (err error) {
	if f.API != nil {
		return nil
	}

	f.API, err = api.NewMux()
	return err
}

/*func (f *Factory) InitStorage() (err error) {
	typeName := f.Config.Storage.Type()
	params := f.Config.Storage.Parameters()
	log := gobagcontext.GetLogger(f.LoggingContext)
	switch typeName {
	case "mongodb":
		f.Storage, err = mongodbstorage.NewStorageConnector(params)

	default:
		err = fmt.Errorf("no connector for type %q", typeName)
	}

	if f.Storage != nil && err == nil {
		log.Infof("using %q storage connector", typeName)
	}

	return err
}*/

func logLevel(level cfg.LogLevel) log.Level {
	l, err := log.ParseLevel(string(level))
	if err != nil {
		l = log.InfoLevel
		log.Warnf("error parsing level %q: %v, using %q", level, err, l)
	}

	return l
}

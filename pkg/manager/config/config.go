package managerconfig

import (
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"reflect"
	"strings"

	cfg "github.com/danielkrainas/gobag/configuration"

	"github.com/tenderbytes/kindermud/pkg/manager"
)

// EnvPrefix is the prefix used when loading configuration from environment.
const EnvPrefix = "kinder"

// DefaultConfig is the application-default configuration.
var DefaultConfig = manager.Config{
	Printing: manager.PrintLabsConfig{
		EnableHolds: false,
	},
	Log: manager.LogConfig{
		Level:     "debug",
		Formatter: "text",
		Fields:    make(map[string]interface{}),
	},
	HTTP: manager.HTTPConfig{
		Debug: false,
		Addr:  ":4188",
		Host:  "localhost",
		CORS: manager.CORSConfig{
			Origins: []string{"*"},
			Methods: []string{"POST", "GET", "PUT", "PATCH", "HEAD", "OPTIONS"},
			Headers: []string{"*"},
		},
	},
}

// New creates a new config based on DefaultConfig
func New() *manager.Config {
	config := DefaultConfig
	return &config
}

// Resolve determines the application's config location and loads it.
func Resolve(configPath string) (*manager.Config, error) {
	if configPath == "" {
		configPath = os.Getenv(strings.ToUpper(EnvPrefix) + "_CONFIG_PATH")
	}

	if configPath == "" {
		return nil, errors.New("configuration path not specified")
	}

	fp, err := os.Open(configPath)
	if err != nil {
		return nil, fmt.Errorf("configuration: %v", err)
	}

	defer fp.Close()
	config, err := Parse(fp)
	if err != nil {
		return nil, fmt.Errorf("error parsing %s: %v", configPath, err)
	}

	return config, nil
}

// Validate determines if the configuration is prepared correctly and valid to use.
func Validate(config *manager.Config) (*manager.Config, error) {
	return config, nil
}

// v1_0Config is used when versioning config files
type v1_0Config manager.Config

// Parse loads and parses the configuration from a reader.
func Parse(rd io.Reader) (*manager.Config, error) {
	in, err := ioutil.ReadAll(rd)
	if err != nil {
		return nil, err
	}

	p := cfg.NewParser(strings.ToLower(EnvPrefix), []cfg.VersionedParseInfo{
		{
			Version: cfg.MajorMinorVersion(1, 0),
			ParseAs: reflect.TypeOf(v1_0Config{}),
			ConversionFunc: func(c interface{}) (interface{}, error) {
				if v1_0, ok := c.(*v1_0Config); ok {
					return Validate((*manager.Config)(v1_0))
				}

				return nil, fmt.Errorf("Expected *v1_0Config, received %#v", c)
			},
		},
	})

	config := new(manager.Config)
	if err = p.Parse(in, config); err != nil {
		return nil, err
	}

	return config, nil
}

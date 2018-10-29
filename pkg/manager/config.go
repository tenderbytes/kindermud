package manager

import (
	cfg "github.com/danielkrainas/gobag/configuration"
)

type LogConfig struct {
	Level     cfg.LogLevel           `yaml:"level"`
	Formatter string                 `yaml:"formatter"`
	Fields    map[string]interface{} `yaml:"fields"`
}

type CORSConfig struct {
	Origins []string `yaml:"origins"`
	Methods []string `yaml:"methods"`
	Headers []string `yaml:"headers"`
}

type HTTPConfig struct {
	Debug bool       `yaml:"debug"`
	Addr  string     `yaml:"addr"`
	Host  string     `yaml:"host"`
	CORS  CORSConfig `yaml:"cors"`
}

type Config struct {
	Storage cfg.Driver `yaml:"storage"`
	Log     LogConfig  `yaml:"log"`
	HTTP    HTTPConfig `yaml:"http"`
}

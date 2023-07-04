package ethmetrics

import (
	"github.com/keRin7/ethMetrics/internal/webServer"
)

type Config struct {
	Url          string `env:"URL"`
	LogLevel     string `env:"LOG_LEVEL" envDefault:"warn"`
	QueryTimeout int    `env:"QUERY_TIMEOUT"  envDefault:"15"`
	WebService   *webServer.Config
}

func NewConfig() *Config {
	return &Config{
		WebService: webServer.NewConfig(),
	}
}

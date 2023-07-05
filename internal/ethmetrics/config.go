package ethmetrics

import (
	"github.com/keRin7/ethMetrics/internal/webServer"
)

type Config struct {
	Hosts        []string `env:"HOSTS" envSeparator:","`
	LogLevel     string   `env:"LOG_LEVEL" envDefault:"info"`
	QueryTimeout int      `env:"QUERY_TIMEOUT"  envDefault:"15"`
	WebService   *webServer.Config
}

func NewConfig() *Config {
	return &Config{
		WebService: webServer.NewConfig(),
	}
}

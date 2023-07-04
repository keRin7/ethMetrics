package ethmetrics

type Config struct {
	Url      string `env:"URL"`
	LogLevel string `default:"warn" env:"LOG_LEVEL"`
}

func NewConfig() *Config {
	return &Config{}
}

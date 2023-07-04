package ethmetrics

type Config struct {
	Url string `env:"URL"`
}

func NewConfig() *Config {
	return &Config{}
}

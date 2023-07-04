package ethmetrics

type EthMetrics struct {
	config *Config
}

func CreateEthMetrics(config *Config) *EthMetrics {
	return &EthMetrics{
		config: config,
	}
}

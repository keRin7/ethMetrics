package ethmetrics

import (
	"github.com/keRin7/ethMetrics/internal/vPrometheus"
	"github.com/keRin7/ethMetrics/internal/webServer"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/sirupsen/logrus"
)

type EthMetrics struct {
	Config    *Config
	Prom      *vPrometheus.AppPrometheus
	WebServer *webServer.WebServer
}

func CreateEthMetrics(config *Config) *EthMetrics {

	switch config.LogLevel {
	case "debug":
		logrus.SetLevel(logrus.DebugLevel)
	case "info":
		logrus.SetLevel(logrus.InfoLevel)
	case "warn":
		logrus.SetLevel(logrus.WarnLevel)
	case "error":
		logrus.SetLevel(logrus.ErrorLevel)
	default:
		logrus.SetLevel(logrus.InfoLevel)
		logrus.Warnf("Home: invalid log level supplied: '%s'", config.LogLevel)
	}

	return &EthMetrics{
		Config:    config,
		Prom:      vPrometheus.CreatePrometheus(),
		WebServer: webServer.CreateWebServer(config.WebService),
	}
}

func (c *EthMetrics) Start() {

	logrus.Printf("Web started on port: %s", c.Config.WebService.Port)
	c.Prom.InitPrometheus()
	c.WebServer.AddHeandler("/metrics", promhttp.Handler())
	c.WebServer.Start()
}

package ethmetrics

import (
	"context"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/keRin7/ethMetrics/internal/vPrometheus"
	"github.com/keRin7/ethMetrics/internal/webServer"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/sirupsen/logrus"

	"time"
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

func (c *EthMetrics) getMetrics(ctx context.Context, host string) {
	rpc, err := rpc.Dial(host)
	if err != nil {
		logrus.Warn(err)
	}
	var result hexutil.Uint64
	logrus.Info("Start watching host:", host)
	for {
		select {
		case <-ctx.Done():
			return
		default:
			time.Sleep(time.Duration(c.Config.QueryTimeout) * time.Second)
			if err := rpc.Call(&result, "net_peerCount"); err != nil {
				logrus.Warn("peerCount getting error")
			} else {
				c.Prom.WriteMetric(host, "net_peerCount", float64(result))
			}
			if err := rpc.Call(&result, "eth_blockNumber"); err != nil {
				logrus.Warn("blockNumber getting error")
			} else {
				c.Prom.WriteMetric(host, "eth_blockNumber", float64(result))
			}
		}
	}
}

func (c *EthMetrics) Start(ctx context.Context) {
	logrus.Printf("Web started on port: %s", c.Config.WebService.Port)
	c.Prom.InitPrometheus()
	c.WebServer.AddHeandler("/metrics", promhttp.Handler())
	for _, host := range c.Config.Hosts {
		go c.getMetrics(ctx, host)
	}
	c.WebServer.Start()
}

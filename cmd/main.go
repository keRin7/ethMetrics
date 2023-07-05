package main

import (
	"context"
	"github.com/caarlos0/env"
	"github.com/keRin7/ethMetrics/internal/ethmetrics"
	"github.com/sirupsen/logrus"
)

func main() {

	ctx, finish := context.WithCancel(context.Background())
	defer finish()
	config := ethmetrics.NewConfig()
	err := env.Parse(config)
	if err != nil {
		logrus.Fatal(err)
	}
	metrics := ethmetrics.CreateEthMetrics(config)
	metrics.Start(ctx)

}

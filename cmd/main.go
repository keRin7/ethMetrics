package main

import (
	"context"

	"github.com/caarlos0/env"
	ethMetrics "github.com/keRin7/ethMetrics/internal/ethmetrics"
	"github.com/sirupsen/logrus"
	//"github.com/ethereum/go-ethereum/ethclient"
)

func main() {

	//	client, err := ethclient.Dial("http://localhost:8545")
	//	if err != nil {
	//		log.Fatal(err)
	//	}
	//	header, err := client.HeaderByNumber(context.Background(), nil)
	//	if err != nil {
	//		log.Fatal(err)
	//	}
	ctx, finish := context.WithCancel(context.Background())
	defer finish()
	config := ethMetrics.NewConfig()
	err := env.Parse(config)
	if err != nil {
		logrus.Fatal(err)
	}
	metrics := ethMetrics.CreateEthMetrics(config)
	metrics.Start(ctx)

}

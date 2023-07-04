package main

import (
	"fmt"
	"log"

	"github.com/caarlos0/env"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/rpc"
	ethMetrics "github.com/keRin7/ethmetrics/internal/ethmetrics"
	"github.com/sirupsen/logrus"
	//"github.com/ethereum/go-ethereum/ethclient"
)

func main() {
	fmt.Println("Hello")

	//	client, err := ethclient.Dial("http://localhost:8545")
	//	if err != nil {
	//		log.Fatal(err)
	//	}
	//	header, err := client.HeaderByNumber(context.Background(), nil)
	//	if err != nil {
	//		log.Fatal(err)
	//	}
	config := ethMetrics.NewConfig()

	err := env.Parse(config)
	if err != nil {
		logrus.Fatal(err)
	}

	rpc, err := rpc.Dial("http://localhost:8545")
	if err != nil {
		log.Fatal(err)
	}

	var result hexutil.Uint64
	if err := rpc.Call(&result, "net_peerCount"); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(float64(result))
	}
	if err := rpc.Call(&result, "eth_blockNumber"); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(float64(result))
	}
}

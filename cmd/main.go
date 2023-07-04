package main

import (
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/rpc"
	ethMetrics "github.com/keRin7/ethMetrics/internal/ethMetrics"
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

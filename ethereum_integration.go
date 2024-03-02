package main

import (
    "context"
    "fmt"
    "log"

    "github.com/ethereum/go-ethereum/common"
    "github.com/ethereum/go-ethereum/ethclient"
)

func main() {
    client, err := ethclient.Dial("https://eth-mainnet.g.alchemy.com/v1/nsKEdR9mvSlYb22X-ih-nUlFONE2cdJk")
    if err != nil {
        log.Fatal(err)
    }

    // Get the latest known block
    block, err := client.BlockByNumber(context.Background(), nil)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println("Latest block:", block.Number().Uint64())

    // Example usage of a function from the common package
    address := common.HexToAddress("0x8d8f7a197bdbdf9cf784bb79eb4b5145f5d419cf")
    fmt.Println("Address:", address.Hex())
}

package main

import (
	"TheForestProject/blockchain"
	"fmt"
	"strconv"
)

func main() {
	chain := blockchain.InitBlockChain()

	chain.AddBlock("First after Gen")
	chain.AddBlock("Second after Gen")
	chain.AddBlock("Third after Gen")

	for _, block := range chain.Blocks {
		fmt.Printf("Data: %s\n", block.Data)
		fmt.Printf("Previous hash: %x\n", block.PrevHash)
		fmt.Printf("Hash: %x\n", block.Hash)

		pow := blockchain.NewProof(block)
		fmt.Printf("PoW: %s\n", strconv.FormatBool(pow.Validate()))
		fmt.Println()
	}
}

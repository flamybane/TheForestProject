package main

import (
	"TheForestProject/cli"
	"TheForestProject/wallet"
	"os"
)

func main() {
	defer os.Exit(0)
	cmd := cli.CommandLine{}
	cmd.Run()

	w := wallet.MakeWallet()
	w.Address()
}

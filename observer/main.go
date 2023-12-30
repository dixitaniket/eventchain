package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/dixitaniket/eventchain/observer/event"
	"golang.org/x/sync/errgroup"
)

const (
	contractAddress = "0x5FbDB2315678afecb367f032d93F642f64180aa3" // replace with your contract address
	infuraURL       = "ws://0.0.0.0:8545"                          // replace with your Infura Endpoint or own Ethereum node
)

func main() {
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGTSTP)
	ctx, cancel := context.WithCancel(context.Background())
	g, gctx := errgroup.WithContext(ctx)
	o := event.NewObserver(contractAddress, infuraURL)
	g.Go(func() error {
		return o.Start(gctx)
	})

	// Process events
	fmt.Println("Started listening for Launch events. Press CTRL + Z to exit.")
	for {
		select {
		case <-sigs:
			cancel()
			fmt.Println("Exiting...")
			return
		}
	}
}

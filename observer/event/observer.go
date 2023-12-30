package event

import (
	"context"
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

type Observer struct {
	relayerAddress  string
	contractAddress string
	rpcUrl          string
}

func NewObserver(relayerAddress, address, rpc string) *Observer {
	observer := &Observer{
		relayerAddress:  relayerAddress,
		contractAddress: address,
		rpcUrl:          rpc,
	}
	return observer
}

func (o *Observer) Start(ctx context.Context) error {
	client, err := ethclient.Dial(o.rpcUrl)
	if err != nil {
		return fmt.Errorf("failed to connect to the eth client: %v", err)
	}

	address := common.HexToAddress(o.contractAddress)
	token, err := NewTestEvent(address, client)
	sub := make(chan *TestEventLaunch)
	see, err := token.WatchLaunch(nil, sub, nil, nil)

	for {
		select {
		case _ = <-see.Err():
		case <-ctx.Done():
			fmt.Println("closing subscription")
			see.Unsubscribe()
			return nil
		case other := <-sub:
			fmt.Println(other.Number, other.Toadd)
		}
	}
}

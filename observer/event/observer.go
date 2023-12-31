package event

import (
	"context"
	"fmt"

	"github.com/dixitaniket/eventchain/observer/client"
	"github.com/dixitaniket/eventchain/x/oracle/types"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/rs/zerolog"
)

type Observer struct {
	logger          zerolog.Logger
	relayerAddress  string
	contractAddress string
	rpcUrl          string
	timeoutHeight   int64
	oc              client.OracleClient
}

func NewObserver(logger zerolog.Logger, relayerAddress, address, rpc string, height int64, oc client.OracleClient) *Observer {
	observer := &Observer{
		logger:          logger.With().Str("module", "observer").Logger(),
		relayerAddress:  relayerAddress,
		contractAddress: address,
		rpcUrl:          rpc,
		timeoutHeight:   height,
		oc:              oc,
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
			o.logger.Info().Msg("closing subscription")
			see.Unsubscribe()
			return nil
		case other := <-sub:
			err := o.processMsg(other.Number, other.Toadd)
			if err != nil {
				o.logger.Err(err).Send()
			}
		}
	}
}

func (o *Observer) processMsg(num uint8, add uint8) error {
	msg := types.MsgPostResult{
		Creator: o.relayerAddress,
		Result: types.Result{
			Num:   int64(num),
			Toadd: int64(add),
		},
	}
	height, err := o.oc.ChainHeight.GetChainHeight()
	if err != nil {
		return err
	}

	err = o.oc.BroadcastTx(height, o.timeoutHeight, &msg)
	if err != nil {
		return fmt.Errorf("error in broadcasting msg %s", err)
	}
	return nil
}

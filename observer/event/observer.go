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
	if err != nil {
		return err
	}
	eventSink := make(chan *TestEventLaunch)
	eventSub, err := token.WatchLaunch(nil, eventSink, nil, nil)
	if err != nil {
		return err
	}
	defer eventSub.Unsubscribe()
	for {
		select {
		case err = <-eventSub.Err():
			o.logger.Err(err).Send()
		case <-ctx.Done():
			o.logger.Info().Msg("closing subscription")
			return nil
		case event := <-eventSink:
			o.logger.Info().Uint8("number", event.Number).
				Uint8("to add", event.Toadd).
				Msg("new event received")
			err := o.processMsg(event.Number, event.Toadd, event.Raw.BlockNumber)
			if err != nil {
				o.logger.Err(err).Send()
			}
		}
	}
}

func (o *Observer) processMsg(num uint8, add uint8, blockNum uint64) error {
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
	msg.ChainHeight = int64(blockNum)
	msg.BlockHeight = height

	o.logger.Info().Int64("chain block height", height).
		Uint64("evm chain height", blockNum).
		Msg("broadcasting msg")

	err = o.oc.BroadcastTx(height, o.timeoutHeight, &msg)
	if err != nil {
		return fmt.Errorf("error in broadcasting msg %s", err)
	}
	return nil
}

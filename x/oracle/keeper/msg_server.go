package keeper

import (
	"context"

	"cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/dixitaniket/eventchain/x/oracle/types"
)

type msgServer struct {
	Keeper
}

// NewMsgServerImpl returns an implementation of the MsgServer interface
// for the provided Keeper.
func NewMsgServerImpl(keeper Keeper) types.MsgServer {
	return &msgServer{Keeper: keeper}
}

var _ types.MsgServer = msgServer{}

func (k msgServer) PostResult(goCtx context.Context, msg *types.MsgPostResult) (*types.MsgPostResultResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	// check whitelist
	operatorAddr, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return nil, err
	}
	if !k.CheckWhitelist(ctx, operatorAddr) {
		return nil, errors.Wrapf(types.ErrNotWhitelisted, "operator not whitelisted %s", operatorAddr.String())
	}

	err = k.SetResult(ctx, operatorAddr, msg.Result)
	if err != nil {
		return nil, err
	}
	return &types.MsgPostResultResponse{}, nil
}

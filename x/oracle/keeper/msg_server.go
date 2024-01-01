package keeper

import (
	"context"

	"cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
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

	if msg.BlockHeight > ctx.BlockHeight() {
		return nil, errors.Wrapf(types.ErrInvalidBlockHeight, "block height %d greater than current block height %d", msg.BlockHeight, ctx.BlockHeight())
	}

	err = k.SetResult(ctx, operatorAddr, msg.Result)
	if err != nil {
		return nil, err
	}
	return &types.MsgPostResultResponse{}, nil
}

func (k msgServer) ProposeWhitelist(goCtx context.Context, msg *types.MsgProposeWhitelist) (*types.MsgProposeWhitelistResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	authority := msg.Authority
	if authority != authtypes.NewModuleAddress(govtypes.ModuleName).String() {
		return nil, errors.Wrapf(types.ErrNotAuthorized, "authority is not gov")
	}
	whitelist, err := sdk.AccAddressFromBech32(msg.WhitelistOperator)
	if err != nil {
		return nil, err
	}
	k.SetWhitelist(ctx, whitelist)
	return &types.MsgProposeWhitelistResponse{}, nil
}

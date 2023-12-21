package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/dixitaniket/eventchain/x/oracle/types"
)

func (k msgServer) PostResult(goCtx context.Context, msg *types.MsgPostResult) (*types.MsgPostResultResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgPostResultResponse{}, nil
}

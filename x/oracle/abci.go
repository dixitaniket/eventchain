package oracle

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/dixitaniket/eventchain/x/oracle/keeper"
)

func EndBlocker(ctx sdk.Context, k keeper.Keeper) {
	err := k.SetFinalResult(ctx)
	if err != nil {
		ctx.Logger().Error("Error in final result end blocker", "err", err)
	}
}

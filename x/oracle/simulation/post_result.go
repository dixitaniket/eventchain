package simulation

import (
	"math/rand"

	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/dixitaniket/eventchain/x/oracle/keeper"
	"github.com/dixitaniket/eventchain/x/oracle/types"
)

func SimulateMsgPostResult(
	ak types.AccountKeeper,
	bk types.BankKeeper,
	k keeper.Keeper,
) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		simAccount, _ := simtypes.RandomAcc(r, accs)
		msg := &types.MsgPostResult{
			Creator: simAccount.Address.String(),
		}

		// TODO: Handling the PostResult simulation

		return simtypes.NoOpMsg(types.ModuleName, msg.Type(), "PostResult simulation not implemented"), nil, nil
	}
}

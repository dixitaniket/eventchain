package oracle

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/dixitaniket/eventchain/x/oracle/keeper"
	"github.com/dixitaniket/eventchain/x/oracle/types"
)

// InitGenesis initializes the module's state from a provided genesis state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	// this line is used by starport scaffolding # genesis/module/init
	k.SetParams(ctx, genState.Params)
	// store whitelisted address

	for _, addr := range genState.Whitelist {
		operatorAddr, err := sdk.AccAddressFromBech32(addr)
		if err != nil {
			panic(err)
		}
		k.SetWhitelist(ctx, operatorAddr)
	}
}

// ExportGenesis returns the module's exported genesis
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()
	genesis.Params = k.GetParams(ctx)

	// this line is used by starport scaffolding # genesis/module/export

	return genesis
}

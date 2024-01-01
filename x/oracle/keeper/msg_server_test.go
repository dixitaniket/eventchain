package keeper_test

import (
	"context"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
	keepertest "github.com/dixitaniket/eventchain/testutil/keeper"
	"github.com/dixitaniket/eventchain/testutil/sample"
	"github.com/dixitaniket/eventchain/x/oracle/keeper"
	"github.com/dixitaniket/eventchain/x/oracle/types"
	"github.com/stretchr/testify/require"
)

func setupMsgServer(t testing.TB) (keeper.Keeper, types.MsgServer, context.Context) {
	k, ctx := keepertest.OracleKeeper(t)
	ctx = ctx.WithBlockHeight(10)
	return *k, keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}

func TestPostResult(t *testing.T) {
	k, ms, ctx := setupMsgServer(t)
	creator := sample.AccAddress()

	msg := &types.MsgPostResult{
		Creator:     creator.String(),
		BlockHeight: 9,
		Result: types.Result{
			Num:   10,
			Toadd: 20,
		},
		ChainHeight: 10,
	}
	_, err := ms.PostResult(ctx, msg)
	require.Error(t, err)

	// Whitelist operator
	k.SetWhitelist(sdk.UnwrapSDKContext(ctx), creator)
	_, err = ms.PostResult(ctx, msg)
	require.NoError(t, err)

	// Test for whitelisting error
	msgInvalid := &types.MsgPostResult{
		Creator: "INVALID", // Invalid creator
	}
	_, err = ms.PostResult(ctx, msgInvalid)
	require.Error(t, err)

	// check in keeper
	result, err := k.GetResult(sdk.UnwrapSDKContext(ctx), creator)
	require.NoError(t, err)
	require.Equal(t, msg.Result, result)
}

func TestProposeWhitelist(t *testing.T) {
	k, ms, ctx := setupMsgServer(t)
	toWhitelist := sample.AccAddress()
	// Test for successful ProposeWhitelist
	msg := &types.MsgProposeWhitelist{
		Authority:         sample.AccAddress().String(), //invalid authority
		WhitelistOperator: toWhitelist.String(),
	}
	_, err := ms.ProposeWhitelist(ctx, msg)
	require.ErrorIs(t, err, types.ErrNotAuthorized)

	// with valid authority
	msg.Authority = authtypes.NewModuleAddress(govtypes.ModuleName).String()
	_, err = ms.ProposeWhitelist(ctx, msg)
	require.NoError(t, err)

	require.True(t, k.CheckWhitelist(sdk.UnwrapSDKContext(ctx), toWhitelist))
}

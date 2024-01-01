package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	testkeeper "github.com/dixitaniket/eventchain/testutil/keeper"
	"github.com/dixitaniket/eventchain/testutil/sample"
	"github.com/dixitaniket/eventchain/x/oracle/types"
	"github.com/stretchr/testify/require"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.OracleKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}

func TestSetAndCheckWhitelist(t *testing.T) {
	k, ctx := testkeeper.OracleKeeper(t)
	acc := sample.AccAddress()
	require.False(t, k.CheckWhitelist(ctx, acc))

	k.SetWhitelist(ctx, acc)
	require.True(t, k.CheckWhitelist(ctx, acc))
}

func TestSetAndGetResult(t *testing.T) {
	k, ctx := testkeeper.OracleKeeper(t)
	acc := sample.AccAddress()
	result := types.Result{
		Num:   10,
		Toadd: 5,
	}

	// Set result for account
	err := k.SetResult(ctx, acc, result)
	require.NoError(t, err)

	// Get result for account
	fetchedResult, err := k.GetResult(ctx, acc)
	require.NoError(t, err)
	require.Equal(t, result, fetchedResult)
}

func TestGetWhitelist(t *testing.T) {
	k, ctx := testkeeper.OracleKeeper(t)
	addresses := []sdk.AccAddress{
		sample.AccAddress(),
		sample.AccAddress(),
	}

	// Set addresses to the whitelist
	for _, acc := range addresses {
		k.SetWhitelist(ctx, acc)
	}

	// Fetch the whitelist and ensure it contains the addresses
	whitelist := k.GetWhitelist(ctx)
	require.Len(t, whitelist, len(addresses))
	for _, addr := range addresses {
		require.Contains(t, whitelist, addr.String())
	}
}

func TestSetFinalResult(t *testing.T) {
	k, ctx := testkeeper.OracleKeeper(t)
	params := types.DefaultParams()
	k.SetParams(ctx, params)

	// Verify no final result is set due to not enough votes
	_, err := k.GetFinalResult(ctx)
	require.Error(t, err)

	// cast max votes
	for i := 0; i < int(params.MinVotes); i++ {
		result := types.Result{
			Num:   10,
			Toadd: 20,
		}
		err := k.SetResult(ctx, sample.AccAddress(), result)
		require.NoError(t, err)
	}

	err = k.SetFinalResult(ctx)
	require.NoError(t, err)

	finalResult, err := k.GetFinalResult(ctx)
	require.NoError(t, err)
	require.Equal(t, finalResult.Result, int64(30))
	require.Equal(t, finalResult.Height, ctx.BlockHeight())

	// advance block height and increase min votes
	ctx = ctx.WithBlockHeight(ctx.BlockHeight() + 10)
	params.MinVotes = 5
	k.SetParams(ctx, params)

	for i := 0; i < 4; i++ {
		result := types.Result{
			Num:   20,
			Toadd: 20,
		}
		err = k.SetResult(ctx, sample.AccAddress(), result)
		require.NoError(t, err)
	}

	require.ErrorIs(t, k.SetFinalResult(ctx), types.ErrNoConsensus)

	// should be the old result
	finalResult, err = k.GetFinalResult(ctx)
	require.NoError(t, err)
	require.Equal(t, finalResult.Result, int64(30))

	for i := 0; i < 3; i++ {
		result := types.Result{
			Num:   20,
			Toadd: 20,
		}
		err = k.SetResult(ctx, sample.AccAddress(), result)
		require.NoError(t, err)
	}
	for i := 0; i < 3; i++ {
		result := types.Result{
			Num:   10,
			Toadd: 0,
		}
		err = k.SetResult(ctx, sample.AccAddress(), result)
		require.NoError(t, err)
	}

	// equally divided votes thus no consensus
	require.Error(t, k.SetFinalResult(ctx), types.ErrNoConsensus)

	// should be the old result
	finalResult, err = k.GetFinalResult(ctx)
	require.NoError(t, err)
	require.Equal(t, finalResult.Result, int64(30))

	for i := 0; i < 4; i++ {
		result := types.Result{
			Num:   20,
			Toadd: 20,
		}
		err = k.SetResult(ctx, sample.AccAddress(), result)
		require.NoError(t, err)
	}
	for i := 0; i < 3; i++ {
		result := types.Result{
			Num:   10,
			Toadd: 0,
		}
		err = k.SetResult(ctx, sample.AccAddress(), result)
		require.NoError(t, err)
	}

	require.NoError(t, k.SetFinalResult(ctx))

	// should be the new result
	finalResult, err = k.GetFinalResult(ctx)
	require.NoError(t, err)
	require.Equal(t, finalResult.Result, int64(40))
}

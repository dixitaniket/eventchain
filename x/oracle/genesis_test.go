package oracle_test

import (
	"testing"

	keepertest "github.com/dixitaniket/eventchain/testutil/keeper"
	"github.com/dixitaniket/eventchain/testutil/sample"
	"github.com/dixitaniket/eventchain/x/oracle"
	"github.com/dixitaniket/eventchain/x/oracle/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params:    types.DefaultParams(),
		Whitelist: []string{sample.AccAddress().String()},
	}

	k, ctx := keepertest.OracleKeeper(t)
	oracle.InitGenesis(ctx, *k, genesisState)

	got := oracle.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	require.EqualValues(t, genesisState, *got)
}

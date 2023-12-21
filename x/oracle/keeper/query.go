package keeper

import (
	"github.com/dixitaniket/eventchain/x/oracle/types"
)

var _ types.QueryServer = Keeper{}

package keeper

import (
	"fmt"

	"github.com/cometbft/cometbft/libs/log"
	"github.com/cosmos/cosmos-sdk/codec"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
	"github.com/dixitaniket/eventchain/x/oracle/types"
)

type (
	Keeper struct {
		cdc        codec.BinaryCodec
		storeKey   storetypes.StoreKey
		memKey     storetypes.StoreKey
		paramstore paramtypes.Subspace
	}
)

func NewKeeper(
	cdc codec.BinaryCodec,
	storeKey,
	memKey storetypes.StoreKey,
	ps paramtypes.Subspace,

) *Keeper {
	// set KeyTable if it has not already been set
	if !ps.HasKeyTable() {
		ps = ps.WithKeyTable(types.ParamKeyTable())
	}

	return &Keeper{
		cdc:        cdc,
		storeKey:   storeKey,
		memKey:     memKey,
		paramstore: ps,
	}
}

func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", types.ModuleName))
}

func (k Keeper) setKey(ctx sdk.Context, key, val []byte) {
	kvstore := ctx.KVStore(k.storeKey)
	kvstore.Set(key, val)
}

func (k Keeper) getKey(ctx sdk.Context, key []byte) ([]byte, bool) {
	kvstore := ctx.KVStore(k.storeKey)
	if !kvstore.Has(key) {
		return nil, false
	}

	return kvstore.Get(key), true
}

func (k Keeper) SetWhitelist(ctx sdk.Context, acc sdk.AccAddress) {
	k.setKey(ctx, types.GetWhitelistKey(acc), []byte(""))
}

func (k Keeper) SetResult(ctx sdk.Context, acc sdk.AccAddress, result types.Result) error {
	res, err := k.cdc.Marshal(&result)
	if err != nil {
		return err
	}
	k.setKey(ctx, types.GetResultInt(acc), res)
	return nil
}

func (k Keeper) CheckWhitelist(ctx sdk.Context, acc sdk.AccAddress) bool {
	_, found := k.getKey(ctx, types.GetWhitelistKey(acc))
	return found
}

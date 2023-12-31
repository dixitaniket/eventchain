package keeper

import (
	"fmt"

	sdkerrors "cosmossdk.io/errors"
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

func (k Keeper) GetResultIterator(ctx sdk.Context) sdk.Iterator {
	return sdk.KVStorePrefixIterator(ctx.KVStore(k.storeKey), types.KeyPrefixResultInt)
}

func (k Keeper) SetFinalResult(ctx sdk.Context) error {
	iterator := k.GetResultIterator(ctx)
	defer iterator.Close()

	// could use map iterator as it does not affect order here
	totalCount := make(map[int64]int)
	for ; iterator.Valid(); iterator.Next() {
		var val types.Result
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		total := val.Toadd + val.Num
		if _, found := totalCount[total]; !found {
			totalCount[total] = 0
		}
		totalCount[total] += 1
	}
	maxcounts := 0
	finalVal := int64(0)
	for val, counts := range totalCount {
		if counts > maxcounts {
			finalVal = val
			maxcounts = counts
		}
	}

	res := types.FinalResult{
		Result: finalVal,
		Height: ctx.BlockHeight(),
	}
	resBuf, err := k.cdc.Marshal(&res)
	if err != nil {
		return err
	}
	k.setKey(ctx, types.KeyFinalResult, resBuf)
	return nil
}

func (k Keeper) GetFinalResult(ctx sdk.Context) (types.FinalResult, error) {
	buf, found := k.getKey(ctx, types.KeyFinalResult)
	if !found {
		return types.FinalResult{}, sdkerrors.Wrap(types.ErrNoFinalResult, "not found")
	}
	var result types.FinalResult
	err := k.cdc.UnmarshalInterface(buf, &result)
	if err != nil {
		return types.FinalResult{}, sdkerrors.Wrap(types.ErrNoFinalResult, "not found")
	}

	return result, nil
}

func (k Keeper) CheckWhitelist(ctx sdk.Context, acc sdk.AccAddress) bool {
	_, found := k.getKey(ctx, types.GetWhitelistKey(acc))
	return found
}

package keeper

import (
	"fmt"

	sdkerrors "cosmossdk.io/errors"
	"github.com/cometbft/cometbft/libs/log"
	"github.com/cosmos/cosmos-sdk/codec"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerror "github.com/cosmos/cosmos-sdk/types/errors"
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

func (k Keeper) delKeys(ctx sdk.Context, keys [][]byte) {
	kvstore := ctx.KVStore(k.storeKey)
	for _, key := range keys {
		kvstore.Delete(key)
	}
}

func (k Keeper) getKey(ctx sdk.Context, key []byte) ([]byte, bool) {
	kvstore := ctx.KVStore(k.storeKey)
	if !kvstore.Has(key) {
		return nil, false
	}

	return kvstore.Get(key), true
}

func (k Keeper) SetWhitelist(ctx sdk.Context, acc sdk.AccAddress) {
	k.setKey(ctx, types.GetWhitelistKey(acc), acc)
}

func (k Keeper) SetResult(ctx sdk.Context, acc sdk.AccAddress, result types.Result) error {
	res, err := k.cdc.Marshal(&result)
	if err != nil {
		return err
	}
	k.setKey(ctx, types.GetResultInt(acc), res)
	return nil
}

func (k Keeper) GetResult(ctx sdk.Context, acc sdk.AccAddress) (types.Result, error) {
	res, found := k.getKey(ctx, types.GetResultInt(acc))
	if !found {
		return types.Result{}, sdkerrors.Wrap(sdkerror.ErrNotFound, "result not found")
	}
	var result types.Result
	err := k.cdc.Unmarshal(res, &result)
	if err != nil {
		return types.Result{}, err
	}
	return result, nil
}

func (k Keeper) GetIterator(ctx sdk.Context, key []byte) sdk.Iterator {
	return sdk.KVStorePrefixIterator(ctx.KVStore(k.storeKey), key)
}

func (k Keeper) SetFinalResult(ctx sdk.Context) error {
	params := k.GetParams(ctx)
	minVotes := params.GetMinVotes()
	iterator := k.GetIterator(ctx, types.KeyPrefixResultInt)
	defer iterator.Close()

	// could use map iterator as it does not affect order here
	totalCount := make(map[int64]int)
	totalVoters := make([][]byte, 0)
	for ; iterator.Valid(); iterator.Next() {
		var val types.Result
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		total := val.Toadd + val.Num
		if _, found := totalCount[total]; !found {
			totalCount[total] = 0
		}
		totalCount[total] += 1
		totalVoters = append(totalVoters, iterator.Key())
	}
	if len(totalVoters) == 0 {
		// no votes present
		return nil
	}
	if len(totalVoters) < int(minVotes) {
		ctx.Logger().Info("total votes less than min required", "total votes", len(totalVoters), "block height", ctx.BlockHeight())
		k.delKeys(ctx, totalVoters)
		return nil
	}
	maxCounts := 0
	finalVal := int64(0)
	for val, counts := range totalCount {
		if counts > maxCounts {
			finalVal = val
			maxCounts = counts
		}
	}

	// check if it is in consensus
	requiredConsensus := (len(totalVoters) + 1) / 2
	if maxCounts < requiredConsensus {
		ctx.Logger().Info("total votes less than consensus", "total votes received", maxCounts, "total required", requiredConsensus)
		k.delKeys(ctx, totalVoters)
		return nil
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
	// delete all previous aggregated results
	k.delKeys(ctx, totalVoters)
	return nil
}

func (k Keeper) GetFinalResult(ctx sdk.Context) (types.FinalResult, error) {
	buf, found := k.getKey(ctx, types.KeyFinalResult)
	if !found {
		return types.FinalResult{}, sdkerrors.Wrap(types.ErrNoFinalResult, "not found")
	}

	var result types.FinalResult
	err := k.cdc.Unmarshal(buf, &result)
	if err != nil {
		return types.FinalResult{}, sdkerrors.Wrap(types.ErrNoFinalResult, "not found")
	}

	return result, nil
}

func (k Keeper) CheckWhitelist(ctx sdk.Context, acc sdk.AccAddress) bool {
	_, found := k.getKey(ctx, types.GetWhitelistKey(acc))
	return found
}

func (k Keeper) GetWhitelist(ctx sdk.Context) (whitelist []string) {
	iterator := k.GetIterator(ctx, types.KeyPrefixWhitelist)
	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		operator := sdk.AccAddress(iterator.Value())
		whitelist = append(whitelist, operator.String())
	}

	return whitelist
}

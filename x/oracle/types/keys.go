package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/address"
)

const (
	// ModuleName defines the module name
	ModuleName = "oracle"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// RouterKey defines the module's message routing key
	RouterKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_oracle"
)

var (
	KeyPrefixWhitelist = []byte{0x01}
	KeyPrefixResultInt = []byte{0x02}
	KeyFinalResult     = []byte{0x03}
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}

func GetWhitelistKey(acc sdk.AccAddress) []byte {
	return append(KeyPrefixWhitelist, address.MustLengthPrefix(acc)...)
}

func GetResultInt(acc sdk.AccAddress) []byte {
	return append(KeyPrefixResultInt, address.MustLengthPrefix(acc)...)
}

package types

// DONTCOVER

import (
	sdkerrors "cosmossdk.io/errors"
)

// x/oracle module sentinel errors
var (
	ErrNotWhitelisted     = sdkerrors.Register(ModuleName, 1100, "Not Whitelisted")
	ErrNoFinalResult      = sdkerrors.Register(ModuleName, 1101, "No Final Result")
	ErrNotAuthorized      = sdkerrors.Register(ModuleName, 1102, "Not authorized")
	ErrInvalidBlockHeight = sdkerrors.Register(ModuleName, 1103, "Invalid block height")
	ErrNoConsensus        = sdkerrors.Register(ModuleName, 1104, "No consensus")
)

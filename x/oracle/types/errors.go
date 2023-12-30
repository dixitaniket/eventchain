package types

// DONTCOVER

import (
	sdkerrors "cosmossdk.io/errors"
)

// x/oracle module sentinel errors
var (
	ErrNotWhitelisted = sdkerrors.Register(ModuleName, 1100, "Not Whitelisted")
)

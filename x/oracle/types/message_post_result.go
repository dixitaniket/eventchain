package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgPostResult = "post_result"

var _ sdk.Msg = &MsgPostResult{}

func NewMsgPostResult(creator string, result string) *MsgPostResult {
	return &MsgPostResult{
		Creator: creator,
		Result:  result,
	}
}

func (msg *MsgPostResult) Route() string {
	return RouterKey
}

func (msg *MsgPostResult) Type() string {
	return TypeMsgPostResult
}

func (msg *MsgPostResult) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgPostResult) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgPostResult) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

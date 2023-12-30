package types

import (
	"errors"
	"fmt"
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	gov1b1 "github.com/cosmos/cosmos-sdk/x/gov/types/v1beta1"
)

const TypeMsgPostResult = "post_result"

var _ sdk.Msg = &MsgPostResult{}

func NewMsgPostResult(creator string, num, toadd int64) *MsgPostResult {
	return &MsgPostResult{
		Creator: creator,
		Result: Result{
			Num:   num,
			Toadd: toadd,
		},
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

func NewMsgGovAddDenoms(
	authority, title, description string,
	operatorAcc sdk.AccAddress,
) *MsgProposeWhitelist {
	return &MsgProposeWhitelist{
		Authority:         authority,
		Title:             title,
		Description:       description,
		WhitelistOperator: operatorAcc.String(),
	}
}

// Type implements Msg interface
func (msg MsgProposeWhitelist) Type() string { return sdk.MsgTypeURL(&msg) }

// GetSignBytes implements Msg
func (msg MsgProposeWhitelist) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(&msg)
	return sdk.MustSortJSON(bz)
}

// GetSigners implements Msg
func (msg MsgProposeWhitelist) GetSigners() []sdk.AccAddress {
	return Signers(msg.Authority)
}

// ValidateBasic implements Msg
func (msg MsgProposeWhitelist) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.WhitelistOperator)
	if err != nil {
		return err
	}
	return ValidateProposal(msg.Title, msg.Description, msg.Authority)
}

func Signers(signers ...string) []sdk.AccAddress {
	as := make([]sdk.AccAddress, len(signers))
	for i := range signers {
		a, _ := sdk.AccAddressFromBech32(signers[i])
		as[i] = a
	}
	return as
}

func ValidateProposal(title, description, authority string) error {
	_, err := sdk.AccAddressFromBech32(authority)
	if err != nil {
		return err
	}

	if len(strings.TrimSpace(title)) == 0 {
		return errors.New("proposal title cannot be blank")
	}
	if len(title) > gov1b1.MaxTitleLength {
		return fmt.Errorf("proposal title is longer than max length of %d", gov1b1.MaxTitleLength)
	}

	if len(description) == 0 {
		return errors.New("proposal description cannot be blank")
	}
	if len(description) > gov1b1.MaxDescriptionLength {
		return fmt.Errorf(
			"proposal description is longer than max length of %d",
			gov1b1.MaxDescriptionLength,
		)
	}

	return nil
}

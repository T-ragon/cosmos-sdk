package types

import (
	sdk "github.com/T-ragon/cosmos-sdk/v3/types"
)

// verify interface at compile time
var (
	_ sdk.Msg = &MsgUnjail{}
	_ sdk.Msg = &MsgUpdateParams{}
)

// NewMsgUnjail creates a new MsgUnjail instance
func NewMsgUnjail(validatorAddr string) *MsgUnjail {
	return &MsgUnjail{
		ValidatorAddr: validatorAddr,
	}
}

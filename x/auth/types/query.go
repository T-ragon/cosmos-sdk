package types

import (
	codectypes "github.com/T-ragon/cosmos-sdk/codec/types"
	sdk "github.com/T-ragon/cosmos-sdk/types"
)

func (m *QueryAccountResponse) UnpackInterfaces(unpacker codectypes.AnyUnpacker) error {
	var account sdk.AccountI
	return unpacker.UnpackAny(m.Account, &account)
}

var _ codectypes.UnpackInterfacesMessage = &QueryAccountResponse{}

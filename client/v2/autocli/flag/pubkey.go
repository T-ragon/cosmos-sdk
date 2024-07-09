package flag

import (
	"context"
	"fmt"

	"google.golang.org/protobuf/reflect/protoreflect"

	"github.com/T-ragon/cosmos-sdk/v3/codec"
	"github.com/T-ragon/cosmos-sdk/v3/codec/types"
	cryptocodec "github.com/T-ragon/cosmos-sdk/v3/crypto/codec"
	cryptotypes "github.com/T-ragon/cosmos-sdk/v3/crypto/types"
)

type pubkeyType struct{}

func (a pubkeyType) NewValue(_ *context.Context, _ *Builder) Value {
	return &pubkeyValue{}
}

func (a pubkeyType) DefaultValue() string {
	return ""
}

type pubkeyValue struct {
	value *types.Any
}

func (a pubkeyValue) Get(protoreflect.Value) (protoreflect.Value, error) {
	return protoreflect.ValueOf(a.value), nil
}

func (a pubkeyValue) String() string {
	return a.value.String()
}

func (a *pubkeyValue) Set(s string) error {
	registry := types.NewInterfaceRegistry()
	cryptocodec.RegisterInterfaces(registry)
	cdc := codec.NewProtoCodec(registry)

	var pk cryptotypes.PubKey
	err := cdc.UnmarshalInterfaceJSON([]byte(s), &pk)
	if err != nil {
		return fmt.Errorf("input isn't a pubkey: %w", err)
	}

	any, err := types.NewAnyWithValue(pk)
	if err != nil {
		return fmt.Errorf("error converting to any type")
	}

	a.value = any

	return nil
}

func (a pubkeyValue) Type() string {
	return "pubkey"
}

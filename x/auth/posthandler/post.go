package posthandler

import (
	sdk "github.com/T-ragon/cosmos-sdk/v3/types"
)

// HandlerOptions are the options required for constructing a default SDK PostHandler.
type HandlerOptions struct{}

// NewPostHandler returns an empty PostHandler chain.
func NewPostHandler(_ HandlerOptions) (sdk.PostHandler, error) {
	postDecorators := []sdk.PostDecorator{}

	return sdk.ChainPostDecorators(postDecorators...), nil
}

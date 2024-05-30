package crosschain_bridge

import (
	"context"
	"github.com/onchainpay/go-sdk/types/requests"
	"github.com/onchainpay/go-sdk/types/responses"
)

func (d *Domain) Create(ctx context.Context, payload requests.CrosschainBridge) responses.BaseResponseGeneric[*responses.CrosschainBridge] {
	var res responses.CrosschainBridge

	_res := d.requester.Request(ctx, "bridge/create", payload.ToMap(), &res)

	return responses.ConvertBase[*responses.CrosschainBridge](_res)
}

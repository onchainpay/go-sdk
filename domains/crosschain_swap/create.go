package crosschain_swap

import (
	"context"
	"onchainpay_sdk/types/requests"
	"onchainpay_sdk/types/responses"
)

func (d *Domain) Create(ctx context.Context, payload requests.CrosschainSwap) responses.BaseResponseGeneric[*responses.CrosschainSwap] {
	var res responses.CrosschainSwap

	_res := d.requester.Request(ctx, "swaps/create", payload.ToMap(), &res)

	return responses.ConvertBase[*responses.CrosschainSwap](_res)
}

package base

import (
	"context"
	"onchainpay_sdk/types/requests"
	"onchainpay_sdk/types/responses"
)

func (d *Domain) PriceRate(ctx context.Context, payload requests.PriceRate) responses.BaseResponseGeneric[*string] {
	var res string

	_res := d.requester.Request(ctx, "price-rate", payload.ToMap(), &res)

	return responses.ConvertBase[*string](_res)
}

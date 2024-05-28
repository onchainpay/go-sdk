package orders

import (
	"context"
	"onchainpay_sdk/types/requests"
	"onchainpay_sdk/types/responses"
)

func (d *Domain) GetOrders(ctx context.Context, payload requests.GetOrders) responses.BaseResponseGeneric[*responses.GetOrders] {
	var res responses.GetOrders

	if payload.Limit == 0 {
		payload.Limit = 100
	}

	_res := d.requester.Request(ctx, "orders", payload.ToMap(), &res)

	return responses.ConvertBase[*responses.GetOrders](_res)
}

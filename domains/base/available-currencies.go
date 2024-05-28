package base

import (
	"context"
	"onchainpay_sdk/types/responses"
)

func (d *Domain) AvailableCurrencies(ctx context.Context) responses.BaseResponseGeneric[*[]responses.AvailableCurrency] {
	res := []responses.AvailableCurrency{}

	_res := d.requester.Request(ctx, "available-currencies", map[string]string{}, &res)

	return responses.ConvertBase[*[]responses.AvailableCurrency](_res)
}

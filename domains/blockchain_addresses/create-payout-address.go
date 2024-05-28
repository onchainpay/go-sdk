package blockchain_addresses

import (
	"context"
	"onchainpay_sdk/types/requests"
	"onchainpay_sdk/types/responses"
)

func (d *Domain) CreatePayoutAddress(ctx context.Context, payload requests.CreatePayOutAddress) responses.BaseResponseGeneric[*responses.PayOutAddress] {
	var res responses.PayOutAddress

	_res := d.requester.Request(ctx, "payout-address", payload.ToMap(), &res)

	return responses.ConvertBase[*responses.PayOutAddress](_res)
}

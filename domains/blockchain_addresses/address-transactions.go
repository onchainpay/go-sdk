package blockchain_addresses

import (
	"context"
	"onchainpay_sdk/types/requests"
	"onchainpay_sdk/types/responses"
)

func (d *Domain) AddressTransactions(ctx context.Context, payload *requests.AddressTransactions) responses.BaseResponseGeneric[*[]responses.AddressTransaction] {
	res := []responses.AddressTransaction{}

	_res := d.requester.Request(ctx, "addresses/address-transactions", payload.ToMap(), &res)

	return responses.ConvertBase[*[]responses.AddressTransaction](_res)
}

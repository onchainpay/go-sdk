package address_book

import (
	"context"
	"onchainpay_sdk/types/requests"
	"onchainpay_sdk/types/responses"
)

func (d *Domain) Update(ctx context.Context, payload requests.UpdateAddressBook) responses.BaseResponse {
	var res string

	_res := d.requester.Request(ctx, "address-book/update", payload.ToMap(), &res)

	return _res
}

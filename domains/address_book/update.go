package address_book

import (
	"context"
	"github.com/onchainpay/go-sdk/types/requests"
	"github.com/onchainpay/go-sdk/types/responses"
)

func (d *Domain) Update(ctx context.Context, payload requests.UpdateAddressBook) responses.BaseResponse {
	var res string

	_res := d.requester.Request(ctx, "address-book/update", payload.ToMap(), &res)

	return _res
}

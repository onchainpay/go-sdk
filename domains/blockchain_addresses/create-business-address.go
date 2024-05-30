package blockchain_addresses

import (
	"context"
	"github.com/onchainpay/go-sdk/types/requests"
	"github.com/onchainpay/go-sdk/types/responses"
)

func (d *Domain) CreateBusinessAddress(ctx context.Context, payload requests.CreateBusinessAddress) responses.BaseResponseGeneric[*responses.BusinessAddress] {
	var res responses.BusinessAddress

	_res := d.requester.Request(ctx, "business-address", payload.ToMap(), &res)

	return responses.ConvertBase[*responses.BusinessAddress](_res)
}

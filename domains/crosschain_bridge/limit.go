package crosschain_bridge

import (
	"context"
	"onchainpay_sdk/types/responses"
)

func (d *Domain) Limit(ctx context.Context) responses.BaseResponseGeneric[*responses.CrosschainLimit] {
	var res responses.CrosschainLimit

	_res := d.requester.Request(ctx, "bridge/limit", map[string]string{}, &res)

	return responses.ConvertBase[*responses.CrosschainLimit](_res)
}

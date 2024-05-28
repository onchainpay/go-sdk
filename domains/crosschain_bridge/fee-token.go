package crosschain_bridge

import (
	"context"
	"onchainpay_sdk/types/requests"
	"onchainpay_sdk/types/responses"
)

func (d *Domain) FeeToken(ctx context.Context, payload requests.CrosschainBridgeFeeToken) responses.BaseResponseGeneric[*responses.CrosschainBridgeFeeToken] {
	var res responses.CrosschainBridgeFeeToken

	_res := d.requester.Request(ctx, "bridge/fee-token", payload.ToMap(), &res)

	return responses.ConvertBase[*responses.CrosschainBridgeFeeToken](_res)
}

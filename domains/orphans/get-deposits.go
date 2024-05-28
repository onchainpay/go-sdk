package orphans

import (
	"context"
	"onchainpay_sdk/types/requests"
	"onchainpay_sdk/types/responses"
)

func (d *Domain) GetDeposits(ctx context.Context, payload requests.GetOrphansDeposits) responses.BaseResponseGeneric[*[]responses.OrphanTransaction] {
	res := []responses.OrphanTransaction{}

	_res := d.requester.Request(ctx, "orphan-deposits/get-deposits", payload.ToMap(), &res)

	return responses.ConvertBase[*[]responses.OrphanTransaction](_res)
}

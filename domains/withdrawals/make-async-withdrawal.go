package withdrawals

import (
	"context"
	"onchainpay_sdk/types/requests"
	"onchainpay_sdk/types/responses"
)

func (d *Domain) MakeWithdrawalAsync(ctx context.Context, payload requests.CreateAsyncWithdrawal) responses.BaseResponseGeneric[*responses.AsyncWithdrawal] {
	var res responses.AsyncWithdrawal

	_res := d.requester.Request(ctx, "make-withdrawal-async", payload.ToMap(), &res)

	return responses.ConvertBase[*responses.AsyncWithdrawal](_res)
}

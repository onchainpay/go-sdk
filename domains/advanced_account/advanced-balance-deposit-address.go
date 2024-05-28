package advanced_account

import (
	"context"
	"onchainpay_sdk/types/requests"
	"onchainpay_sdk/types/responses"
)

func (d *Domain) AdvancedBalanceDepositAddress(ctx context.Context, payload requests.AdvancedBalanceDepositAddress) responses.BaseResponseGeneric[*responses.AdvancedBalanceDepositAddress] {
	var res responses.AdvancedBalanceDepositAddress

	_res := d.requester.Request(ctx, "advanced-balance-deposit-address", payload.ToMap(), &res)

	return responses.ConvertBase[*responses.AdvancedBalanceDepositAddress](_res)
}

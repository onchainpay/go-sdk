package advanced_account

import (
	"context"
	"github.com/onchainpay/go-sdk/types/requests"
	"github.com/onchainpay/go-sdk/types/responses"
)

func (d *Domain) AdvancedBalanceDepositAddress(ctx context.Context, payload requests.AdvancedBalanceDepositAddress) responses.BaseResponseGeneric[*responses.AdvancedBalanceDepositAddress] {
	var res responses.AdvancedBalanceDepositAddress

	_res := d.requester.Request(ctx, "advanced-balance-deposit-address", payload.ToMap(), &res)

	return responses.ConvertBase[*responses.AdvancedBalanceDepositAddress](_res)
}

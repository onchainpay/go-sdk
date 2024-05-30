package invoices

import (
	"context"
	"github.com/onchainpay/go-sdk/types/requests"
	"github.com/onchainpay/go-sdk/types/responses"
)

func (d *Domain) MakeInvoice(ctx context.Context, payload requests.CreateInvoice) responses.BaseResponseGeneric[*responses.Invoice] {
	var res responses.Invoice

	_res := d.requester.Request(ctx, "make-invoice", payload.ToMap(), &res)

	return responses.ConvertBase[*responses.Invoice](_res)
}

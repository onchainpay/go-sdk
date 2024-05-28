package invoices

import (
	"context"
	"onchainpay_sdk/types/requests"
	"onchainpay_sdk/types/responses"
)

func (d *Domain) GetInvoices(ctx context.Context, payload requests.GetInvoices) responses.BaseResponseGeneric[*responses.InvoicesList] {
	var res responses.InvoicesList

	_res := d.requester.Request(ctx, "get-invoices", payload.ToMap(), &res)

	return responses.ConvertBase[*responses.InvoicesList](_res)
}

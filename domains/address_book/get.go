package address_book

import (
	"context"
	"onchainpay_sdk/types/requests"
	"onchainpay_sdk/types/responses"
)

func (d *Domain) Get(ctx context.Context, payload requests.GetAddressBook) responses.BaseResponseGeneric[*responses.AddressBookList] {
	var res responses.AddressBookList

	_res := d.requester.Request(ctx, "address-book/get", payload.ToMap(), &res)

	return responses.ConvertBase[*responses.AddressBookList](_res)
}

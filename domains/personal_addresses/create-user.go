package personal_addresses

import (
	"context"
	"onchainpay_sdk/types/requests"
	"onchainpay_sdk/types/responses"
)

func (d *Domain) CreateUser(ctx context.Context, payload requests.CreatePersonalUser) responses.BaseResponseGeneric[*responses.PersonalUserWithAddresses] {
	var res responses.PersonalUserWithAddresses

	_res := d.requester.Request(ctx, "personal_addresses/create-user", payload.ToMap(), &res)

	return responses.ConvertBase[*responses.PersonalUserWithAddresses](_res)
}

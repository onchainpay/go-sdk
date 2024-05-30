package crosschain_bridge

import "github.com/onchainpay/go-sdk/requester"

type Domain struct {
	requester *requester.Requester
}

func New(requester *requester.Requester) *Domain {
	return &Domain{
		requester: requester,
	}
}

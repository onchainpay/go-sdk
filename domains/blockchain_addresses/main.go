package blockchain_addresses

import "github.com/onchainpay/go-sdk/requester"

type Domain struct {
	requester *requester.Requester
}

func New(requester *requester.Requester) *Domain {
	return &Domain{
		requester: requester,
	}
}

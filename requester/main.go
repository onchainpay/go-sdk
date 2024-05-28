package requester

import "onchainpay_sdk/noncer"

type Requester struct {
	publicKey string
	secretKey string

	noncer *noncer.Noncer
}

func New(public, secret string) *Requester {
	return &Requester{
		publicKey: public,
		secretKey: secret,

		noncer: noncer.New(),
	}
}

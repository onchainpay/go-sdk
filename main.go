package onchainpay_sdk

import (
	"onchainpay_sdk/domains/address_book"
	"onchainpay_sdk/domains/advanced_account"
	"onchainpay_sdk/domains/auto_swaps"
	"onchainpay_sdk/domains/base"
	"onchainpay_sdk/domains/blockchain_addresses"
	"onchainpay_sdk/domains/crosschain_bridge"
	"onchainpay_sdk/domains/crosschain_swap"
	"onchainpay_sdk/domains/invoices"
	"onchainpay_sdk/domains/orders"
	"onchainpay_sdk/domains/orphans"
	"onchainpay_sdk/domains/personal_addresses"
	"onchainpay_sdk/domains/webhooks"
	"onchainpay_sdk/domains/withdrawals"
	"onchainpay_sdk/noncer"
	"onchainpay_sdk/requester"
)

type Client struct {
	noncer *noncer.Noncer

	AddressBook         *address_book.Domain
	AdvancedBalance     *advanced_account.Domain
	AutoSwaps           *auto_swaps.Domain
	Base                *base.Domain
	BlockchainAddresses *blockchain_addresses.Domain
	Bridge              *crosschain_bridge.Domain
	Swaps               *crosschain_swap.Domain
	Invoices            *invoices.Domain
	Orders              *orders.Domain
	Orphans             *orphans.Domain
	PersonalAddresses   *personal_addresses.Domain
	Webhooks            *webhooks.Domain
	Withdrawals         *withdrawals.Domain
}

func New(public, private string) *Client {
	_requester := requester.New(public, private)

	return &Client{
		noncer: noncer.New(),

		AddressBook:         address_book.New(_requester),
		AdvancedBalance:     advanced_account.New(_requester),
		AutoSwaps:           auto_swaps.New(_requester),
		Base:                base.New(_requester),
		BlockchainAddresses: blockchain_addresses.New(_requester),
		Bridge:              crosschain_bridge.New(_requester),
		Swaps:               crosschain_swap.New(_requester),
		Invoices:            invoices.New(_requester),
		Orders:              orders.New(_requester),
		Orphans:             orphans.New(_requester),
		PersonalAddresses:   personal_addresses.New(_requester),
		Webhooks:            webhooks.New(_requester),
		Withdrawals:         withdrawals.New(_requester),
	}
}

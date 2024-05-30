# OnchainPay SDK

OnchainPay SDK is a comprehensive JavaScript library designed to streamline the integration of blockchain-based payment solutions into various applications. This SDK provides developers with the tools necessary to facilitate secure, transparent, and efficient transactions on supported blockchain networks.

Key Features:

- **Ease of Integration**: Simplifies the process of adding blockchain payment capabilities to your applications.
- **Security**: Ensures high-level security for all transactions using blockchain technology.
- **Transparency**: Leverages the transparency of blockchain networks to provide clear and verifiable transaction records.
- **Multi-Network Support**: Supports multiple blockchain networks, providing flexibility and scalability.

For detailed documentation, installation guides, and API references, please visit OnchainPay Documentation.

This package makes it easy [OnChainPay Api](https://docs.onchainpay.io/).

## Installation

`go get github.com/onchainpay/go-sdk`

## Use

Go to your personal account
[https://ocp.onchainpay.io/api-keys](https://ocp.onchainpay.io/api-keys)
and get api-keys.

*Substitute keys in class call:*

```go
package main

import (
	"context"
	onchainpay_sdk "github.com/onchainpay/go-sdk"
)

func main() {
	client := onchainpay_sdk.New("__PUBLIC_KEY__", "__PRIVATE_KEY__")
}
```

### Check signature

You can test your signature within this method.

```go
checkSignature := client.Base.CheckSignature(context.Background())

if !checkSignature.Success {
	panic(checkSignature.Error)
}

if checkSignature.Result.CheckSignatureResult {
    fmt.Println("Signature correct")
} else {
    fmt.Println("Signature incorrect", checkSignature.Result.Errors)
}
```

### Fetch available currencies

Get list of available currencies for depositing/withdrawing

```go
availableCurrencies := client.Base.AvailableCurrencies(context.Background())

if !availableCurrencies.Success {
    panic(availableCurrencies.Error)
}

list := *availableCurrencies.Result

for _, currency := range list {
	fmt.Println("%s (%s) = %s", currency.Currency, currency.Alias, currency.PriceUSD)

    if len(currency.Networks) > 0 {
        fmt.Println("\tnetworks:");

		for _, network := range currency.Networks {
			fmt.Println("\t\t%s (%s)", network.Name, network.Alias)
        }
    }
}
```

### Get currencies price-rate

Get price rate from one currency to another


```go
price := client.Base.PriceRate(context.Background(), "ETH", "USDT")

if !price.Success {
    panic(price.Error)
}

fmt.Println("Price: %s", price.Result)
```

### Get advanced balances info

Get info about advanced balance by its id

```go
balance := client.AdvancedBalance.AdvancedBalance(context.Background(), balanceId)

if !balance.Success {
    panic(balance.Error)
}

fmt.Println("[%s] (%s)\n\tAvalable for deposit: %s", balance.Result.AdvancedBalanceId, balance.Result.Currency, string.Join(balance.Result.AvailableCurrenciesForDeposit, ", "))
```

Or get list of advanced balances of user

```go
balances := client.AdvancedBalance.AdvancedBalance(context.Background(), balanceId)

if !balances.Success {
    panic(balances.Error)
}

list := *balances.Result

for _, balance := range list {
    fmt.Println("[%s] (%s)\n\tAvalable for deposit: %s", balance.Result.AdvancedBalanceId, balance.Result.Currency, string.Join(balance.Result.AvailableCurrenciesForDeposit, ", "))
}
```
# Fincra Golang SDK

An unofficial Go SDK for working with [Fincra API](https://fincra.com/)

[![fincra-go-sdk](https://github.com/neymarjimoh/fincra-go/actions/workflows/go.yml/badge.svg)](https://github.com/neymarjimoh/fincra-go/actions/workflows/go.yml)

## Installation

#### To install, run:

```sh
go get github.com/neymarjimoh/fincra-go
```

## Usage

#### Import the package:

```sh
import "github.com/neymarjimoh/fincra-go"
```

#### Initialize a new client:

```go
import (
    "time"
    fincra "github.com/neymarjimoh/fincra-go"
)

const (
    SECRET_KEY = "<Your secret key>"
)

// to indicate environment when integrating, use `WithSandbox()`
newClient := fincra.NewClient(SECRET_KEY, fincra.WithSandbox(true))

// to indicate http request timeout to use, use `WithTimeout()`
newClient := fincra.NewClient(SECRET_KEY, fincra.WithTimeout(5*time.Second))
```

Note:

- The `WithSandbox` argument is optional. If not specified, it defaults to false and you will be using the **Production(live)** API. For example:
  `newClient := fincra.NewClient(SECRET_KEY)`
- The `WithTimeout` argument is optional. If not specified, it defaults to 60 seconds
- Get your keys from your Fincra dashboard and be sure to add it as environment variables

## Functionalities Currently Supported

### 1. Business

#### - Get business details:

This method lets you retrieve the unique Identifier of your business and other information such as your email etc.

Usage example:

```go
resp, err := client.GetBusinessId()
```

### 2. Beneficiaries

#### - Create a beneficiary:

This method is used for creating a Beneficiary.

Usage example:

```go
data := &fincra.CreateBeneficiaryBody{
  FirstName: "efe",
  LastName: "ebieroma",
  Email: "efe@hahaha.com",
  PhoneNumber: "09090909090",
  AccountHolderName: "efe stephen ebieroma",
  Bank: fincra.Bank{
    Name: "Wema Bank",
    Code: "06",
    SortCode: "928927",
    Branch: "Ota",
    Address: fincra.Address{
      Country: "GB",
      State: "Lagos",
      Zip: "123455",
      City: "Paris",
      Street: "Osapa London",
    },
  },
  Type: fincra.Individual, // or Corporate
  Currency: "GBP",
  PaymentDestination: fincra.CryptoWallet,
  UniqueIdentifier: "4",
  DestinationAddress: "Osapa London",
  BusinessId: "6457d39b12b4401f99a54772",
}

resp, err := client.CreateBeneficiary(data)
```

**NOTE**: PaymentDestination accepts an enum of CryptoWallet (crypto_wallet), BankAccount (bank_account) or MobileMoneyWallet (mobile_money_wallet)

#### - Fetch a beneficiary:

This method is used for retrieving a single beneficiary attached to a business.

Usage example:

```go
data := &fincra.GetAllBeneficiariesParams{
  BusinessId: "617fefbe4a68ec99ba6af0be",
  Page: "1", // optional, defaults to 1
  PerPage: "20", // optional, defaults to 10
}

resp, err := client.GetAllBeneficiaries(data);
```

#### - List beneficiaries:

This method is used for retrieving a single beneficiary attached to a business.

Usage example:

```go
data := &fincra.GetBeneficiaryParams{
  BusinessId: "617fefbe4a68ec99ba6af0be",
  BeneficiaryId: "618fefbe4a68ec99ba5af0be",
}

resp, err := client.GetBeneficiary(data);
```

#### - Update a beneficiary:

This method is used for updating a Beneficiary.

Usage example:

```go
data := &fincra.UpdateBeneficiaryBody{
  FirstName: "efe",
  LastName: "ebieroma",
  Email: "efe@hahaha.com",
  PhoneNumber: "09090909090",
  AccountHolderName: "efe stephen ebieroma",
  DestinationAddress: "Osapa London",
  BusinessId: "6457d39b12b4401f99a54772",
  BeneficiaryId: "646db15a8cdec23981165184",
}

resp, err := client.UpdateBeneficiary(data)
```

More details about the parameters for the above method [here](https://docs.fincra.com/reference/update-a-beneficiary)

#### - Delete a beneficiary:

This method is used for deleting a beneficiary.

Usage example:

```go
data := &fincra.GetBeneficiaryParams{
  BusinessId: "617fefbe4a68ec99ba6af0be",
  BeneficiaryId: "618fefbe4a68ec99ba5af0be",
}

resp, err := client.DeleteBeneficiary(data);
```

### 3. Conversions

#### - Convert a currency:

This method can convert one currency to another provided that it's a supported conversion currency e.g NGN to USD.

Usage example:

```go
data := &fincra.CreateConversionBody{
  BusinessId: "617fefbe4a68ec99ba6af0be",
  QuoteReference: "124246677268282782728",
}

resp, err := client.CreateConversion(data);
```

#### - List conversions:

This method provides a list of all conversions performed by a business.

Usage example:

```go
businessId := "617fefbe4a68ec99ba6af0be"
resp, err := client.GetBusinessConversions(businessId);
```

#### - Fetch a conversion:

This method fetches a specific conversion performed by a parent Business or sub account.

Usage example:

```go
conversionId := "617fefbe4a68ec99ba6af0bh"
resp, err := client.GetConversion(conversionId);
```

### 4. Quotes

The Quotes service provides a method that allows you to generate quotes for Real-time transactions occurring on your integration.

#### - Create a quote:

This method is used for generating a quote.

Usage example:

```go
data := &fincra.CreateQuoteBody{
  Action:              "send",
  TransactionType:     "conversion",
  FeeBearer:           "business",
  PaymentDestination:  fincra.FliqPayWallet,
  BeneficiaryType:     fincra.Individual,
  Business:            "6457d39b12b4401f99a54772",
  Amount:              "150",
  DestinationCurrency: "USD",
  SourceCurrency:      "NGN",
}

resp, err := client.CreateQuote(data);
```

### 5. Chargebacks

#### - List chargebacks:

This method lets you list all the chargebacks incurred on your account.

Usage example:

```go
businessId := "6457d39b12b4401f99a54772"
resp, err := client.ListChargeBacks(businessId)
```

#### - Accept a chargeback:

This method lets you accept a chargeback

Usage example:

```go
data := &fincra.AcceptChargeBackDto{
  BusinessId: "617fefbe4a68ec99ba6af0be",
  ChargeBackId: "7171892",
}

resp, err := client.AcceptChargeBack(data)
```

#### - Reject a chargeback:

This method lets you reject a chargeback

Usage example:

```go
data := &fincra.RejectChargeBackDto{
  BusinessId: "617fefbe4a68ec99ba6af0be",
  ChargeBackId: "7171892",
  Reason: "no money on ground",
}

resp, err := client.RejectChargeBack(data)
```

### 6. Wallets (Balance)

#### - Fetch all balances:

This method lists all the account balance information of a business

Usage example:

```go
businessId := "6457d39b12b4401f99a54772"
resp, err := client.ListWallets(businessId)
```

#### - Fetch a blance:

This method provides information to the merchant about a specific account balance

Usage example:

```go
walletId := "66433"
resp, err := client.ListWallet(walletId)
```

#### - List account balance logs:

This method fetches all pay-ins and pay-outs that occurred on your integration

Usage example:

```go
data := fincra.LogsDto{
 Business: "6457d39b12b4401f99a54772",
 Action:   fincra.Credit,
 Page:     "2",
 PerPage:  "10",
 Amount:   "500",
}

resp, err := client.ListWalletLogs(data)
```

### 7. Identity Management

#### - Verify account number:

This method lets you verify a bank account

Usage example:

```go
data := fincra.VerifyBankAccountBody{
 AccountNumber: "0929292929",
 Type:          fincra.Nuban, // it can be fincra.Iban
 BankCode:      "044",
 Iban:          "999",
}

resp, err := client.VerifyBankAccount(data)
```

#### - BVN Resolution:

This method lets you verify a bank account

Usage example:

```go
data := fincra.VerifyBVNBody{
 Bvn:      "09292929221",
 Business: "6457d39b12b4401f99a54772",
}

resp, err := client.VerifyBVN(data)
```

### 8. Virtual Account

#### - Create a virtual account [here](https://docs.fincra.com/reference/request-virtual-accounts):

This method lets you create NGN and MCY virtual accounts. You’ll need to complete currency-specific fields as well as shared fields applicable to all currencies. E.g firstName and lastName are applicable to both NGN and MCY account requests while meansofID is applicable to only MCY accounts.

Usage example:

```go
data := fincra.CreateVirtualAccountDto{
 Currency:    "NGN",
 UtilityBill: "https://www.planetware.com/wpimages/2020/02/france-in-pictures-beautiful-places-to-photograph-eiffel-tower.jpg",
 AccountType: "individual",
 KYCInformation: fincra.KYCInformationDto{
  FirstName:    "John",
  LastName:     "Doe",
  Email:        "abc@abc.com",
  BusinessName: "JohnDoe",
 },
 Channel: "providus",
}

resp, err := client.CreateVirtualAccount(data)
```

#### - List virtual accounts [here](https://docs.fincra.com/reference/get-merchant-virtual-account-requests):

This method fetches all virtual accounts belonging to a merchant.

Usage example:

```go
options := fincra.Options{
 Currency:    "NGN",
}

resp, err := client.ListVirtualAccounts(options)
```

_N/B_: In `Options`, one or more of `Currency`, `BusinessName`, `IssuedDate`, `RequestedDate`, `AccountNumber` and `Status` must be passed as payload.

#### - List virtual account requests [here](https://docs.fincra.com/reference/get-virtual-account-requests):

This method is used for getting all account requests belonging to a merchant.

Usage example:

```go
resp, err := client.ListVirtualAccountRequests()
```

#### - Fetch a virtual account by currency [here](https://docs.fincra.com/reference/get-merchant-virtual-account-by-currency):

This method is used for retrieving an account that is belongs to a merchant by currency.

Usage example:

```go
resp, err := client.ListVirtualAccountByCurrency("EUR")
```

#### - Fetch a virtual account by BVN [here](https://docs.fincra.com/reference/fetch-a-virtual-account-by-bvn):

This method is used for retrieving an account that is belongs to a merchant by BVN.

Usage example:
Accepts two parameters. First parameter represents the BVN and second represents the business ID.

```go
resp, err := client.ListVirtualAccountByBvn("0123456789", "6457d39b12b4401f99a54772")
```

#### - Fetch a virtual account [here](https://docs.fincra.com/reference/get-one-virtual-account):

This method is used for retrieving a virtual account.

Usage example:
Accepts a parameter that represents the Virtual Account ID.

```go
resp, err := client.ListVirtualAccount("6457d39b12b4401f99a54772")
```

### Todos:

- [ ] Payins endpoints, tests and update README
- [ ] Payouts endpoints, tests and update README
- [ ] Checkouts endpoints, tests and update README


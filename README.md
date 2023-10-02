# Fincra Golang SDK
An unofficial Go SDK for working with [Fincra API](https://fincra.com/)

## Installation
#### To install, run:

```
go get github.com/neymarjimoh/fincra-go
```

## Usage
#### Import the package:

```
import "github.com/neymarjimoh/fincra-go"
```

#### Initialize a new client:

```
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

```
resp, err := client.GetBusinessId()
```

### 2. Beneficiaries
#### - Create a beneficiary:
This method is used for creating a Beneficiary.

Usage example:

```
data := &fincra.CreateBeneficiaryBody{
  FirstName: 'efe',
  LastName: 'ebieroma',
  Email: 'efe@hahaha.com',
  PhoneNumber: '09090909090',
  AccountHolderName: 'efe stephen ebieroma',
  Bank: fincra.Bank{
    Name: 'Wema Bank',
    Code: '06',
    SortCode: '928927',
    Branch: 'Ota',
    Address: fincra.Address{
      Country: 'GB',
      State: 'Lagos',
      Zip: '123455',
      City: 'Paris',
      Street: 'Osapa London',
    },
  },
  Type: fincra.Individual, // or Corporate 
  Currency: 'GBP',
  PaymentDestination: fincra.CryptoWallet,
  UniqueIdentifier: '4',
  DestinationAddress: 'Osapa London',
  BusinessId: '6457d39b12b4401f99a54772',
}

resp, err := client.CreateBeneficiary(data)
```

**NOTE**: PaymentDestination accepts an enum of CryptoWallet (crypto_wallet), BankAccount (bank_account) or MobileMoneyWallet (mobile_money_wallet)

#### - Fetch a beneficiary:
This method is used for retrieving a single beneficiary attached to a business.

Usage example:

```
data := &fincra.GetAllBeneficiariesParams{
  BusinessId: '617fefbe4a68ec99ba6af0be',
  Page: '1', // optional, defaults to 1
  PerPage: '20', // optional, defaults to 10
}

resp, err := client.GetAllBeneficiaries(data);
```

#### - List beneficiaries:
This method is used for retrieving a single beneficiary attached to a business.

Usage example:

```
data := &fincra.GetBeneficiaryParams{
  BusinessId: '617fefbe4a68ec99ba6af0be',
  BeneficiaryId: '618fefbe4a68ec99ba5af0be',
}

resp, err := client.GetBeneficiary(data);
```

#### - Update a beneficiary:
This method is used for updating a Beneficiary.

Usage example:

```
data := &fincra.UpdateBeneficiaryBody{
  FirstName: 'efe',
  LastName: 'ebieroma',
  Email: 'efe@hahaha.com',
  PhoneNumber: '09090909090',
  AccountHolderName: 'efe stephen ebieroma',
  DestinationAddress: 'Osapa London',
  BusinessId: '6457d39b12b4401f99a54772',
  BeneficiaryId: '646db15a8cdec23981165184',
}

resp, err := client.UpdateBeneficiary(data)
```

More details about the parameters for the above method [here](https://docs.fincra.com/reference/update-a-beneficiary)

#### - Delete a beneficiary:
This method is used for deleting a beneficiary.

Usage example:

```
data := &fincra.GetBeneficiaryParams{
  BusinessId: '617fefbe4a68ec99ba6af0be',
  BeneficiaryId: '618fefbe4a68ec99ba5af0be',
}

resp, err := client.DeleteBeneficiary(data);
```

### 3. Conversions
#### - Convert a currency:
This method can convert one currency to another provided that it's a supported conversion currency e.g NGN to USD.

Usage example:

```
data := &fincra.CreateConversionBody{
  BusinessId: '617fefbe4a68ec99ba6af0be',
  QuoteReference: '124246677268282782728',
}

resp, err := client.CreateConversion(data);
```

#### - List conversions:
This method provides a list of all conversions performed by a business.

Usage example:

```
businessId := "617fefbe4a68ec99ba6af0be"
resp, err := client.GetBusinessConversions(businessId);
```


#### - Fetch a conversion:
This method fetches a specific conversion performed by a parent Business or sub account.

Usage example:

```
conversionId := "617fefbe4a68ec99ba6af0bh"
resp, err := client.GetConversion(conversionId);
```

### 4. Quotes
The Quotes service provides a method that allows you to generate quotes for Real-time transactions occurring on your integration.
#### - Create a quote:
This method is used for generating a quote.

Usage example:

```
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

```
businessId := "6457d39b12b4401f99a54772"
resp, err := client.ListChargeBacks(businessId)
```

#### - Accept a chargeback:
This method lets you accept a chargeback

Usage example:

```
data := &fincra.AcceptChargeBackDto{
  BusinessId: "617fefbe4a68ec99ba6af0be",
  ChargeBackId: "7171892",
}

resp, err := client.AcceptChargeBack(data)
```

#### - Reject a chargeback:
This method lets you reject a chargeback

Usage example:

```
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

```
businessId := "6457d39b12b4401f99a54772"
resp, err := client.ListWallets(businessId)
```

#### - Fetch a blance:
This method provides information to the merchant about a specific account balance

Usage example:

```
walletId := "66433"
resp, err := client.ListWallet(walletId)
```

#### - List account balance logs:
This method fetches all pay-ins and pay-outs that occurred on your integration

Usage example:

```
data := fincra.LogsDto{
 Business: "6457d39b12b4401f99a54772",
 Action:   fincra.Credit,
 Page:     "2",
 PerPage:  "10",
 Amount:   "500",
}

resp, err := client.ListWalletLogs(data)
```

### 6. Identity Management
#### - Verify account number:
This method lets you verify a bank account

Usage example:

```
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

```
data := fincra.VerifyBVNBody{
 Bvn:      "09292929221",
 Business: "6457d39b12b4401f99a54772",
}

resp, err := client.VerifyBVN(data)
```

### Todos:
- [x] Set up project with base client wrapper done, write tests
- [x] update README with setup
- [x] Business endpoints, tests and update README
- [x] Beneficiaries endpoints, tests and update README
- [x] Conversions endpoints, tests and update README
- [x] Add context to the client methods
- [x] Quotes endpoints, tests and update README
- [x] Chargebacks endpoints, tests and update README
- [x] Wallets endpoints, tests and update README
- [x] Verification endpoints, tests and update README
- [ ] Collections endpoints, tests and update README
- [ ] Set up CI/CD with PR templates and auto tests for every build
- [ ] Payouts endpoints, tests and update README
- [ ] add faker emails, ids etc to test for avoiding duplicate error
- [ ] Virtual accounts endpoints, tests and update README
- [ ] Add link to contributions guide on README
- [ ] Update README completely for easy usage (re-arrange based on the API reference)
- [ ] Test as a user and see how it works, fix bug fixes and prepare for first release
- [ ] Make public and stay jiggy with more grinding
- [ ] Build and deploy a release after review from some devs

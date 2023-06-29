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
    fincra "github.com/neymarjimoh/fincra-go"
)

const (
    SECRET_KEY = "<Your secret key>"
)

newClient := fincra.NewClient(SECRET_KEY, fincra.WithSandbox(true))
```
Note: 
- The second argument is optional. If not specified, it defaults to false and you will be using the **Production(live)** API. For example:
`newClient := fincra.NewClient(SECRET_KEY)`
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

resp, err := client.AcceptChargeBack(data);
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

resp, err := client.RejectChargeBack(data);
```

### Todos:
- [x] Set up project with base client wrapper done, write tests
- [x] update README with setup
- [x] Business endpoints, tests and update README
- [x] Beneficiaries endpoints, tests and update README
- [ ] add faker emails, ids etc to test for avoiding dplicate error
- [x] Conversions endpoints, tests and update README
- [x] Quotes endpoints, tests and update README
- [ ] Collections endpoints, tests and update README
- [ ] Set up CI/CD with PR templates and auto tests for every build
- [ ] Payouts endpoints, tests and update README
- [x] Chargebacks endpoints, tests and update README
- [ ] Verification endpoints, tests and update README
- [ ] Subaccounts endpoints, tests and update README
- [ ] Virtual accounts endpoints, tests and update README
- [ ] Wallets endpoints, tests and update README
- [ ] Add link to contributions guide on README
- [ ] Update README completely for easy usage
- [ ] Test as a user and see how it works, fix bug fixes and prepare for first release
- [ ] Make public and stay jiggy with more grinding
- [ ] Build and deploy a release after review from some devs

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
This method lets you retrieves the unique Identifier of your business and other information such as your email etc.

Usage example:
```
resp, err := client.GetBusinessId()
```

### Todos:
- [x] Set up project with base client wrapper done, write tests
- [x] update README with setup
- [x] Business endpoints, tests and update README
- [ ] Beneficiaries endpoints, tests and update README
- [ ] Set up CI/CD with PR templates and auto tests for every build
- [ ] Chargebacks endpoints, tests and update README
- [ ] Conversions endpoints, tests and update README
- [ ] Collections endpoints, tests and update README
- [ ] Payouts endpoints, tests and update README
- [ ] Quotes endpoints, tests and update README
- [ ] Subaccounts endpoints, tests and update README
- [ ] Verification endpoints, tests and update README
- [ ] Subaccounts endpoints, tests and update README
- [ ] Virtual accounts endpoints, tests and update README
- [ ] Wallets endpoints, tests and update README
- [ ] Add link to contributions guide on README
- [ ] Update README completely for easy usage
- [ ] Test as a user and see how it works, fix bug fixes and prepare for first release
- [ ] Make public and stay jiggy with more grinding
- [ ] Build and deploy a release after review from some devs

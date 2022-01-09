# veil-api-go

veil-api-go is a Go client library for accessing the [ECP VeiL REST API](https://veil.mashtab.org/docs/latest/api/).

## Использование

```
import "github.com/jsc-masshtab/veil-api-go"
```

First step is constructing ECP VeiL client which allows to use API services.
You can generate token in [ECP VeiL Panel](https://veil.mashtab.org/docs/latest/base/operator_guide/security/users/#_11).

```
client := NewClient("token should be here")
account, _, err := client.Account.Get()
```

Some operations with scalets can be started both sync and async.

```
// Second argument is "wait" which expects boolean value
// true - if you want to wait until the end of operation
// false - if you want this operation to be handled in background
client := NewClient("token should be here")
scalet, _, err := client.Scalet.Rebuild(11111, true)
```

## Tests

You can run tests which make requests straightly to ECP VeiL API.
For now they can't be run together. Run specific test if you want to test some method.

```bash
$ go test -v github.com/jsc-masshtab/veil-api-go -run TestDomainService_List
```

For convenience, you can use "VEIL_API_TOKEN" env for not passing token to every test.

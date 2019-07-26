# Biwse Go Library
[![GoDoc](http://img.shields.io/badge/godoc-reference-blue.svg)](http://godoc.org/github.com/biwse/biwse-go)

This API client generated from OpenAPI specification. 

## Installation

```shell
go get -u github.com/biwse/biwse-go
```

## Example


```golang
import "github.com/biwse/biwse-go"

...
// Prepare client
token := "<api_token>"
config := biwse.NewConfiguration()
ctx := context.WithValue(context.Background(), biwse.ContextAccessToken, token)
client := NewAPIClient(config)
//Call any API endpoint
balance, response, err := client.AppApi.GetBalance(ctx, "<app_id>", "btc")
```



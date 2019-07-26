# Go API client for biwse.com

This API client generated from OpenAPI specification. 

## Installation

```shell
go get github.com/biwse/biwse-go
```

## Example


```golang
import "github.com/biwse/biwse-go"

...

token := "<api_token>"
config := biwse.NewConfiguration()
ctx := context.WithValue(context.Background(), biwse.ContextAccessToken, token)
client := NewAPIClient(config)
balance, response, err := client.AppApi.GetBalance(ctx, "<app_id>", "btc")
```


## Author




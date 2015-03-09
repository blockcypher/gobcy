# blockcy

A Go wrapper for the [BlockCypher](http://blockcypher.com/) API. Targeting support for Bitcoin (main and testnet3) and BlockCypher's internal testnet.

## Configuration

Import the package like so:

```go
import "github.com/acityinohio/blockcy"
```

Then set your configuration variables:

```go
blockcy.Config.Coin = "btc" //only supports "btc" or "bcy"
blockcy.Config.Chain = "main"
blockcy.Config.Token = "your-api-token-here"
```

## Usage

Check the "types.go" file for information on the return types. Almost all API calls are supported, with a few dropped to reduce the number of distinct types.

For more information on the API, check out [BlockCypher's documentation](http://dev.blockcypher.com/). I've heavily commented the code following Golang convention, so you might also find [the GoDoc quite useful.](http://godoc.org/github.com/acityinohio/blockcy) The "blockcy_test.go" file also shows most of the API calls in action.

## Testing

The aforementioned "blockcy_test.go" file contains a number of tests to ensure the wrapper is functioning properly. If you run it yourself, you'll have to insert a valid API token---additionally, you may want to generate a new token, as the test Posts and Deletes WebHooks and Payment Forwarding requests.

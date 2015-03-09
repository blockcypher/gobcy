//Pacakge blockcy implements a wrapper for the http://blockcypher.com API.
//You can use it to query addresses, transactions, and blocks from the bitcoin
//main and test3 chains, and the BlockCypher test chain (other blockchains currently
//not supported.)
package blockcy

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/url"
)

const (
	BlockURL = "https://api.blockcypher.com/v1/"
)

//Config stores your BlockCypher Token, and the coin/chain
//you're querying. This wrapper currently only supports the "btc" and "bcy"
//coins and their respective chains ("main", "test3", and "test").
//Update them in your code like so:
//	blockcy.Config.Token = "your-api-token"
//	blockcy.Config.Coin = "btc"
//	blockcy.Config.Chain = "main"
var Config struct {
	Token, Coin, Chain string
}

//getResponse is a boilerplate for HTTP GET responses.
func getResponse(target *url.URL) (resp *http.Response, err error) {
	resp, err = http.Get(target.String())
	if err != nil {
		return
	}
	if resp.StatusCode != http.StatusOK {
		msg := make(map[string]string)
		dec := json.NewDecoder(resp.Body)
		dec.Decode(&msg)
		resp.Body.Close()
		err = errors.New(resp.Status + ", Message: " + msg["error"])
	}
	return
}

//postResponse is a boilerplate for HTTP POST responses.
func postResponse(target *url.URL, data io.Reader) (resp *http.Response, err error) {
	resp, err = http.Post(target.String(), "application/json", data)
	if err != nil {
		return
	}
	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusCreated {
		msg := make(map[string]string)
		dec := json.NewDecoder(resp.Body)
		dec.Decode(&msg)
		resp.Body.Close()
		err = errors.New(resp.Status + ", Message: " + msg["error"])
	}
	return
}

//deleteResponse is a boilerplate for HTTP DELETE responses.
func deleteResponse(target *url.URL) (resp *http.Response, err error) {
	req, err := http.NewRequest("DELETE", target.String(), nil)
	if err != nil {
		return
	}
	resp, err = http.DefaultClient.Do(req)
	if err != nil {
		return
	}
	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusNoContent {
		msg := make(map[string]string)
		dec := json.NewDecoder(resp.Body)
		dec.Decode(&msg)
		resp.Body.Close()
		err = errors.New(resp.Status + ", Message: " + msg["error"])
	}
	return
}

//constructs BlockCypher URLs for requests
func buildURL(u string) (target *url.URL, err error) {
	target, err = url.Parse(BlockURL + Config.Coin +
		"/" + Config.Chain + u)
	if err != nil {
		return
	}
	//add token to url, if present
	if Config.Token != "" {
		values := target.Query()
		values.Set("token", Config.Token)
		target.RawQuery = values.Encode()
	}
	return
}

//constructs BlockCypher URLs with parameters for requests
func buildURLParams(u string, params map[string]string) (target *url.URL, err error) {
	target, err = url.Parse(BlockURL + Config.Coin +
		"/" + Config.Chain + u)
	if err != nil {
		return
	}
	values := target.Query()
	//Set parameters
	for k, v := range params {
		values.Set(k, v)
	}
	//add token to url, if present
	if Config.Token != "" {
		values.Set("token", Config.Token)
	}
	target.RawQuery = values.Encode()
	return
}

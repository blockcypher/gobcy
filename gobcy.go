//Package gobcy implements a wrapper for the http://www.blockcypher.com API.
//You can use it to interact with addresses, transactions, and blocks from
//various blockchains, including Bitcoin's main and test3 chains,
//and the BlockCypher test chain.
//
//Please note: we assume you use are using a 64-bit architecture for deployment,
//which automatically makes `int` types 64-bit. Without 64-bit ints, some values
//might overflow on certain calls, depending on the blockchain you are querying.
//If you are using a 32-bit system, you can change all `int` types to `int64` to
//explicitly work around this issue.
package gobcy

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/url"
	"strconv"
)

const baseURL = "https://api.blockcypher.com/v1/"

//API stores your BlockCypher Token, and the coin/chain
//you're querying. Coins can be "btc","bcy","ltc", and "doge".
//Chains can be "main", "test3", or "test", depending on the Coin.
//Check http://dev.blockcypher.com/ for more information.
//All your credentials are stored within an API struct, as are
//many of the API methods.
//You can allocate an API struct like so:
//	bc = gobcy.API{"your-api-token","btc","main"}
//Then query as you like:
//	chain = bc.GetChain()
type API struct {
	Token, Coin, Chain string
}

//getResponse is a boilerplate for HTTP GET responses.
func getResponse(target *url.URL, decTarget interface{}) (err error) {
	resp, err := http.Get(target.String())
	if err != nil {
		return
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		err = respErrorMaker(resp.StatusCode, resp.Body)
		return
	}
	dec := json.NewDecoder(resp.Body)
	err = dec.Decode(decTarget)
	return
}

//postResponse is a boilerplate for HTTP POST responses.
func postResponse(target *url.URL, encTarget interface{}, decTarget interface{}) (err error) {
	var data bytes.Buffer
	enc := json.NewEncoder(&data)
	if err = enc.Encode(encTarget); err != nil {
		return
	}
	resp, err := http.Post(target.String(), "application/json", &data)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusCreated {
		err = respErrorMaker(resp.StatusCode, resp.Body)
		return
	}
	dec := json.NewDecoder(resp.Body)
	err = dec.Decode(decTarget)
	return
}

//putResponse is a boilerplate for HTTP PUT responses.
func putResponse(target *url.URL, encTarget interface{}) (err error) {
	var data bytes.Buffer
	enc := json.NewEncoder(&data)
	if err = enc.Encode(encTarget); err != nil {
		return
	}
	req, err := http.NewRequest("PUT", target.String(), &data)
	if err != nil {
		return
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusNoContent {
		err = respErrorMaker(resp.StatusCode, resp.Body)
	}
	return
}

//deleteResponse is a boilerplate for HTTP DELETE responses.
func deleteResponse(target *url.URL) (err error) {
	req, err := http.NewRequest("DELETE", target.String(), nil)
	if err != nil {
		return
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusNoContent {
		err = respErrorMaker(resp.StatusCode, resp.Body)
	}
	return
}

//respErrorMaker checks error messages/if they are multiple errors
//serializes them into a single error message
func respErrorMaker(statusCode int, body io.Reader) (err error) {
	status := "HTTP " + strconv.Itoa(statusCode) + " " + http.StatusText(statusCode)
	if statusCode == 429 {
		err = errors.New(status)
		return
	}
	type errorJSON struct {
		Err    string `json:"error"`
		Errors []struct {
			Err string `json:"error"`
		} `json:"errors"`
	}
	var msg errorJSON
	dec := json.NewDecoder(body)
	err = dec.Decode(&msg)
	if err != nil {
		return err
	}
	var errtxt string
	errtxt += msg.Err
	for i, v := range msg.Errors {
		if i == len(msg.Errors)-1 {
			errtxt += v.Err
		} else {
			errtxt += v.Err + ", "
		}
	}
	if errtxt == "" {
		err = errors.New(status)
	} else {
		err = errors.New(status + ", Message(s): " + errtxt)
	}
	return
}

//constructs BlockCypher URLs with parameters for requests
func (api *API) buildURL(u string, params map[string]string) (target *url.URL, err error) {
	target, err = url.Parse(baseURL + api.Coin + "/" + api.Chain + u)
	if err != nil {
		return
	}
	values := target.Query()
	//Set parameters
	for k, v := range params {
		values.Set(k, v)
	}
	//add token to url, if present
	if api.Token != "" {
		values.Set("token", api.Token)
	}
	target.RawQuery = values.Encode()
	return
}

//checks token usage
func (api *API) CheckUsage() (usage TokenUsage, err error) {
	u, err := url.Parse(baseURL + "tokens/" + api.Token)
	if err != nil {
		return
	}
	err = getResponse(u, &usage)
	return
}

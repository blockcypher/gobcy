package blockcy

import (
	"encoding/json"
	"errors"
	"net/url"
	"strconv"
)

//GetChain returns the current state of the
//configured Coin/Chain.
func GetChain() (chain Chain, err error) {
	u, err := buildURL("")
	resp, err := getResponse(u)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	//decode JSON into Chain
	dec := json.NewDecoder(resp.Body)
	err = dec.Decode(&chain)
	return
}

//GetBlock returns a Block based on either height
//or hash. If both height and hash are sent, it will
//throw an error.
func GetBlock(height int, hash string) (block Block, err error) {
	block, err = GetBlockPage(height, hash, 0, 0)
	return
}

//GetBlockPage returns a Block based on either height
//or hash, and includes custom variables for txstart/limit of txs.
//If both height and hash are sent, it will throw an error. If txstart/limit = 0,
//it will use the API-defaults for both.
func GetBlockPage(height int, hash string, txstart int, limit int) (block Block, err error) {
	var u *url.URL
	ustr := "/blocks/"
	if height != 0 && hash != "" {
		err = errors.New("Func GetBlock: Cannot send both height and hash")
		return
	} else if height != 0 {
		ustr = ustr + strconv.Itoa(height)
	} else if hash != "" {
		ustr = ustr + hash
	}
	if txstart == 0 && limit == 0 {
		u, err = buildURL(ustr)
	} else {
		params := map[string]string{
			"txstart": strconv.Itoa(txstart),
			"limit":   strconv.Itoa(limit),
		}
		u, err = buildURLParams(ustr, params)
	}
	resp, err := getResponse(u)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	//decode JSON into Block
	dec := json.NewDecoder(resp.Body)
	err = dec.Decode(&block)
	return
}

//GetUnTX returns an array of the latest unconfirmed TXs.
func GetUnTX() (txs []TX, err error) {
	u, err := buildURL("/txs")
	resp, err := getResponse(u)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	//decode JSON into []TX
	dec := json.NewDecoder(resp.Body)
	err = dec.Decode(&txs)
	return
}

//GetTX returns a TX represented by the passed hash.
func GetTX(hash string) (tx TX, err error) {
	u, err := buildURL("/txs/" + hash)
	resp, err := getResponse(u)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	//decode JSON into TX
	dec := json.NewDecoder(resp.Body)
	err = dec.Decode(&tx)
	return
}

//GetTXConf returns a float [0,1] representing BlockCypher's
//confidence that an unconfirmed transaction will be confirmed
//in the next block. Returns an error if the transaction has
//already been confirmed.
func GetTXConf(hash string) (conf float64, err error) {
	u, err := buildURL("/txs/" + hash + "/confidence")
	resp, err := getResponse(u)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	//decode JSON into map[string]interface{} then float
	result := make(map[string]interface{})
	dec := json.NewDecoder(resp.Body)
	err = dec.Decode(&result)
	conf = result["confidence"].(float64)
	return
}

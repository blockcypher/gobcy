package blockcy

import (
	"encoding/json"
	"errors"
	"net/url"
	"strconv"
)

//GetChain returns the current state of the
//configured Coin/Chain.
func (api *API) GetChain() (chain Blockchain, err error) {
	u, err := api.buildURL("")
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
func (api *API) GetBlock(height int, hash string) (block Block, err error) {
	block, err = api.GetBlockPage(height, hash, 0, 0)
	return
}

//GetBlockNextTXs returns the the next page of TXids based
//on the NextTXs URL in this Block. If NextTXs is empty,
//this will return an error.
func (api *API) GetBlockNextTXs(this Block) (next Block, err error) {
	if this.NextTXs == "" {
		err = errors.New("Func GetNextTXs: This Block doesn't have more transactions")
		return
	}
	if len(this.TXids) == 0 {
		err = errors.New("Func GetNextTXs: This Block doesn't have any TXids in the array, meaning no more transactions")
		return
	}
	txurl, err := url.Parse(this.NextTXs)
	if err != nil {
		return
	}
	params := txurl.Query()
	txstart, err := strconv.Atoi(params.Get("txstart"))
	limit, err := strconv.Atoi(params.Get("limit"))
	if err != nil {
		return
	}
	next, err = api.GetBlockPage(0, this.Hash, txstart, limit)
	return
}

//GetBlockPage returns a Block based on either height
//or hash, and includes custom variables for txstart/limit of txs.
//If both height and hash are sent, it will throw an error. If txstart/limit = 0,
//it will use the API-defaults for both.
func (api *API) GetBlockPage(height int, hash string, txstart int, limit int) (block Block, err error) {
	var u *url.URL
	ustr := "/blocks/"
	if height != 0 && hash != "" {
		err = errors.New("Func GetBlockPage: Cannot send both height and hash")
		return
	} else if height != 0 {
		ustr = ustr + strconv.Itoa(height)
	} else if hash != "" {
		ustr = ustr + hash
	}
	if txstart == 0 && limit == 0 {
		u, err = api.buildURL(ustr)
	} else {
		params := map[string]string{
			"txstart": strconv.Itoa(txstart),
			"limit":   strconv.Itoa(limit),
		}
		u, err = api.buildURLParams(ustr, params)
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

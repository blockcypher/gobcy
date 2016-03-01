package gobcy

import (
	"errors"
	"net/url"
	"strconv"
)

//GetChain returns the current state of the
//configured Coin/Chain.
func (api *API) GetChain() (chain Blockchain, err error) {
	u, err := api.buildURL("", nil)
	if err != nil {
		return
	}
	err = getResponse(u, &chain)
	return
}

//GetBlock returns a Block based on either height
//or hash. If both height and hash are sent, it will
//throw an error.
func (api *API) GetBlock(height int, hash string, params map[string]string) (block Block, err error) {
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
	u, err = api.buildURL(ustr, params)
	if err != nil {
		return
	}
	err = getResponse(u, &block)
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
	params := make(map[string]string)
	query := txurl.Query()
	for k := range query {
		params[k] = query.Get(k)
	}
	next, err = api.GetBlock(0, this.Hash, params)
	return
}

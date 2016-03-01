package gobcy

//CreateHDWallet creates a public-address watching HDWallet
//associated with this token/coin/chain, usable anywhere
//in the API where an Address might be used (just use
//the wallet name instead). For example, with checking
//a wallet name balance:
//  addr, err := api.GetAddrBal("your-hd-wallet-name")
func (api *API) CreateHDWallet(req HDWallet) (wal HDWallet, err error) {
	u, err := api.buildURL("/wallets/hd", nil)
	if err != nil {
		return
	}
	err = postResponse(u, &req, &wal)
	return
}

//ListHDWallets lists all known HDWallets associated with
//this token/coin/chain.
//Currently not supported! Use ListWallets() instead.
/*func (api *API) ListHDWallets() (names []string, err error) {
	u, err := api.buildURL("/wallets/hd", nil)
	resp, err := getResponse(u)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	jsonResp := new(struct {
		List []string `json:"hd_wallet_names"`
	})
	dec := json.NewDecoder(resp.Body)
	err = dec.Decode(jsonResp)
	names = jsonResp.List
	return
}*/

//GetHDWallet gets a HDWallet based on its name
//and the associated API token/coin/chain.
func (api *API) GetHDWallet(name string) (wal HDWallet, err error) {
	u, err := api.buildURL("/wallets/hd/"+name, nil)
	if err != nil {
		return
	}
	err = getResponse(u, &wal)
	return
}

//GetAddrHDWallet returns addresses associated with
//a named HDWallet, associated with the API token/coin/chain.
//It also optionally accepts URL parameters.
func (api *API) GetAddrHDWallet(name string, params map[string]string) (addrs HDWallet, err error) {
	u, err := api.buildURL("/wallets/hd/"+name+"/addresses", params)
	if err != nil {
		return
	}
	err = getResponse(u, &addrs)
	return
}

//DeriveAddrHDWallet derives a new address within the named Wallet,
//associated with the API token/coin/chain. It will only return a partial
//HDWallet, ONLY containing the new address derived.
func (api *API) DeriveAddrHDWallet(name string, params map[string]string) (wal HDWallet, err error) {
	u, err := api.buildURL("/wallets/hd/"+name+"/addresses/derive", params)
	if err != nil {
		return
	}
	err = postResponse(u, nil, &wal)
	return
}

//DeleteHDWallet deletes a named HDWallet associated with the
//API token/coin/chain.
func (api *API) DeleteHDWallet(name string) (err error) {
	u, err := api.buildURL("/wallets/hd/"+name, nil)
	if err != nil {
		return
	}
	err = deleteResponse(u)
	return
}

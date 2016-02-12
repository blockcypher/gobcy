package gobcy

import (
	"bytes"
	"encoding/json"
	"strconv"
	"strings"
)

//CreateWallet creates a public-address watching wallet
//associated with this token/coin/chain, usable anywhere
//in the API where an Address might be used (just use
//the wallet name instead). For example, with checking
//a wallet name balance:
//  addr, err := api.GetAddrBal("your-wallet-name", nil)
func (api *API) CreateWallet(req Wallet) (wal Wallet, err error) {
	u, err := api.buildURL("/wallets", nil)
	if err != nil {
		return
	}
	var data bytes.Buffer
	enc := json.NewEncoder(&data)
	if err = enc.Encode(&req); err != nil {
		return
	}
	resp, err := postResponse(u, &data)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	dec := json.NewDecoder(resp.Body)
	err = dec.Decode(&wal)
	return
}

//ListWallets lists all known Wallets associated with
//this token/coin/chain.
func (api *API) ListWallets() (names []string, err error) {
	u, err := api.buildURL("/wallets", nil)
	resp, err := getResponse(u)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	jsonResp := new(struct {
		List []string `json:"wallet_names"`
	})
	dec := json.NewDecoder(resp.Body)
	err = dec.Decode(jsonResp)
	names = jsonResp.List
	return
}

//GetWallet gets a Wallet based on its name, the associated
//API token/coin/chain, and whether it's an HD wallet or
//not.
func (api *API) GetWallet(name string) (wal Wallet, err error) {
	u, err := api.buildURL("/wallets/"+name, nil)
	resp, err := getResponse(u)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	dec := json.NewDecoder(resp.Body)
	err = dec.Decode(&wal)
	return
}

//AddAddrWallet adds a slice of addresses to a named Wallet,
//associated with the API token/coin/chain. In addition to your
//list of addresses to add, takes one additional parameter:
//  "omitAddr," if true will omit wallet addresses in your
//  response. Useful to speed up the API call for larger wallets.
func (api *API) AddAddrWallet(name string, addrs []string, omitAddr bool) (wal Wallet, err error) {
	params := map[string]string{"omitWalletAddresses": strconv.FormatBool(omitAddr)}
	u, err := api.buildURL("/wallets/"+name+"/addresses", params)
	if err != nil {
		return
	}
	var data bytes.Buffer
	enc := json.NewEncoder(&data)
	if err = enc.Encode(&Wallet{Addresses: addrs}); err != nil {
		return
	}
	resp, err := postResponse(u, &data)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	dec := json.NewDecoder(resp.Body)
	err = dec.Decode(&wal)
	return
}

//GetAddrWallet returns a slice of addresses associated with
//a named Wallet, associated with the API token/coin/chain.
//Takes an optionally-nil URL parameter map.
func (api *API) GetAddrWallet(name string, params map[string]string) (addrs []string, err error) {
	u, err := api.buildURL("/wallets/"+name+"/addresses", params)
	resp, err := getResponse(u)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	var wal Wallet
	dec := json.NewDecoder(resp.Body)
	err = dec.Decode(&wal)
	addrs = wal.Addresses
	return
}

//DeleteAddrWallet deletes a slice of addresses associated with
//a named Wallet, associated with the API token/coin/chain.
func (api *API) DeleteAddrWallet(name string, addrs []string) (err error) {
	u, err := api.buildURL("/wallets/"+name+"/addresses",
		map[string]string{"address": strings.Join(addrs, ";")})
	resp, err := deleteResponse(u)
	if err != nil {
		return
	}
	resp.Body.Close()
	return
}

//GenAddrWallet generates a new address within the named Wallet,
//associated with the API token/coin/chain. Also returns the
//private/WIF/public key of address via an Address Keychain.
func (api *API) GenAddrWallet(name string) (wal Wallet, addr AddrKeychain, err error) {
	u, err := api.buildURL("/wallets/"+name+"/addresses/generate", nil)
	resp, err := postResponse(u, nil)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	dec := json.NewDecoder(resp.Body)
	//weird anonymous struct composition FTW
	err = dec.Decode(&struct {
		*Wallet
		*AddrKeychain
	}{&wal, &addr})
	return
}

//DeleteWallet deletes a named wallet associated with the
//API token/coin/chain.
func (api *API) DeleteWallet(name string) (err error) {
	u, err := api.buildURL("/wallets/"+name, nil)
	resp, err := deleteResponse(u)
	if err != nil {
		return
	}
	resp.Body.Close()
	return
}

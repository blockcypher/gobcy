package blockcy

import (
	"bytes"
	"encoding/json"
	"net/url"
	"strings"
)

//CreateWallet creates a public-address watching wallet
//associated with this token/coin/chain, usable anywhere
//in the API where an Address might be used (just use
//the wallet name instead). For example, with checking
//a wallet name balance:
//  addr, err := api.GetAddrBal("your-wallet-name")
//Wallet is either a normal list of addresses, or an HD
//wallet derived via an extended public seed and subchains.
//Make sure your Wallet has "HD" set to "true" if you're
//requesting to construct an HD wallet.
func (api *API) CreateWallet(req Wallet) (wal Wallet, err error) {
	//Decide whether to create HD or regular Wallet
	var u *url.URL
	if req.HD {
		u, err = api.buildURL("/wallets/hd")
	} else {
		u, err = api.buildURL("/wallets")
	}
	if err != nil {
		return
	}
	//encode post data into ReadWriter
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
	//Decode JSON into result
	dec := json.NewDecoder(resp.Body)
	err = dec.Decode(&wal)
	return
}

//GetWallet gets a Wallet based on its name, the associated
//API token/coin/chain, and whether it's an HD wallet or
//not.
func (api *API) GetWallet(name string, hd bool) (wal Wallet, err error) {
	//Decide whether to get HD or regular Wallet
	var u *url.URL
	if hd {
		u, err = api.buildURL("/wallets/hd/" + name)
	} else {
		u, err = api.buildURL("/wallets/" + name)
	}
	resp, err := getResponse(u)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	//decode JSON into result
	dec := json.NewDecoder(resp.Body)
	err = dec.Decode(&wal)
	return
}

//AddAddrWallet adds a slice of addresses to a named Wallet,
//associated with the API token/coin/chain. Only works with
//normal (non-HD) wallets; trying to add Addresses to an HD
//wallet will return an error.
func (api *API) AddAddrWallet(name string, addrs []string) (wal Wallet, err error) {
	u, err := api.buildURL("/wallets/" + name + "/addresses")
	if err != nil {
		return
	}
	//encode post data into ReadWriter
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
	//Decode JSON into result
	dec := json.NewDecoder(resp.Body)
	err = dec.Decode(&wal)
	return
}

//GetAddrWallet returns a slice of addresses associated with
//a named Wallet, associated with the API token/coin/chain. Must
//denote whether a Wallet is HD or not.
func (api *API) GetAddrWallet(name string, hd bool) (addrs []string, err error) {
	//Decide whether to query HD or regular Wallet
	var u *url.URL
	if hd {
		u, err = api.buildURL("/wallets/hd/" + name + "/addresses")
	} else {
		u, err = api.buildURL("/wallets/" + name + "/addresses")
	}
	if err != nil {
		return
	}
	resp, err := getResponse(u)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	//decode JSON into result
	var wal Wallet
	dec := json.NewDecoder(resp.Body)
	err = dec.Decode(&wal)
	addrs = wal.Addresses
	return
}

//DeleteAddrWallet deletes a slice of addresses associated with
//a named Wallet, associated with the API token/coin/chain. Does
//not work with HD wallets, as their addresses cannot be deleted.
func (api *API) DeleteAddrWallet(name string, addrs []string) (err error) {
	u, err := api.buildURLParams("/wallets/"+name+"/addresses",
		map[string]string{"addresses": strings.Join(addrs, ";")})
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
//Only works on normal wallets; for HD wallets, use DeriveAddrWallet.
func (api *API) GenAddrWallet(name string) (wal Wallet, addr AddrKeychain, err error) {
	u, err := api.buildURL("/wallets/" + name + "/addresses/generate")
	resp, err := postResponse(u, nil)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	//decode JSON into result
	dec := json.NewDecoder(resp.Body)
	//weird anonymous struct composition FTW
	err = dec.Decode(&struct {
		*Wallet
		*AddrKeychain
	}{&wal, &addr})
	return
}

//DeleteWallet deletes a named wallet associated with the
//API token/coin/chain. Must specify if it's an HD wallet.
func (api *API) DeleteWallet(name string, hd bool) (err error) {
	//Decide whether to delete HD or regular Wallet
	var u *url.URL
	if hd {
		u, err = api.buildURL("/wallets/hd/" + name)
	} else {
		u, err = api.buildURL("/wallets/" + name)
	}
	if err != nil {
		return
	}
	resp, err := deleteResponse(u)
	if err != nil {
		return
	}
	resp.Body.Close()
	return
}

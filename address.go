package gobcy

import (
	"errors"
	"strconv"
)

//GetAddrBal returns balance information for a given public
//address. Fastest Address API call, but does not
//include transaction details.
func (api *API) GetAddrBal(hash string, params map[string]string) (addr Addr, err error) {
	u, err := api.buildURL("/addrs/"+hash+"/balance", params)
	if err != nil {
		return
	}
	err = getResponse(u, &addr)
	return
}

//GetAddr returns information for a given public
//address, including a slice of confirmed and unconfirmed
//transaction outpus via the TXRef arrays in the Address
//type. Returns more information than GetAddrBal, but
//slightly slower.
func (api *API) GetAddr(hash string, params map[string]string) (addr Addr, err error) {
	u, err := api.buildURL("/addrs/"+hash, params)
	if err != nil {
		return
	}
	err = getResponse(u, &addr)
	return
}

//GetAddrNext returns a given Addr's next page of TXRefs,
//if Addr.HasMore is true. If HasMore is false, will
//return an error. It assumes default API URL parameters.
func (api *API) GetAddrNext(this Addr) (next Addr, err error) {
	if !this.HasMore {
		err = errors.New("Func GetAddrNext: this Addr doesn't have more TXRefs according to its HasMore")
		return
	}
	before := this.TXRefs[len(this.TXRefs)-1].BlockHeight
	next, err = api.GetAddr(this.Address, map[string]string{"before": strconv.Itoa(before)})
	return
}

//GetAddrFull returns information for a given public
//address, including a slice of TXs associated
//with this address. Returns more data than GetAddr since
//it includes full transactions, but slowest Address query.
func (api *API) GetAddrFull(hash string, params map[string]string) (addr Addr, err error) {
	u, err := api.buildURL("/addrs/"+hash+"/full", params)
	if err != nil {
		return
	}
	err = getResponse(u, &addr)
	return
}

//GetAddrFullNext returns a given Addr's next page of TXs,
//if Addr.HasMore is true. If HasMore is false, will
//return an error. It assumes default API URL parameters, like GetAddrFull.
func (api *API) GetAddrFullNext(this Addr) (next Addr, err error) {
	if !this.HasMore {
		err = errors.New("Func GetAddrFullNext: this Addr doesn't have more TXs according to its HasMore")
		return
	}
	before := this.TXs[len(this.TXs)-1].BlockHeight
	next, err = api.GetAddrFull(this.Address, map[string]string{"before": strconv.Itoa(before)})
	return
}

//GenAddrKeychain generates a public/private key pair for use with
//transactions within the specified coin/chain. Please note that
//this call must be made over SSL, and it is not recommended to keep
//large amounts in these addresses, or for very long.
func (api *API) GenAddrKeychain() (pair AddrKeychain, err error) {
	u, err := api.buildURL("/addrs", nil)
	if err != nil {
		return
	}
	err = postResponse(u, nil, &pair)
	return
}

//GenAddrMultisig generates a P2SH multisignature address using an array
//of PubKeys and the ScriptType from a AddrKeychain. Other fields are
//ignored, and the ScriptType must be a "multisig-n-of-m" type. Returns
//an AddrKeychain with the same PubKeys, ScriptType, and the proper
//P2SH address in the AddrKeychain's address field.
func (api *API) GenAddrMultisig(multi AddrKeychain) (addr AddrKeychain, err error) {
	if len(multi.PubKeys) == 0 || multi.ScriptType == "" {
		err = errors.New("GenAddrMultisig: PubKeys or ScriptType are empty.")
		return
	}
	u, err := api.buildURL("/addrs", nil)
	if err != nil {
		return
	}
	err = postResponse(u, &multi, &addr)
	return
}

//Faucet funds the AddrKeychain with an amount. Only works on BlockCypher's
//Testnet and Bitcoin Testnet3. Returns the transaction hash funding
//your AddrKeychain.
func (api *API) Faucet(a AddrKeychain, amount int) (txhash string, err error) {
	if !(api.Coin == "bcy" && api.Chain == "test") && !(api.Coin == "btc" && api.Chain == "test3") {
		err = errors.New("Faucet: Cannot use Faucet unless on BlockCypher Testnet or Bitcoin Testnet3.")
		return
	}
	u, err := api.buildURL("/faucet", nil)
	if err != nil {
		return
	}
	type FauxAddr struct {
		Address string `json:"address"`
		Amount  int    `json:"amount"`
	}
	var addr string
	//for easy funding/testing of OAPAddresses
	if a.OriginalAddress != "" {
		addr = a.OriginalAddress
	} else {
		addr = a.Address
	}
	txref := make(map[string]string)
	err = postResponse(u, &FauxAddr{addr, amount}, &txref)
	txhash = txref["tx_ref"]
	return
}

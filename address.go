package gobcy

import (
	"bytes"
	"encoding/json"
	"errors"
	"strconv"
)

//GetAddrBal returns balance information for a given public
//address. Fastest Address API call, but does not
//include transaction details.
func (api *API) GetAddrBal(hash string) (addr Addr, err error) {
	addr, err = api.GetAddrBalCustom(hash, false)
	return
}

//GetAddrBalCustom returns balance information for a given public
//address. Fastest Address API call, but does not
//include transaction details. Takes one additional parameter
//compared to GetAddrBal:
//  "omitWalletAddr," if true will omit wallet addresses if
//  you're querying a wallet instead of an address. Useful to
//  speed up the API call for larger wallets.
func (api *API) GetAddrBalCustom(hash string, omitWalletAddr bool) (addr Addr, err error) {
	params := map[string]string{"omitWalletAddresses": strconv.FormatBool(omitWalletAddr)}
	u, err := api.buildURLParams("/addrs/"+hash+"/balance", params)
	resp, err := getResponse(u)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	//decode JSON into Addr
	dec := json.NewDecoder(resp.Body)
	err = dec.Decode(&addr)
	return
}

//GetAddr returns information for a given public
//address, including a slice of confirmed and unconfirmed
//transaction outpus via the TXRef arrays in the Address
//type. Returns more information than GetAddrBal, but
//slightly slower.
func (api *API) GetAddr(hash string) (addr Addr, err error) {
	addr, err = api.GetAddrCustom(hash, false, 0, 0, 0, 0, false, false)
	return
}

//GetAddrNext returns a given Addr's next page of TXRefs,
//if Addr.HasMore is true. If HasMore is false, will
//return an error. It assumes default API flags, like GetAddr.
func (api *API) GetAddrNext(this Addr) (next Addr, err error) {
	if !this.HasMore {
		err = errors.New("Func GetAddrNext: this Addr doesn't have more TXRefs according to its HasMore")
		return
	}
	before := this.TXRefs[len(this.TXRefs)-1].BlockHeight
	next, err = api.GetAddrCustom(this.Address, false, 0, before, 0, 0, false, false)
	return
}

//GetAddrCustom returns information for a given public
//address, including a slice of confirmed and unconfirmed
//transaction outpus via the TXRef arrays in the Address
//type. Takes 5 additional parameters compared to GetAddr:
//  "unspent," which if true will only return TXRefs
//  that are unpsent outputs (UTXOs).
//  "confirms," which will only return TXRefs
//  that have reached this number of confirmations or more.
//  Set it to 0 to ignore this parameter.
//  "before," which will only return transactions below
//  this height in the blockchain. Useful for paging. Set it
//  to 0 to ignore this parameter.
//	"after," which will only return transaction above
//	this height in the blockchain. Useful for paging. Set it
//	to 0 to ignore this parameter.
//  "limit," which return this number of TXRefs per call.
//  The default is 50, maximum is 200. Set it to 0 to ignore
//  this parameter and use the API-set default.
//  "omitWalletAddr," if true will omit wallet addresses if
//  you're querying a wallet instead of an address. Useful to
//  speed up the API call for larger wallets.
//	"includeConfidence," if true, includes confidence information
//	for unconfirmed transactions.
func (api *API) GetAddrCustom(hash string, unspent bool, confirms int, before int,
	after int, limit int, omitWalletAddr bool, includeConfidence bool) (addr Addr, err error) {
	params := map[string]string{"unspentOnly": strconv.FormatBool(unspent), "omitWalletAddresses": strconv.FormatBool(omitWalletAddr), "includeConfidence": strconv.FormatBool(includeConfidence)}
	if confirms > 0 {
		params["confirmations"] = strconv.Itoa(confirms)
	}
	if before > 0 {
		params["before"] = strconv.Itoa(before)
	}
	if after > 0 {
		params["after"] = strconv.Itoa(after)
	}
	if limit > 0 {
		params["limit"] = strconv.Itoa(limit)
	}
	u, err := api.buildURLParams("/addrs/"+hash, params)
	resp, err := getResponse(u)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	//decode JSON into Addr
	dec := json.NewDecoder(resp.Body)
	err = dec.Decode(&addr)
	return
}

//GetAddrFull returns information for a given public
//address, including a slice of TXs associated
//with this address. Returns more data than GetAddr since
//it includes full transactions, but slowest Address query.
func (api *API) GetAddrFull(hash string) (addr Addr, err error) {
	addr, err = api.GetAddrFullCustom(hash, false, 0, 0, 0, 0, false, false)
	return
}

//GetAddrFullNext returns a given Addr's next page of TXs,
//if Addr.HasMore is true. If HasMore is false, will
//return an error. It assumes default API flags, like GetAddrFull.
func (api *API) GetAddrFullNext(this Addr) (next Addr, err error) {
	if !this.HasMore {
		err = errors.New("Func GetAddrFullNext: this Addr doesn't have more TXs according to its HasMore")
		return
	}
	before := this.TXs[len(this.TXs)-1].BlockHeight
	next, err = api.GetAddrFullCustom(this.Address, false, 0, 0, before, 0, false, false)
	return
}

//GetAddrFullCustom returns information for a given public
//address, including a slice of TXs associated
//with this address. Returns more data than GetAddr since
//it includes full transactions, but slower. Takes 4
//additional parameters compared to GetAddrFull:
//  "hex," which if true will return the full hex-encoded
//  raw transaction for each TX. False by default.
//  "confirms," which will only return TXRefs
//  that have reached this number of confirmations or more.
//  Set it to 0 to ignore this parameter.
//  "before," which will only return transactions below
//  this height in the blockchain. Useful for paging. Set it
//  to 0 to ignore this parameter.
//	"after," which will only return transaction above
//	this height in the blockchain. Useful for paging. Set it
//	to 0 to ignore this parameter.
//  "limit," which return this number of TXs per call.
//  The default is 10, maximum is 50. Set it to 0 to ignore
//  this parameter and use the API-set default.
//  "omitWalletAddr," if true will omit wallet addresses if
//  you're querying a wallet instead of an address. Useful to
//  speed up the API call for larger wallets.
//	"includeConfidence," if true, includes confidence information
//	for unconfirmed transactions.
func (api *API) GetAddrFullCustom(hash string, hex bool, confirms int, before int, after int,
	limit int, omitWalletAddr bool, includeConfidence bool) (addr Addr, err error) {
	params := map[string]string{"includeHex": strconv.FormatBool(hex), "omitWalletAddresses": strconv.FormatBool(omitWalletAddr), "includeConfidence": strconv.FormatBool(includeConfidence)}
	if confirms > 0 {
		params["confirmations"] = strconv.Itoa(confirms)
	}
	if before > 0 {
		params["before"] = strconv.Itoa(before)
	}
	if after > 0 {
		params["after"] = strconv.Itoa(after)
	}
	if limit > 0 {
		params["limit"] = strconv.Itoa(limit)
	}
	u, err := api.buildURLParams("/addrs/"+hash+"/full", params)
	resp, err := getResponse(u)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	//decode JSON into Addr
	dec := json.NewDecoder(resp.Body)
	err = dec.Decode(&addr)
	return
}

//GenAddrKeychain generates a public/private key pair for use with
//transactions within the specified coin/chain. Please note that
//this call must be made over SSL, and it is not recommended to keep
//large amounts in these addresses, or for very long.
func (api *API) GenAddrKeychain() (pair AddrKeychain, err error) {
	u, err := api.buildURL("/addrs")
	resp, err := postResponse(u, nil)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	//decode JSON into AddrKeychain
	dec := json.NewDecoder(resp.Body)
	err = dec.Decode(&pair)
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
	u, err := api.buildURL("/addrs")
	if err != nil {
		return
	}
	var data bytes.Buffer
	enc := json.NewEncoder(&data)
	if err = enc.Encode(&multi); err != nil {
		return
	}
	resp, err := postResponse(u, &data)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	dec := json.NewDecoder(resp.Body)
	err = dec.Decode(&addr)
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
	if amount > 1e7 {
		err = errors.New("Faucet: Cannot fund with more than 10,000,000 coins at a time.")
		return
	}
	u, err := api.buildURL("/faucet")
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
	var data bytes.Buffer
	enc := json.NewEncoder(&data)
	if err = enc.Encode(&FauxAddr{addr, amount}); err != nil {
		return
	}
	resp, err := postResponse(u, &data)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	//decode JSON into map[string]string
	txref := make(map[string]string)
	dec := json.NewDecoder(resp.Body)
	err = dec.Decode(&txref)
	txhash = txref["tx_ref"]
	return
}

package blockcy

import (
	"bytes"
	"encoding/json"
	"errors"
	"strconv"
)

//GetAddr returns balance information for a given public
//address. Does not include transaction details.
func GetAddr(hash string) (addr Addr, err error) {
	u, err := buildURL("/addrs/" + hash + "/balance")
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
//address, including a slice of transactions associated
//with this address. Takes an additional parameter "unspent."
//If true, unspent will only return transactions with unspent
//outputs (UTXO).
func GetAddrFull(hash string, unspent bool) (addr Addr, err error) {
	params := map[string]string{"unspentOnly": strconv.FormatBool(unspent)}
	u, err := buildURLParams("/addrs/"+hash+"/full", params)
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

//GenAddrPair generates a public/private key pair for use with
//transactions within the specified coin/chain. Please note that
//this call must be made over SSL, and it is not recommended to keep
//large amounts in these addresses, or for very long.
func GenAddrPair() (pair AddrPair, err error) {
	u, err := buildURL("/addrs")
	resp, err := postResponse(u, nil)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	//decode JSON into AddrPair
	dec := json.NewDecoder(resp.Body)
	err = dec.Decode(&pair)
	return
}

//Faucet funds the AddrPair with an amount. Only works on BlockCypher's
//testnet, where Config.Coin = "bcy" and Config.Chain = "test."
func (a *AddrPair) Faucet(amount int) (txhash string, err error) {
	if Config.Coin != "bcy" && Config.Chain != "test" {
		err = errors.New("Faucet: Cannot use Faucet unless on BlockCypher testnet.")
		return
	}
	if amount > 1e7 {
		err = errors.New("Faucet: Cannot fund with more than 10,000,000 coins at a time.")
		return
	}
	u, err := buildURL("/faucet")
	if err != nil {
		return
	}
	type FauxAddr struct {
		Address string `json:"address"`
		Amount  int    `json:"amount"`
	}
	var data bytes.Buffer
	enc := json.NewEncoder(&data)
	if err = enc.Encode(&FauxAddr{a.Address, amount}); err != nil {
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

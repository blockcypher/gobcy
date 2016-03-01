package gobcy

import (
	"encoding/hex"
	"errors"
	"strconv"

	"github.com/btcsuite/btcd/btcec"
)

//GetUnTX returns an array of the latest unconfirmed TXs.
func (api *API) GetUnTX() (txs []TX, err error) {
	u, err := api.buildURL("/txs", nil)
	if err != nil {
		return
	}
	err = getResponse(u, &txs)
	return
}

//GetTX returns a TX represented by the passed hash. Takes
//an optionally-nil URL parameter map.
func (api *API) GetTX(hash string, params map[string]string) (tx TX, err error) {
	u, err := api.buildURL("/txs/"+hash, params)
	if err != nil {
		return
	}
	err = getResponse(u, &tx)
	return
}

//GetTXConf returns a TXConf containing a float [0,1] that
//represents BlockCypher's confidence that an unconfirmed transaction
//won't be successfully double-spent against. If the confidence is 1,
//the transaction has already been confirmed.
func (api *API) GetTXConf(hash string) (conf TXConf, err error) {
	u, err := api.buildURL("/txs/"+hash+"/confidence", nil)
	if err != nil {
		return
	}
	err = getResponse(u, &conf)
	return
}

//TempNewTX creates a simple template transaction, suitable for
//use in NewTX. Takes an input/output address and amount.
func TempNewTX(inAddr string, outAddr string, amount int) (trans TX) {
	trans.Inputs = make([]TXInput, 1)
	trans.Outputs = make([]TXOutput, 1)
	trans.Inputs[0].Addresses = make([]string, 1)
	trans.Outputs[0].Addresses = make([]string, 1)
	trans.Inputs[0].Addresses[0] = inAddr
	trans.Outputs[0].Addresses[0] = outAddr
	trans.Outputs[0].Value = amount
	return
}

//TempMultiTX creates a skeleton multisig transaction,
//suitable for use in NewTX. If outAddr == "", then the
//returned TX will be a skeleton to fund a multisig address.
//If inAddr == "", then the returned TX will be a skeleton to
//send from a multisig address (/series of public keys).
//n represents the number of valid signatures required, and m
//is derived from the number of pubkeys.
func TempMultiTX(inAddr string, outAddr string, amount int, n int, pubkeys []string) (trans TX, err error) {
	m := len(pubkeys)
	if inAddr != "" && outAddr != "" {
		err = errors.New("TempMultiTX: Can't have both inAddr and outAddr != \"\"")
		return
	}
	if n > m {
		err = errors.New("TempMultiTX: n cannot be greater than the number of pubkeys")
		return
	}
	scripttype := "multisig-" + strconv.Itoa(n) + "-of-" + strconv.Itoa(m)
	trans.Inputs = make([]TXInput, 1)
	trans.Outputs = make([]TXOutput, 1)
	if inAddr != "" {
		trans.Inputs[0].Addresses = make([]string, 1)
		trans.Inputs[0].Addresses[0] = inAddr
		trans.Outputs[0].Addresses = pubkeys
		trans.Outputs[0].ScriptType = scripttype
	} else if outAddr != "" {
		trans.Inputs[0].Addresses = pubkeys
		trans.Inputs[0].ScriptType = scripttype
		trans.Outputs[0].Addresses = make([]string, 1)
		trans.Outputs[0].Addresses[0] = outAddr
	}
	trans.Outputs[0].Value = amount
	return
}

//NewTX takes a partially formed TX and returns a TXSkel
//with the data that needs to be signed. Can use TempNewTX
//or TempMultiTX to streamline input transaction, or customize
//transaction as described in the BlockCypher docs:
//http://dev.blockcypher.com/#customizing-transaction-requests
//If verify is true, will include "ToSignTX," which can be used
//to locally verify the "ToSign" data is valid.
func (api *API) NewTX(trans TX, verify bool) (skel TXSkel, err error) {
	u, err := api.buildURL("/txs/new",
		map[string]string{"includeToSignTx": strconv.FormatBool(verify)})
	if err != nil {
		return
	}
	err = postResponse(u, &trans, &skel)
	return
}

//Sign takes a hex-encoded string slice of private
//keys and uses them to sign the ToSign data in a
//TXSkel, generating the proper Signatures and PubKeys
//array, both hex-encoded. This is meant as a helper
//function, and leverages btcd's btcec library.
func (skel *TXSkel) Sign(priv []string) (err error) {
	//num of private keys must match len(ToSign)
	//Often this might mean repeating private keys
	if len(priv) != len(skel.ToSign) {
		err = errors.New("*TXSkel.Sign error: number of private keys != length of ToSign array")
		return
	}
	//Loop through keys, append sigs/public key
	for i, k := range priv {
		privDat, err := hex.DecodeString(k)
		tosign, err := hex.DecodeString(skel.ToSign[i])
		if err != nil {
			return err
		}
		privkey, pubkey := btcec.PrivKeyFromBytes(btcec.S256(), privDat)
		sig, err := privkey.Sign(tosign)
		if err != nil {
			return err
		}
		skel.Signatures = append(skel.Signatures, hex.EncodeToString(sig.Serialize()))
		skel.PubKeys = append(skel.PubKeys, hex.EncodeToString(pubkey.SerializeCompressed()))
	}
	return
}

//SendTX takes a TXSkel, returns the completed
//transaction and sends it across the Coin/Chain
//network. TXSkel requires a fully formed TX, Signatures,
//and PubKeys. PubKeys should not be included in the
//special case of multi-sig addresses.
func (api *API) SendTX(skel TXSkel) (trans TXSkel, err error) {
	u, err := api.buildURL("/txs/send", nil)
	if err != nil {
		return
	}
	err = postResponse(u, &skel, &trans)
	return
}

//PushTX takes a hex-encoded transaction string
//and pushes it directly to the Coin/Chain network.
func (api *API) PushTX(hex string) (trans TXSkel, err error) {
	u, err := api.buildURL("/txs/push", nil)
	if err != nil {
		return
	}
	err = postResponse(u, &map[string]string{"tx": hex}, &trans)
	return
}

//DecodeTX takes a hex-encoded transaction string
//and decodes it into a TX object, without sending
//it along to the Coin/Chain network.
func (api *API) DecodeTX(hex string) (trans TXSkel, err error) {
	u, err := api.buildURL("/txs/decode", nil)
	if err != nil {
		return
	}
	err = postResponse(u, &map[string]string{"tx": hex}, &trans)
	return
}

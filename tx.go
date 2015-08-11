package blockcy

import (
	"bytes"
	"encoding/json"
	"errors"
	"strconv"
)

//GetUnTX returns an array of the latest unconfirmed TXs.
func (self *API) GetUnTX() (txs []TX, err error) {
	u, err := self.buildURL("/txs")
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
func (self *API) GetTX(hash string) (tx TX, err error) {
	u, err := self.buildURL("/txs/" + hash)
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
//in the next block. If it returns a 1, the transaction has
//already been confirmed.
func (self *API) GetTXConf(hash string) (conf float64, err error) {
	u, err := self.buildURL("/txs/" + hash + "/confidence")
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

//SkelTX creates a skeleton transaction,
//suitable for use in NewTX.
func SkelTX(inAddr string, outAddr string, amount int, confirm bool) (trans TX) {
	trans.Inputs = make([]TXInput, 1)
	trans.Outputs = make([]TXOutput, 1)
	trans.Inputs[0].Addresses = make([]string, 1)
	trans.Outputs[0].Addresses = make([]string, 1)
	trans.Inputs[0].Addresses[0] = inAddr
	trans.Outputs[0].Addresses[0] = outAddr
	trans.Outputs[0].Value = amount
	if confirm {
		trans.Confirmations = 1
	}
	return
}

//SkelMultiTX creates a skeleton multisig transaction,
//suitable for use in NewTX. If outAddr == "", then the
//returned TX will be a skeleton to fund a multisig address.
//If inAddr == "", then the returned TX will be a skeleton to
//send from a multisig address (/series of public keys).
//n represents the number of valid signatures required, and m
//is derived from the number of pubkeys.
func SkelMultiTX(inAddr string, outAddr string, amount int, confirm bool, n int, pubkeys []string) (trans TX, err error) {
	m := len(pubkeys)
	if inAddr != "" && outAddr != "" {
		err = errors.New("SkelMultiTX: Can't have both inAddr and outAddr != \"\"")
		return
	}
	if n > m {
		err = errors.New("SkelMultiTX: n cannot be greater than the number of pubkeys")
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
	if confirm {
		trans.Confirmations = 1
	}
	return
}

//NewTX takes a partially formed TX and returns
//a WipTX with the data that needs to be signed.
func (self *API) NewTX(trans TX) (wip TXSkel, err error) {
	u, err := self.buildURL("/txs/new")
	if err != nil {
		return
	}
	//encode response into ReadWriter
	var data bytes.Buffer
	enc := json.NewEncoder(&data)
	if err = enc.Encode(&trans); err != nil {
		return
	}
	resp, err := postResponse(u, &data)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	//Decode JSON into result
	dec := json.NewDecoder(resp.Body)
	err = dec.Decode(&wip)
	return
}

//SendTX takes a WipTX, returns the completed
//transaction and sends it across the Coin/Chain
//network. WipTX requires a fully formed TX, Signatures,
//and PubKeys. PubKeys should not be included in the
//special case of multi-sig addresses.
func (self *API) SendTX(wip TXSkel) (trans TXSkel, err error) {
	u, err := self.buildURL("/txs/send")
	if err != nil {
		return
	}
	//encode response into ReadWriter
	var data bytes.Buffer
	enc := json.NewEncoder(&data)
	if err = enc.Encode(&wip); err != nil {
		return
	}
	resp, err := postResponse(u, &data)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	//Decode JSON into result
	dec := json.NewDecoder(resp.Body)
	err = dec.Decode(&trans)
	return
}

//PushTX takes a hex-encoded transaction string
//and pushes it directly to the Coin/Chain network.
func (self *API) PushTX(hex string) (trans TXSkel, err error) {
	u, err := self.buildURL("/txs/push")
	if err != nil {
		return
	}
	//encode response into ReadWriter
	var data bytes.Buffer
	enc := json.NewEncoder(&data)
	if err = enc.Encode(map[string]string{"tx": hex}); err != nil {
		return
	}
	resp, err := postResponse(u, &data)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	//Decode JSON into result
	dec := json.NewDecoder(resp.Body)
	err = dec.Decode(&trans)
	return
}

//DecodeTX takes a hex-encoded transaction string
//and decodes it into a TX object, without sending
//it along to the Coin/Chain network.
func (self *API) DecodeTX(hex string) (trans TXSkel, err error) {
	u, err := self.buildURL("/txs/decode")
	if err != nil {
		return
	}
	//encode response into ReadWriter
	var data bytes.Buffer
	enc := json.NewEncoder(&data)
	if err = enc.Encode(map[string]string{"tx": hex}); err != nil {
		return
	}
	resp, err := postResponse(u, &data)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	//Decode JSON into result
	dec := json.NewDecoder(resp.Body)
	err = dec.Decode(&trans)
	return
}

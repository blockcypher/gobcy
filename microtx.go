package blockcy

import (
	"bytes"
	"encoding/json"
)

//SendMicro sends a Micro through the Coin/Chain
//network. It will return a Micro with a proper hash
//if it successfully sent. If using public (instead of
//private) keys, you'll need to sign the returned Micro
//and run SendMicro again until you will get a hash.
func (api *API) SendMicro(mic MicroTX) (result MicroTX, err error) {
	u, err := api.buildURL("/txs/micro")
	if err != nil {
		return
	}
	//encode response into ReadWriter
	var data bytes.Buffer
	enc := json.NewEncoder(&data)
	if err = enc.Encode(&mic); err != nil {
		return
	}
	resp, err := postResponse(u, &data)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	//Decode JSON into result
	dec := json.NewDecoder(resp.Body)
	err = dec.Decode(&result)
	return
}

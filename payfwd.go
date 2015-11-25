package gobcy

import (
	"bytes"
	"encoding/json"
)

//CreatePayFwd creates a new PayFwd forwarding
//request associated with your API.Token, and
//returns a PayFwd with a BlockCypher-assigned id.
func (api *API) CreatePayFwd(payment PayFwd) (result PayFwd, err error) {
	u, err := api.buildURL("/payments")
	if err != nil {
		return
	}
	var data bytes.Buffer
	enc := json.NewEncoder(&data)
	if err = enc.Encode(&payment); err != nil {
		return
	}
	resp, err := postResponse(u, &data)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	dec := json.NewDecoder(resp.Body)
	err = dec.Decode(&result)
	return
}

//ListPayFwds returns a PayFwds slice
//associated with your API.Token.
func (api *API) ListPayFwds() (payments []PayFwd, err error) {
	u, err := api.buildURL("/payments")
	resp, err := getResponse(u)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	//decode JSON into payments
	dec := json.NewDecoder(resp.Body)
	err = dec.Decode(&payments)
	return
}

//GetPayFwd returns a PayFwd based on its id.
func (api *API) GetPayFwd(id string) (payment PayFwd, err error) {
	u, err := api.buildURL("/payments/" + id)
	resp, err := getResponse(u)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	//decode JSON into payments
	dec := json.NewDecoder(resp.Body)
	err = dec.Decode(&payment)
	return
}

//DeletePayFwd deletes a PayFwd request from
//BlockCypher's database, based on its id.
func (api *API) DeletePayFwd(id string) (err error) {
	u, err := api.buildURL("/payments/" + id)
	resp, err := deleteResponse(u)
	if err != nil {
		return
	}
	resp.Body.Close()
	return
}

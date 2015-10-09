package blockcy

import (
	"bytes"
	"encoding/json"
)

//CreatePayment creates a new Payment forwarding
//request associated with your API.Token,
//and returns a result Payment with a
//BlockCypher-assigned ID.
func (api *API) CreatePayment(payment PayFwd) (result PayFwd, err error) {
	u, err := api.buildURL("/payments")
	if err != nil {
		return
	}
	//encode response into ReadWriter
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
	//Decode JSON into result
	dec := json.NewDecoder(resp.Body)
	err = dec.Decode(&result)
	return
}

//ListPayments returns a slice of Payments
//associated with your token.
func (api *API) ListPayments() (payments []PayFwd, err error) {
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

//DeletePayment deletes a Payment forwarding
//request from BlockCypher's database, based
//on its ID field.
func (api *API) DeletePayment(payment PayFwd) (err error) {
	u, err := api.buildURL("/payments/" + payment.ID)
	resp, err := deleteResponse(u)
	if err != nil {
		return
	}
	resp.Body.Close()
	return
}

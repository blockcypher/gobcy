package blockcy

import (
	"bytes"
	"encoding/json"
)

//ListPayments returns a slice of Payments
//associated with your API.Token.
func (api *API) ListPayments() (payments []PaymentFwd, err error) {
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

//PostPayment creates a new Payment forwarding
//request associated with your API.Token,
//and returns a result Payment with a
//BlockCypher-assigned Id.
func (api *API) PostPayment(payment PaymentFwd) (result PaymentFwd, err error) {
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

//DeletePayment deletes a Payment forwarding
//request from BlockCypher's database, based
//on its Id field.
func (api *API) DeletePayment(payment PaymentFwd) (err error) {
	u, err := api.buildURL("/payments/" + payment.ID)
	resp, err := deleteResponse(u)
	if err != nil {
		return
	}
	resp.Body.Close()
	return
}

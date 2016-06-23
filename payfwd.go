package gobcy

import "strconv"

//CreatePayFwd creates a new PayFwd forwarding
//request associated with your API.Token, and
//returns a PayFwd with a BlockCypher-assigned id.
func (api *API) CreatePayFwd(payment PayFwd) (result PayFwd, err error) {
	u, err := api.buildURL("/payments", nil)
	if err != nil {
		return
	}
	err = postResponse(u, &payment, &result)
	return
}

//ListPayFwds returns a PayFwds slice
//associated with your API.Token.
func (api *API) ListPayFwds() (payments []PayFwd, err error) {
	u, err := api.buildURL("/payments", nil)
	if err != nil {
		return
	}
	err = getResponse(u, &payments)
	return
}

//ListPayFwdsPage returns a PayFwds slice
//associated with your API.Token, starting at the start index.
//Useful for paging past the 200 payment forward limit.
func (api *API) ListPayFwdsPage(start int) (payments []PayFwd, err error) {
	params := map[string]string{"start": strconv.Itoa(start)}
	u, err := api.buildURL("/payments", params)
	if err != nil {
		return
	}
	err = getResponse(u, &payments)
	return
}

//GetPayFwd returns a PayFwd based on its id.
func (api *API) GetPayFwd(id string) (payment PayFwd, err error) {
	u, err := api.buildURL("/payments/"+id, nil)
	if err != nil {
		return
	}
	err = getResponse(u, &payment)
	return
}

//DeletePayFwd deletes a PayFwd request from
//BlockCypher's database, based on its id.
func (api *API) DeletePayFwd(id string) (err error) {
	u, err := api.buildURL("/payments/"+id, nil)
	if err != nil {
		return
	}
	err = deleteResponse(u)
	return
}

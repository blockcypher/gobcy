package gobcy

//CreateHook creates a new WebHook associated
//with your API.Token, and returns a WebHook
//with a BlockCypher-assigned id.
func (api *API) CreateHook(hook Hook) (result Hook, err error) {
	u, err := api.buildURL("/hooks", nil)
	if err != nil {
		return
	}
	err = postResponse(u, &hook, &result)
	return
}

//ListHooks returns a slice of WebHooks
//associated with your API.Token.
func (api *API) ListHooks() (hooks []Hook, err error) {
	u, err := api.buildURL("/hooks", nil)
	if err != nil {
		return
	}
	err = getResponse(u, &hooks)
	return
}

//GetHook returns a WebHook by its id.
func (api *API) GetHook(id string) (hook Hook, err error) {
	u, err := api.buildURL("/hooks/"+id, nil)
	if err != nil {
		return
	}
	err = getResponse(u, &hook)
	return
}

//DeleteHook deletes a WebHook notification
//from BlockCypher's database, based on its id.
func (api *API) DeleteHook(id string) (err error) {
	u, err := api.buildURL("/hooks/"+id, nil)
	if err != nil {
		return
	}
	err = deleteResponse(u)
	return
}

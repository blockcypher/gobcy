package blockcy

import (
	"bytes"
	"encoding/json"
)

//ListHooks returns a slice of WebHooks associated
//with your Config.Token.
func (api *API) ListHooks() (hooks []Hook, err error) {
	u, err := api.buildURL("/hooks")
	resp, err := getResponse(u)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	//decode JSON into hooks
	dec := json.NewDecoder(resp.Body)
	err = dec.Decode(&hooks)
	return
}

//PostHook creates a new WebHook associated
//with your Config.Token, and returns a result
//WebHook with a BlockCypher-assigned Id.
func (api *API) PostHook(hook Hook) (result Hook, err error) {
	u, err := api.buildURL("/hooks")
	if err != nil {
		return
	}
	//encode response into ReadWriter
	var data bytes.Buffer
	enc := json.NewEncoder(&data)
	if err = enc.Encode(&hook); err != nil {
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

//DeleteHook deletes a WebHook notification
//from BlockCypher's database, based on its
//Id field.
func (api *API) DeleteHook(hook Hook) (err error) {
	u, err := api.buildURL("/hooks/" + hook.ID)
	resp, err := deleteResponse(u)
	if err != nil {
		return
	}
	resp.Body.Close()
	return
}

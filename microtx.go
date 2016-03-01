package gobcy

import (
	"encoding/hex"
	"errors"

	"github.com/btcsuite/btcd/btcec"
)

//SendMicro sends a Micro through the Coin/Chain
//network. It will return a Micro with a proper hash
//if it successfully sent. If using public (instead of
//private) keys, you'll need to sign the returned Micro
//(using the *Micro.Sign method) and run SendMicro
//again with the signed data, which will then return
//a proper hash.
func (api *API) SendMicro(mic MicroTX) (result MicroTX, err error) {
	u, err := api.buildURL("/txs/micro", nil)
	if err != nil {
		return
	}
	err = postResponse(u, &mic, &result)
	return
}

//Sign takes a hex-encoded string slice of private
//keys and uses them to sign the ToSign data in a
//MicroTX, generating the proper hex-encoded Signatures.
//This is meant as a helperfunction, and leverages
//btcd's btcec library.
func (mic *MicroTX) Sign(priv string) (err error) {
	privDat, err := hex.DecodeString(priv)
	if err != nil {
		return
	}
	privkey, pubkey := btcec.PrivKeyFromBytes(btcec.S256(), privDat)
	if mic.Pubkey != hex.EncodeToString(pubkey.SerializeCompressed()) {
		err = errors.New("*MicroTX.Sign error: public key does not match private key")
		return
	}
	for _, k := range mic.ToSign {
		tosign, err := hex.DecodeString(k)
		if err != nil {
			return err
		}
		sig, err := privkey.Sign(tosign)
		if err != nil {
			return err
		}
		mic.Signatures = append(mic.Signatures, hex.EncodeToString(sig.Serialize()))
	}
	return
}

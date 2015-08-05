package blockcy

import "time"

//Blockchain represents information about
//the state of a blockchain.
type Blockchain struct {
	Name             string    `json:"name"`
	Height           int       `json:"height"`
	Hash             string    `json:"hash"`
	Time             time.Time `json:"time"`
	LatestURL        string    `json:"latest_url"`
	PrevHash         string    `json:"previous_hash"`
	PrevURL          string    `json:"previous_url"`
	PeerCount        int       `json:"peer_count"`
	HighFee          int       `json:"high_fee_per_kb"`
	MediumFee        int       `json:"medium_fee_per_kb"`
	LowFee           int       `json:"low_fee_per_kb"`
	UnconfirmedCount int       `json:"unconfirmed_count"`
	LastForkHeight   int       `json:"last_fork_height"`
	LastForkHash     string    `json:"last_fork_hash"`
}

//Block represents information about the state
//of a given block in a blockchain.
type Block struct {
	Hash         string    `json:"hash"`
	Height       int       `json:"height"`
	Depth        int       `json:"depth"`
	Chain        string    `json:"chain"`
	Total        int       `json:"total"`
	Fees         int       `json:"fees"`
	Ver          int       `json:"ver"`
	Time         time.Time `json:"time"`
	ReceivedTime time.Time `json:"received_time"`
	RelayedBy    string    `json:"relayed_by,omitempty"`
	Bits         int       `json:"bits"`
	Nonce        int       `json:"nonce"`
	NumTX        int       `json:"n_tx"`
	PrevBlock    string    `json:"prev_block"`
	PrevBlockURL string    `json:"prev_block_url"`
	MerkleRoot   string    `json:"mrkl_root"`
	TXids        []string  `json:"txids"`
	NextTXs      string    `json:"next_txids"`
}

//TX represents information about the state
//of a given transaction in a blockchain.
type TX struct {
	BlockHash     string     `json:"block_hash,omitempty"`
	BlockHeight   int        `json:"block_height,omitempty"`
	Hash          string     `json:"hash,omitempty"`
	Addresses     []string   `json:"addresses,omitempty"`
	Total         int        `json:"total,omitempty"`
	Fees          int        `json:"fees,omitempty"`
	Size          int        `json:"size"`
	Preference    string     `json:"preference,omitempty"`
	RelayedBy     string     `json:"relayed_by,omitempty"`
	Received      time.Time  `json:"received,omitempty"`
	Confirmed     time.Time  `json:"confirmed,omitempty"`
	Confirmations int        `json:"confirmations,omitempty"`
	Confidence    float64    `json:"confidence,omitempty"`
	Ver           int        `json:"ver,omitempty"`
	LockTime      int        `json:"lock_time,omitempty"`
	DoubleSpend   bool       `json:"double_spend,omitempty"`
	DoubleOf      string     `json:"double_of,omitempty"`
	ReceiveCount  int        `json:"receive_count,omitempty"`
	VinSize       int        `json:"vin_sz,omitempty"`
	VoutSize      int        `json:"vout_sz,omitempty"`
	Hex           string     `json:"hex,omitempty"`
	ChangeAddress string     `json:"change_address,omitempty"`
	NextInputs    string     `json:"next_inputs,omitempty"`
	NextOutputs   string     `json:"next_outputs,omitempty"`
	Inputs        []TXInput  `json:"inputs"`
	Outputs       []TXOutput `json:"outputs"`
}

//TXInput represents the state of a transaction input
type TXInput struct {
	PrevHash    string   `json:"prev_hash,omitempty"`
	OutputIndex int      `json:"output_index,omitempty"`
	OutputValue int      `json:"output_value,omitempty"`
	Addresses   []string `json:"addresses"`
	Sequence    int      `json:"sequence,omitempty"`
	ScriptType  string   `json:"script_type,omitempty"`
	Script      string   `json:"script,omitempty"`
	Age         int      `json:"age,omitempty"`
	WalletName  string   `json:"wallet_name,omitempty"`
	WalletToken string   `json:"wallet_token,omitempty"`
}

//TXOutput represents the state of a transaction output
type TXOutput struct {
	SpentBy    string   `json:"spent_by,omitempty"`
	Value      int      `json:"value"`
	Addresses  []string `json:"addresses"`
	ScriptType string   `json:"script_type,omitempty"`
	Script     string   `json:"script,omitempty"`
}

//TXConf represents information about the
//confidence of an unconfirmed transaction.
type TXConf struct {
	Age          int     `json:"age_millis"`
	ReceiveCount int     `json:"receive_count,omitempty"`
	Confidence   float64 `json:"confidence"`
	TXHash       string  `json:"txhash"`
}

//Addr represents information about the state
//of a public address.
type Addr struct {
	Address            string `json:"address"`
	TotalReceived      int    `json:"total_received"`
	TotalSent          int    `json:"total_sent"`
	Balance            int    `json:"balance"`
	UnconfirmedBalance int    `json:"unconfirmed_balance"`
	FinalBalance       int    `json:"final_balance"`
	NumTX              int    `json:"n_tx"`
	UnconfirmedNumTX   int    `json:"unconfirmed_n_tx"`
	FinalNumTX         int    `json:"final_n_tx"`
	TXs                []TX   `json:"txs"`
}

//AddrPair represents information about a generated
//public-private key pair from BlockCypher's address
//generation API. Large amounts are not recommended to be
//stored with these addresses.
type AddrPair struct {
	Address string `json:"address"`
	Private string `json:"private"`
	Public  string `json:"public"`
	Wif     string `json:"wif"`
}

//WebHook represents a WebHook event. You can include
//your Token for notification management, but otherwise
//it's optional.
//BlockCypher supports the following events:
//	Event = "unconfirmed-tx"
//	Event = "new-block"
//	Event = "confirmed-tx"
//	Event = "tx-confirmation"
//	Event = "double-spend-tx"
//Hash, Address, and Script are all optional; creating
//a WebHook with any of them will filter the resulting
//notifications, if appropriate. Id is returned by
//BlockCyphers servers after Posting a new WebHook; you
//shouldn't manually generate this field.
type WebHook struct {
	Id      string `json:"id,omitempty"`
	Token   string `json:"token,omitempty"`
	Event   string `json:"event"`
	Url     string `json:"url"`
	Hash    string `json:"hash,omitempty"`
	Address string `json:"address,omitempty"`
	Script  string `json:"script,omitempty"`
}

//Payment represents a reference to a payment forwarding
//request.
type Payment struct {
	Id             string   `json:"id,omitempty"`
	Token          string   `json:"token,omitempty"`
	Destination    string   `json:"destination"`
	InputAddr      string   `json:"input_address,omitempty"`
	ProcessAddr    string   `json:"process_fees_address,omitempty"`
	ProcessPercent float64  `json:"process_fees_percent,omitempty"`
	ProcessSatoshi int      `json:"process_fees_satoshis,omitempty"`
	CallbackUrl    string   `json:"callback_url,omitempty"`
	EnableConfirm  bool     `json:"enable_confirmations,omitempty"`
	MiningFees     int      `json:"mining_fees_satoshis,omitempty"`
	TXHistory      []string `json:"transactions,omitempty"`
}

//Payback represents a Payment Forwarding callback.
//It's more fun to call it a "payback."
type Payback struct {
	Value       int    `json:"value"`
	Destination string `json:"destination"`
	DestHash    string `json:"transaction_hash"`
	InputAddr   string `json:"input_address"`
	InputHash   string `json:"input_transaction_hash"`
}

//WIPTX represents the return call to BlockCypher's
//txs/new endpoint, and includes error information,
//hex transactions that need to be signed, and space
//for the signed transactions and associated public keys.
type WipTX struct {
	Errors     []wipTXerr `json:"errors,omitempty"`
	Trans      TX         `json:"tx"`
	ToSign     []string   `json:"tosign"`
	Signatures []string   `json:"signatures"`
	PubKeys    []string   `json:"pubkeys,omitempty"`
}

//used within WipTX for JSON serialization.
type wipTXerr struct {
	Error string `json:"error,omitempty"`
}

//Micro represents a microtransaction. For small-value
//transactions, BlockCypher will sign the transaction
//on your behalf, with your private key (if provided).
//Setting a separate change address is recommended.
//Where your application model allows it, consider
//only using public keys with microtransactions,
//and sign the microtransaction with your private key
//(without sending to BlockCypher's server.
type Micro struct {
	//Only one of Pubkey/Private/Wif is required
	Pubkey     string   `json:"from_pubkey,omitempty"`
	Private    string   `json:"from_private,omitempty"`
	Wif        string   `json:"from_wif,omitempty"`
	ToAddr     string   `json:"to_address"`
	ChangeAddr string   `json:"change_address,omitempty"`
	Value      int      `json:"value_satoshis"`
	Wait       bool     `json:"wait_guarantee,omitempty"`
	ToSign     []string `json:"tosign,omitempty"`
	Signatures []string `json:"signatures,omitempty"`
	Hash       string   `json:"hash,omitempty"`
}

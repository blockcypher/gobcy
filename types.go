package gobcy

import "time"

//TokenUsage represents information about
//the limits and usage against your token.
type TokenUsage struct {
	Limits      Usage          `json:"limits"`
	Hits        Usage          `json:"hits"`
	HitsHistory []UsageHistory `json:"hits_history"`
}

type Usage struct {
	PerSec       int `json:"api/second,omitempty"`
	PerHour      int `json:"api/hour,omitempty"`
	PerDay       int `json:"api/day,omitempty"`
	HooksPerHour int `json:"hooks/hour,omitempty"`
	ConfPerHour  int `json:"confidence/hour,omitempty"`
	Hooks        int `json:"hooks,omitempty"`
	PayFwds      int `json:"payments,omitempty"`
}

type UsageHistory struct {
	Usage
	Time time.Time `json:",omitempty"`
}

//Blockchain represents information about
//the state of a blockchain.
type Blockchain struct {
	Name             string    `json:"name"`
	Height           int       `json:"height"`
	Hash             string    `json:"hash"`
	Time             time.Time `json:"time"`
	PrevHash         string    `json:"previous_hash"`
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
	DataProtocol  string     `json:"data_protocol,omitempty"`
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
}

//TXOutput represents the state of a transaction output
type TXOutput struct {
	SpentBy    string   `json:"spent_by,omitempty"`
	Value      int      `json:"value"`
	Addresses  []string `json:"addresses"`
	ScriptType string   `json:"script_type,omitempty"`
	Script     string   `json:"script,omitempty"`
	DataHex    string   `json:"data_hex,omitempty"`
	DataString string   `json:"data_string,omitempty"`
}

//TXConf represents information about the
//confidence of an unconfirmed transaction.
type TXConf struct {
	Age          int     `json:"age_millis"`
	ReceiveCount int     `json:"receive_count,omitempty"`
	Confidence   float64 `json:"confidence"`
	TXHash       string  `json:"txhash"`
}

//TXRef represents summarized data about a
//transaction input or output.
type TXRef struct {
	Address       string    `json:"address,omitempty"`
	BlockHeight   int       `json:"block_height"`
	TXHash        string    `json:"tx_hash"`
	TXInputN      int       `json:"tx_input_n"`
	TXOutputN     int       `json:"tx_output_n"`
	Value         int       `json:"value"`
	Pref          string    `json:"preference"`
	Spent         bool      `json:"spent"`
	DoubleSpend   bool      `json:"double_spend"`
	DoubleOf      string    `json:"double_of,omitempty"`
	Confirmations int       `json:"confirmations"`
	Script        string    `json:"script,omitempty"`
	RefBalance    int       `json:"ref_balance,omitempty"`
	Confidence    float64   `json:"confidence,omitempty"`
	Confirmed     time.Time `json:"confirmed,omitempty"`
	SpentBy       string    `json:"spent_by,omitempty"`
	Received      time.Time `json:"received,omitempty"`
	ReceivedCount int       `json:"received_count,omitempty"`
}

//TXSkel represents the return call to BlockCypher's
//txs/new endpoint, and includes error information,
//hex transactions that need to be signed, and space
//for the signed transactions and associated public keys.
type TXSkel struct {
	Trans      TX       `json:"tx"`
	ToSign     []string `json:"tosign"`
	Signatures []string `json:"signatures"`
	PubKeys    []string `json:"pubkeys,omitempty"`
	ToSignTX   []string `json:"tosign_tx,omitempty"`
	Errors     []struct {
		Error string `json:"error,omitempty"`
	} `json:"errors,omitempty"`
}

//NullData represents the call and return to BlockCypher's
//Data API, allowing you to embed up to 80 bytes into
//a blockchain via an OP_RETURN.
type NullData struct {
	Data     string `json:"data"`
	Encoding string `json:"encoding,omitempty"`
	Hash     string `json:"hash,omitempty"`
}

//MicroTX represents a microtransaction. For small-value
//transactions, BlockCypher will sign the transaction
//on your behalf, with your private key (if provided).
//Setting a separate change address is recommended.
//Where your application model allows it, consider
//only using public keys with microtransactions,
//and sign the microtransaction with your private key
//(without sending to BlockCypher's server).
type MicroTX struct {
	//Only one of Pubkey/Private/Wif is required
	Pubkey     string   `json:"from_pubkey,omitempty"`
	Priv       string   `json:"from_private,omitempty"`
	Wif        string   `json:"from_wif,omitempty"`
	ToAddr     string   `json:"to_address"`
	Value      int      `json:"value_satoshis"`
	ChangeAddr string   `json:"change_address,omitempty"`
	Wait       bool     `json:"wait_guarantee,omitempty"`
	ToSign     []string `json:"tosign,omitempty"`
	Signatures []string `json:"signatures,omitempty"`
	Hash       string   `json:"hash,omitempty"`
	Inputs     []struct {
		PrevHash    string `json:"prev_hash"`
		OutputIndex int    `json:"output_index"`
	} `json:"inputs,omitempty"`
	Outputs []struct {
		Value   int    `json:"value"`
		Address string `json:"address"`
	} `json:"outputs,omitempty"`
	Fees int `json:"fees,omitempty"`
}

//Addr represents information about the state
//of a public address.
type Addr struct {
	Address            string   `json:"address,omitempty"`
	Wallet             Wallet   `json:"wallet,omitempty"`
	HDWallet           HDWallet `json:"hd_wallet,omitempty"`
	TotalReceived      int      `json:"total_received"`
	TotalSent          int      `json:"total_sent"`
	Balance            int      `json:"balance"`
	UnconfirmedBalance int      `json:"unconfirmed_balance"`
	FinalBalance       int      `json:"final_balance"`
	NumTX              int      `json:"n_tx"`
	UnconfirmedNumTX   int      `json:"unconfirmed_n_tx"`
	FinalNumTX         int      `json:"final_n_tx"`
	TXs                []TX     `json:"txs,omitempty"`
	TXRefs             []TXRef  `json:"txrefs,omitempty"`
	UnconfirmedTXRefs  []TXRef  `json:"unconfirmed_txrefs,omitempty"`
	HasMore            bool     `json:"hasMore,omitempty"`
}

//AddrKeychain represents information about a generated
//public-private key pair from BlockCypher's address
//generation API. Large amounts are not recommended to be
//stored with these addresses.
type AddrKeychain struct {
	Address         string   `json:"address,omitempty"`
	Private         string   `json:"private,omitempty"`
	Public          string   `json:"public,omitempty"`
	Wif             string   `json:"wif,omitempty"`
	PubKeys         []string `json:"pubkeys,omitempty"`
	ScriptType      string   `json:"script_type,omitempty"`
	OriginalAddress string   `json:"original_address,omitempty"`
	OAPAddress      string   `json:"oap_address,omitempty"`
}

//Wallet represents information about a standard wallet.
//Typically, wallets can be used wherever an address can be
//used within the API.
type Wallet struct {
	Name      string   `json:"name,omitempty"`
	Addresses []string `json:"addresses,omitempty"`
}

//HDWallet represents information about a Hierarchical Deterministic
//(HD) wallet. Like regular Wallets, HDWallets can be used wherever an
//address can be used within the API.
type HDWallet struct {
	Name            string `json:"name,omitempty"`
	ExtPubKey       string `json:"extended_public_key,omitempty"`
	SubchainIndexes []int  `json:"subchain_indexes,omitempty"`
	Chains          []struct {
		ChainAddr []struct {
			Address string `json:"address,omitempty"`
			Path    string `json:"path,omitempty"`
			Public  string `json:"public,omitempty"`
		} `json:"chain_addresses,omitempty"`
		Index int `json:"index,omitempty"`
	} `json:"chains,omitempty"`
}

//Hook represents a WebHook/WebSockets event.
//BlockCypher supports the following events:
//	Event = "unconfirmed-tx"
//	Event = "new-block"
//	Event = "confirmed-tx"
//	Event = "tx-confirmation"
//	Event = "double-spend-tx"
//  Event = "tx-confidence"
//Hash, Address, and Script are all optional; creating
//a WebHook with any of them will filter the resulting
//notifications, if appropriate. ID is returned by
//BlockCyphers servers after Posting a new WebHook; you
//shouldn't manually generate this field.
type Hook struct {
	ID            string  `json:"id,omitempty"`
	Event         string  `json:"event"`
	Hash          string  `json:"hash,omitempty"`
	WalletName    string  `json:"wallet_name,omitempty"`
	Address       string  `json:"address,omitempty"`
	Confirmations int     `json:"confirmations,omitempty"`
	Confidence    float32 `json:"confidence,omitempty"`
	Script        string  `json:"script,omitempty"`
	URL           string  `json:"url,omitempty"`
	CallbackErrs  int     `json:"callback_errors,omitempty"`
}

//PayFwd represents a reference to
//a Payment Forwarding request.
type PayFwd struct {
	ID             string   `json:"id,omitempty"`
	Destination    string   `json:"destination"`
	InputAddr      string   `json:"input_address,omitempty"`
	ProcessAddr    string   `json:"process_fees_address,omitempty"`
	ProcessPercent float64  `json:"process_fees_percent,omitempty"`
	ProcessValue   int      `json:"process_fees_satoshis,omitempty"`
	CallbackURL    string   `json:"callback_url,omitempty"`
	EnableConfirm  bool     `json:"enable_confirmations,omitempty"`
	MiningFees     int      `json:"mining_fees_satoshis,omitempty"`
	TXHistory      []string `json:"transactions,omitempty"`
}

//Payback represents a Payment Forwarding Callback.
//It's more fun to call it a "payback."
type Payback struct {
	Value       int    `json:"value"`
	Destination string `json:"destination"`
	DestHash    string `json:"transaction_hash"`
	InputAddr   string `json:"input_address"`
	InputHash   string `json:"input_transaction_hash"`
}

//OAPIssue represents a request for issuance or transfer of
//an Open Asset on a blockchain.
type OAPIssue struct {
	Priv     string `json:"from_private"`
	ToAddr   string `json:"to_address"`
	Amount   int    `json:"amount"`
	Metadata string `json:"metadata,omitempty"`
}

//OAPTX represents an Open Asset protocol transaction, generated
//when issuing or transferring assets.
type OAPTX struct {
	Ver         int       `json:"ver"`
	AssetID     string    `json:"assetid"`
	Hash        string    `json:"hash"`
	Confirmed   time.Time `json:"confirmed,omitempty"`
	Received    time.Time `json:"received"`
	Metadata    string    `json:"oap_meta,omitempty"`
	DoubleSpend bool      `json:"double_spend"`
	Inputs      []struct {
		PrevHash    string `json:"prev_hash"`
		OutputIndex int    `json:"output_index"`
		OAPAddress  string `json:"address"`
		OutputValue int    `json:"output_value"`
	} `json:"inputs"`
	Outputs []struct {
		OAPAddress      string `json:"address"`
		Value           int    `json:"value"`
		OrigOutputIndex int    `json:"original_output_index"`
	} `json:"outputs"`
}

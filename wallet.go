package blockcy

//CreateWallet creates a public-address watching wallet
//associated with this token/coin/chain, usable anywhere
//in the API where an Address might be used (just use
//the wallet name instead). For example, with checking
//a wallet name balance:
//  addr, err := api.GetAddrBal("your-wallet-name")
//Wallet is either a normal list of addresses, or an HD
//wallet derived via an extended public seed and subchains.
//Make sure your Wallet has "HD" set to "true" if you're
//requesting to construct an HD wallet.
func (api *API) CreateWallet(req Wallet) (wal Wallet, err error) {
}

//GetWallet gets a Wallet based on its name, the associated
//API token/coin/chain, and whether it's an HD wallet or
//not.
func (api *API) GetWallet(name string, hd bool) (wal Wallet, err error) {
}

//AddAddrWallet adds a slice of addresses to a named Wallet,
//associated with the API token/coin/chain. Only works with
//normal (non-HD) wallets; trying to add Addresses to an HD
//wallet will return an error.
func (api *API) AddAddrWallet(name string, addrs []string) (wal Wallet, err error) {
}

//GetAddrWallet returns a slice of addresses associated with
//a named Wallet, associated with the API token/coin/chain. Must
//denote whether a Wallet is HD or not.
func (api *API) GetAddrWallet(name string, hd bool) (addrs []string, err error) {
}

//DeleteAddrWallet deletes a slice of addresses associated with
//a named Wallet, associated with the API token/coin/chain. Does
//not work with HD wallets, as their address cannot be deleted.
func (api *API) DeleteAddrWallet(name string, addrs []string) (err error) {
}

//GenAddrWallet generates an Address associated with a named
//Wallet, associated with the API token/coin/chain. If a normal
//wallet, will also return the private/WIF key with the Address
//Keychain. If an HD wallet, will only return the public key and
//address within the Address Keychain. With an HD wallet,
//Can optionally include a subchain index to generate an address on
//a particular subchain; otherwise defaults to the first subchain
//on the wallet.
func (api *API) GenAddrWallet(name string, hd bool, subchain int) (wal Wallet, addr AddrKeychain, err error) {
}

//DeleteWallet deletes a named wallet associated with the
//API token/coin/chain. Must specify if it's an HD wallet.
func (api *API) DeleteWallet(name string, hd bool) (err error) {
}

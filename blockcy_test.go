package blockcy

import (
	"log"
	"os"
	"testing"
)

var Wallet1, Wallet2 AddrPair
var txhash1, txhash2 string

func TestMain(m *testing.M) {
	//Set Coin/Chain to BlockCypher testnet
	Config.Coin = "bcy"
	Config.Chain = "test"
	//Set Your Token
	Config.Token = ""
	//Create/fund the wallets
	var err error
	Wallet1, err = GenAddrPair()
	Wallet2, err = GenAddrPair()
	if err != nil {
		log.Fatal("Error generating test wallets: ", err)
	}
	txhash1, err = Wallet1.Faucet(1e5)
	txhash2, err = Wallet2.Faucet(2e5)
	if err != nil {
		log.Fatal("Error funding test wallets: ", err)
	}
	os.Exit(m.Run())
}

func TestGetChain(t *testing.T) {
	ch, err := GetChain()
	if err != nil {
		t.Error("Error encountered: ", err)
	}
	t.Logf("%+v\n", ch)
	return
}

func TestGetBlock(t *testing.T) {
	_, err := GetBlock(187621, "")
	if err != nil {
		t.Error("Error encountered: ", err)
	}
	bl, err := GetBlock(0, "0000ffeb0031885f2292475eac7f9c6f7bf5057e3b0017a09cd1994e71b431a4")
	if err != nil {
		t.Error("Error encountered: ", err)
	}
	t.Logf("%+v\n", bl)
	_, err = GetBlock(187621, "0000ffeb0031885f2292475eac7f9c6f7bf5057e3b0017a09cd1994e71b431a4")
	if err == nil {
		t.Error("Expected error when querying for both height and hash, did not receive one")
	}
	return
}

func TestUnTX(t *testing.T) {
	txs, err := GetUnTX()
	if err != nil {
		t.Error("Error encountered: ", err)
	}
	t.Logf("%+v\n", txs)
	return
}

func TestGetTX(t *testing.T) {
	tx, err := GetTX(txhash1)
	if err != nil {
		t.Error("Error encountered: ", err)
	}
	t.Logf("%+v\n", tx)
	return
}

func TestGetTXConf(t *testing.T) {
	tx, err := GetTXConf(txhash2)
	if err != nil {
		t.Error("Error encountered: ", err)
	}
	t.Logf("%+v\n", tx)
	return
}

func TestGetAddr(t *testing.T) {
	addr, err := GetAddr(Wallet1.Address)
	if err != nil {
		t.Error("Error encountered: ", err)
	}
	t.Logf("%+v\n", addr)
	return
}

func TestGetAddrFull(t *testing.T) {
	addr, err := GetAddrFull(Wallet2.Address, false)
	if err != nil {
		t.Error("Error encountered: ", err)
	}
	t.Logf("%+v\n", addr)
	return
}

func TestHooks(t *testing.T) {
	hook, err := PostHook(WebHook{Event: "new-block", Url: "https://my.domain.com/api/callbacks/doublespend?secret=justbetweenus"})
	if err != nil {
		t.Error("Error encountered: ", err)
	}
	t.Logf("%+v\n", hook)
	hooks, err := ListHooks()
	if err != nil {
		t.Error("Error encountered: ", err)
	}
	t.Logf("%+v\n", hooks)
	if err = DeleteHook(hooks[0]); err != nil {
		t.Error("Error encountered: ", err)
	}
	hooks, err = ListHooks()
	t.Logf("%+v\n", hooks)
	return
}

func TestPayments(t *testing.T) {
	pay, err := PostPayment(Payment{Destination: Wallet1.Address})
	if err != nil {
		t.Error("Error encountered: ", err)
	}
	t.Logf("%+v\n", pay)
	pays, err := ListPayments()
	if err != nil {
		t.Error("Error encountered: ", err)
	}
	t.Logf("%+v\n", pays)
	if err = DeletePayment(pay); err != nil {
		t.Error("Error encountered: ", err)
	}
	pays, err = ListPayments()
	t.Logf("%+v\n", pays)
	return
}

func TestNewTX(t *testing.T) {
	skel := SkelTX(Wallet2.Address, Wallet1.Address, 25000, false)
	wip, err := NewTX(skel)
	if err != nil {
		t.Error("Error encountered: ", err)
	}
	t.Logf("%+v\n", wip)
}

func TestMicro(t *testing.T) {
	mic := Micro{Private: Wallet2.Private, ToAddr: Wallet1.Address, Value: 25000}
	result, err := SendMicro(mic)
	if err != nil {
		t.Error("Error encountered: ", err)
	}
	t.Logf("%+v\n", result)
	txmic, err := GetTX(result.Hash)
	if err != nil {
		t.Error("Error encountered: ", err)
	}
	t.Logf("%+v\n", txmic)
}

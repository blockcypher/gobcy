package blockcy

import (
	"log"
	"os"
	"testing"
)

var Wallet1, Wallet2 AddrPair
var txhash1, txhash2 string
var bcy API

func TestMain(m *testing.M) {
	//Set Coin/Chain to BlockCypher testnet
	bcy.Coin = "bcy"
	bcy.Chain = "test"
	//Set Your Token
	bcy.Token = ""
	//Create/fund the wallets
	var err error
	Wallet1, err = bcy.GenAddrPair()
	Wallet2, err = bcy.GenAddrPair()
	if err != nil {
		log.Fatal("Error generating test wallets: ", err)
	}
	txhash1, err = bcy.Faucet(Wallet1, 1e5)
	txhash2, err = bcy.Faucet(Wallet2, 2e5)
	if err != nil {
		log.Fatal("Error funding test wallets: ", err)
	}
	os.Exit(m.Run())
}

func TestGetChain(t *testing.T) {
	ch, err := bcy.GetChain()
	if err != nil {
		t.Error("Error encountered: ", err)
	}
	t.Logf("%+v\n", ch)
	return
}

func TestGetBlock(t *testing.T) {
	_, err := bcy.GetBlock(187621, "")
	if err != nil {
		t.Error("Error encountered: ", err)
	}
	bl, err := bcy.GetBlock(0, "0000ffeb0031885f2292475eac7f9c6f7bf5057e3b0017a09cd1994e71b431a4")
	if err != nil {
		t.Error("Error encountered: ", err)
	}
	t.Logf("%+v\n", bl)
	_, err = bcy.GetBlock(187621, "0000ffeb0031885f2292475eac7f9c6f7bf5057e3b0017a09cd1994e71b431a4")
	if err == nil {
		t.Error("Expected error when querying for both height and hash, did not receive one")
	}
	return
}

func TestUnTX(t *testing.T) {
	txs, err := bcy.GetUnTX()
	if err != nil {
		t.Error("Error encountered: ", err)
	}
	t.Logf("%+v\n", txs)
	return
}

func TestGetTX(t *testing.T) {
	tx, err := bcy.GetTX(txhash1)
	if err != nil {
		t.Error("Error encountered: ", err)
	}
	t.Logf("%+v\n", tx)
	return
}

func TestGetTXConf(t *testing.T) {
	conf, err := bcy.GetTXConf(txhash2)
	if err != nil {
		t.Error("Error encountered: ", err)
	}
	t.Logf("%+v\n", conf)
	return
}

func TestGetAddr(t *testing.T) {
	addr, err := bcy.GetAddr(Wallet1.Address)
	if err != nil {
		t.Error("Error encountered: ", err)
	}
	t.Logf("%+v\n", addr)
	return
}

func TestGetAddrFull(t *testing.T) {
	addr, err := bcy.GetAddrFull(Wallet2.Address, false)
	if err != nil {
		t.Error("Error encountered: ", err)
	}
	t.Logf("%+v\n", addr)
	return
}

func TestHooks(t *testing.T) {
	hook, err := bcy.PostHook(WebHook{Event: "new-block", Url: "https://my.domain.com/api/callbacks/doublespend?secret=justbetweenus"})
	if err != nil {
		t.Error("Error encountered: ", err)
	}
	t.Logf("%+v\n", hook)
	hooks, err := bcy.ListHooks()
	if err != nil {
		t.Error("Error encountered: ", err)
	}
	t.Logf("%+v\n", hooks)
	if err = bcy.DeleteHook(hooks[0]); err != nil {
		t.Error("Error encountered: ", err)
	}
	hooks, err = bcy.ListHooks()
	t.Logf("%+v\n", hooks)
	return
}

func TestPayments(t *testing.T) {
	pay, err := bcy.PostPayment(Payment{Destination: Wallet1.Address})
	if err != nil {
		t.Error("Error encountered: ", err)
	}
	t.Logf("%+v\n", pay)
	pays, err := bcy.ListPayments()
	if err != nil {
		t.Error("Error encountered: ", err)
	}
	t.Logf("%+v\n", pays)
	if err = bcy.DeletePayment(pay); err != nil {
		t.Error("Error encountered: ", err)
	}
	pays, err = bcy.ListPayments()
	t.Logf("%+v\n", pays)
	return
}

func TestNewTX(t *testing.T) {
	skel := SkelTX(Wallet2.Address, Wallet1.Address, 25000, false)
	wip, err := bcy.NewTX(skel)
	if err != nil {
		t.Error("Error encountered: ", err)
	}
	t.Logf("%+v\n", wip)
}

func TestMicro(t *testing.T) {
	mic := Micro{Private: Wallet2.Private, ToAddr: Wallet1.Address, Value: 25000}
	result, err := bcy.SendMicro(mic)
	if err != nil {
		t.Error("Error encountered: ", err)
	}
	t.Logf("%+v\n", result)
	txmic, err := bcy.GetTX(result.Hash)
	if err != nil {
		t.Error("Error encountered: ", err)
	}
	t.Logf("%+v\n", txmic)
}

package tutorial

import (
	"ethereum/config"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"math"
	"math/big"
)

var a1 = common.HexToAddress("0xE5D477e97D7Ec6AcD908C9865CFECc998927164b")
var a2 = common.HexToAddress("0x9612f45f6e3C8D11cFda471501E5aD167B8d760b")

func GetEtherFromTestNetwork() {

	//generate2KeyStores()

	client , err := ethclient.Dial(config.Kovan)
	if err != nil {
		panic(err)
	}


	b1Wei, err := client.BalanceAt(ctx,a1, nil)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	fBlance1 := new(big.Float)
	fBlance1.SetString(b1Wei.String())
	b1Ether := new(big.Float).Quo(fBlance1, big.NewFloat(math.Pow10(18)))



	b2Wei, err := client.BalanceAt(ctx,a2, nil)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	fBlance2 := new(big.Float)
	fBlance2.SetString(b2Wei.String())
	b2Ether := new(big.Float).Quo(fBlance2, big.NewFloat(math.Pow10(18)))


	//fmt.Println("첫번째 지갑의 잔액 (Wei 단위):", b1Wei)
	//fmt.Println("두번째 지갑의 잔액 (Wei 단위):" ,b2Wei)

	fmt.Println("첫번째 지갑의 ether 잔액 :", b1Ether)
	fmt.Println("두번째 지갑의 ether 잔액 :", b2Ether)


}

func generate2KeyStores() {
	ks := keystore.NewKeyStore("./wallet", keystore.StandardScryptN, keystore.StandardScryptP)
	account1, err := ks.NewAccount(config.Password)
	if err != nil {
		panic(err)
	}
	account2, err := ks.NewAccount(config.Password)
	if err != nil {
		panic(err)
	}

	fmt.Println(account1.Address)
	//"0xE5D477e97D7Ec6AcD908C9865CFECc998927164b"
	fmt.Println(account2.Address)
	//"0x9612f45f6e3C8D11cFda471501E5aD167B8d760b"
}
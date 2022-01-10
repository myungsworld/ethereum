package tutorial

import (

	"ethereum/config"
	todo "ethereum/gen"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"io/ioutil"
	"math/big"
)

func DeploySmartContract() {
	b , err := ioutil.ReadFile("wallet/UTC--2022-01-05T05-14-39.630076000Z--e5d477e97d7ec6acd908c9865cfecc998927164b")
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	key , err := keystore.DecryptKey(b, config.Password)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	client, err := ethclient.Dial(config.Kovan)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	defer client.Close()

	add := crypto.PubkeyToAddress(key.PrivateKey.PublicKey)
	nonce , err := client.PendingNonceAt(ctx, add)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	gasPrice , err := client.SuggestGasPrice(ctx)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	chainID, err := client.NetworkID(ctx)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	auth, err := bind.NewKeyedTransactorWithChainID(key.PrivateKey, chainID)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	auth.GasPrice = gasPrice
	auth.GasLimit = uint64(300000)
	auth.Nonce = big.NewInt(int64(nonce))

	a , tx , _ , err := todo.DeployTodo(auth,client)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	fmt.Println("-----")
	fmt.Println(a.Hex())
	fmt.Println(tx.Hash().Hex())
	fmt.Println("-----")

}

package tutorial

import (
	"ethereum/config"
	todo "ethereum/gen"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"io/ioutil"
)

func InteractWithSmartContract() {
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

	// 사용할 컨트랙트 주소와 정보 struct에 넣기
	addressOfContract := common.HexToAddress("테스트넷에 배포한 내 컨트랙트 주소")
	t , err := todo.NewTodo(addressOfContract, client)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	// 트랜잭션 옵션 설정
	tx, err := bind.NewKeyedTransactorWithChainID(key.PrivateKey, chainID)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	tx.GasLimit = 3000000
	tx.GasPrice = gasPrice

	//// Write
	//tr, err :=t.Add(tx, "First Task")
	//if err != nil {
	//	fmt.Println(err)
	//	panic(err)
	//}
	//// 생성된 트랜잭션 가져오기
	//fmt.Println(tr.Hash())


	// Read
	add := crypto.PubkeyToAddress(key.PrivateKey.PublicKey)
	tasks, err := t.List(&bind.CallOpts{
		From: add,
	})
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	fmt.Println(tasks)

	// Update
	//tr , err := t.Update(tx, big.NewInt(0), "update task context")
	//if err != nil {
	//	fmt.Println(err)
	//	panic(err)
	//}
	//fmt.Println("tx :", tr.Hash())

	// Delete
	//tr , err := t.Delete(tx, big.NewInt(0))
	//if err != nil {
	//	fmt.Println(err)
	//	panic(err)
	//}
	//fmt.Println("tx :", tr.Hash())

}

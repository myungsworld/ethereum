package tutorial

import (
	"context"
	"ethereum/config"
	middlewares "ethereum/middleware"

	"fmt"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"io/ioutil"
)



func MakeTransaction() {

	client , err := ethclient.Dial(config.Kovan)

	// PendingNonceAt 으로 nonce 값 가져오기
	nonce, err := client.PendingNonceAt(ctx, a1)
	if err != nil {
		panic(err)
	}
	// 보낼 양 0.001 eth
	amount := middlewares.Ether("0.001")

	// SuggestGasPrice 로 평균 가스값 가져오기
	gasPrice , _ := client.SuggestGasPrice(context.Background())
	//gasEther := middlewares.WeiToEther(gasPrice)

	fmt.Println("마켓 가스 값 :", gasPrice)

	// 트랜잭션 생성
	tx := types.NewTransaction(nonce, a2, amount, 21000 , gasPrice , nil)

	// 네트워크 체인 아이디 가져오기
	chainID , err := client.NetworkID(ctx)

	// 송금할 주소의 프라이빗키 가져오기
	b , _ := ioutil.ReadFile("wallet/UTC--2022-01-05T05-14-39.630076000Z--e5d477e97d7ec6acd908c9865cfecc998927164b")
	key , _ := keystore.DecryptKey(b, config.Password)
	privateKey := key.PrivateKey

	// 트랜잭션 체인아이디와 프라이빗키 확인
	tx , err = types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)
	if err != nil {
		panic(err)
	}

	// 트랜잭션 전송
	err = client.SendTransaction(ctx, tx)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	fmt.Printf("트랜잭션 해시값 : %s\n", tx.Hash().Hex())



}

package tutorial

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"math/big"
)

var ctx = context.Background()

func GetLatestBlockNumber(client *ethclient.Client) *big.Int {

	// number 를 지정할수 있는데 nil 이면 마지막 블록 넘버를 가져옴
	block, err := client.BlockByNumber(ctx, nil)
	if err != nil {
		log.Fatalf("블럭 가져오기 실패:%v",err)
	}

	fmt.Println("이더리움 메인넷 마지막 블록 넘버:",block.Number())

	return block.Number()

}

func GetSpecificBlockHash(client *ethclient.Client) {
	block , err := client.BlockByHash(ctx, common.HexToHash("0x6b1a5cdc5e51172c2f4eb97e3fea0b99b4cac1372fba99f7e4d3b75bc32f0472"))
	if err != nil {
		fmt.Println("client.BlockByHash Err :", err )
		panic(err)
	}

	fmt.Println(block)
}
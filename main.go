package main

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"math"
	"math/big"
)

var infuraURL = "https://mainnet.infura.io/v3/e983371ca418429cb62af7f7f7767830"

// sudo npm install -g ganache-cli
// ganache-cli ( 테스트 로컬 메인넷 사용 )
var ganacheURL = "http://localhost:8545"
func main() {

	client , err := ethclient.DialContext(context.Background(), infuraURL)
	if err != nil {
		log.Fatalf("클라 접속 실패:%v", err)
	}
	defer client.Close()

	ctx := context.Background()

	// nil 이면 마지막 블록을 가져옴
	block, err := client.BlockByNumber(ctx, nil)
	if err != nil {
		log.Fatalf("블럭 가져오기 실패:%v",err)
	}

	fmt.Println("이더리움 메인넷 마지막 블록 넘버:",block.Number())

	addr := "0xca02616f14cb40bbc429886807967d84a07c256a"
	address := common.HexToAddress(addr)

	//웨이(wei)는 이더리움에서 사용하는 암호화폐인 이더의 가장 작은 단위이다. 1 웨이(wei)는 10-18 이더(ether)와 같다. 다시 말하면, 1 이더(ether)는 1018 웨이(wei)와 같다
	// 1 이더리움 = 10^18 wei
	balanceWei , err := client.BalanceAt(ctx,address,nil)
	fmt.Println("웨이 단위 (10의 -18승):",balanceWei)
	fBlance := new(big.Float)
	fBlance.SetString(balanceWei.String())
	balanceEther := new(big.Float).Quo(fBlance, big.NewFloat(math.Pow10(18)))
	fmt.Println("이더 단위 :", balanceEther)
}
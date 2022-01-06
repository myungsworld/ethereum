package tutorial

import (
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"math"
	"math/big"
)

func GetEtherAmountFromAddress(client *ethclient.Client , addr string) {
	address := common.HexToAddress(addr)

	//웨이(wei)는 이더리움에서 사용하는 암호화폐인 이더의 가장 작은 단위이다. 1 웨이(wei)는 10-18 이더(ether)와 같다. 다시 말하면, 1 이더(ether)는 1018 웨이(wei)와 같다
	//1 이더리움 = 10^18 wei
	balanceWei , err := client.BalanceAt(ctx,address,nil)
	if err != nil {
		fmt.Println("client.BalancerAt Err :" , err)
		panic(err)
	}

	fBlance := new(big.Float)
	fBlance.SetString(balanceWei.String())
	balanceEther := new(big.Float).Quo(fBlance, big.NewFloat(math.Pow10(18)))

	fmt.Println("지갑주소 :", addr)
	fmt.Println("웨이 단위 (10의 -18승):",balanceWei)
	fmt.Println("이더 단위 :", balanceEther)

}

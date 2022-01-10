package main

import (
	"context"
	"ethereum/config"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
)

func main() {

	// config.TestNet = 테스트넷 , config.EtherMainNet = 이더리움 메인 네트워크
	client, err := ethclient.DialContext(context.Background(), config.TestNet)
	if err != nil {
		log.Fatalf("이더리움 메인넷 접속 실패:%v", err)
	}
	defer client.Close()

	// 마지막 블록 가져오기
	//tutorial.GetLatestBlockNumber(client)

	// 해시값으로 블록 가져오기
	//tutorial.GetSpecificBlockHash(client)

	// 지갑주소로 balance 가져오기
	//addr := "0xca02616f14cb40bbc429886807967d84a07c256a"
	//tutorial.GetEtherAmountFromAddress(client, addr)

	// 이더 지갑 생성
	//tutorial.GenerateWallet()

	// 암호화 지갑 생성
	//tutorial.GenerateKeystore()

	// 테스트 네트워크에서 이더 가져오기
	//tutorial.GetEtherFromTestNetwork()

	// 트랜잭션 생성 ( 계좌송금 )
	//tutorial.GetEtherFromTestNetwork()
	//tutorial.MakeTransaction()

	// 스마트컨트랙트 배포
	//tutorial.DeploySmartContract()

	// 스마트컨트랙트 사용
	//tutorial.InteractWithSmartContract()
}

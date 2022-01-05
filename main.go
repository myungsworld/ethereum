package main

import (
	"context"
	"ethereum/config"
	"ethereum/tutorial"
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

	// 테스트 네트워크에서 이더 가져와보기
	tutorial.GetEtherFromTestNetwork()
}

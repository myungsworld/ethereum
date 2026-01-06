package main

import (
	"context"
	"fmt"
	"log"

	"github.com/myungsworld/ethereum/internal/config"
	"github.com/myungsworld/ethereum/internal/infrastructure/ethereum"
	"github.com/myungsworld/ethereum/internal/usecase"
)

func main() {
	ctx := context.Background()

	// 0. Config 로드
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("설정 로드 실패: %v", err)
	}
	fmt.Printf("환경: %s\n", cfg.Env)

	// 1. Infrastructure 레이어: 이더리움 클라이언트 생성
	var client *ethereum.Client
	var ethUsecase *usecase.EthereumUsecase

	if cfg.HasPrivateKey() {
		// 지갑이 있는 클라이언트 생성
		client, err = ethereum.NewClientWithWallet(cfg.RpcURL, cfg.PrivateKey, cfg.ChainID)
		if err != nil {
			log.Fatalf("클라이언트 생성 실패: %v", err)
		}
		fmt.Printf("지갑 주소: %s\n", client.Address())
		ethUsecase = usecase.NewEthereumUsecaseWithWriter(client)
	} else {
		// 읽기 전용 클라이언트 생성
		client, err = ethereum.NewClient(cfg.RpcURL)
		if err != nil {
			log.Fatalf("클라이언트 생성 실패: %v", err)
		}
		ethUsecase = usecase.NewEthereumUsecase(client)
	}
	defer client.Close()

	fmt.Println("이더리움 네트워크 연결 성공!")

	// 2. 네트워크 정보 조회
	networkInfo, err := ethUsecase.GetNetworkInfo(ctx)
	if err != nil {
		log.Fatalf("네트워크 정보 조회 실패: %v", err)
	}
	fmt.Printf("Chain ID: %d\n", networkInfo.ChainID)
	fmt.Printf("최신 블록 번호: %d\n", networkInfo.BlockNumber)

	// 3. 내 지갑 잔액 조회 (지갑이 있는 경우)
	if client.HasWallet() {
		fmt.Println("\n=== 내 지갑 정보 ===")
		myAccount, err := ethUsecase.GetAccountInfo(ctx, client.Address())
		if err != nil {
			log.Fatalf("계정 정보 조회 실패: %v", err)
		}
		fmt.Printf("주소: %s\n", myAccount.Address)
		fmt.Printf("잔액: %.6f ETH\n", myAccount.BalanceETH)
		fmt.Printf("Nonce: %d\n", myAccount.Nonce)

		// 4. ETH 전송 테스트 (잔액이 있는 경우만)
		if myAccount.BalanceETH > 0.001 {
			fmt.Println("\n=== ETH 전송 테스트 ===")
			// 자기 자신에게 0.001 ETH 전송 (테스트용)
			result, err := ethUsecase.SendETH(ctx, &usecase.SendETHRequest{
				To:     client.Address(), // 자기 자신에게 전송
				Amount: 0.001,
			})
			if err != nil {
				log.Printf("ETH 전송 실패: %v", err)
			} else {
				fmt.Printf("전송 성공!\n")
				fmt.Printf("Tx Hash: %s\n", result.TxHash)
				fmt.Printf("To: %s\n", result.To)
				fmt.Printf("Amount: %.6f ETH\n", result.Amount)
				fmt.Printf("Sepolia Etherscan: https://sepolia.etherscan.io/tx/%s\n", result.TxHash)
			}
		} else {
			fmt.Println("\n테스트 ETH가 부족합니다. Faucet에서 받아주세요:")
			fmt.Println("https://sepoliafaucet.com")
			fmt.Printf("주소: %s\n", client.Address())
		}
	}
}

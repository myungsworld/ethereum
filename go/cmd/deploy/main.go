package main

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"

	"github.com/myungsworld/ethereum/internal/config"
	"github.com/myungsworld/ethereum/internal/infrastructure/contracts"
)

func main() {
	ctx := context.Background()

	// 1. 설정 로드
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("설정 로드 실패: %v", err)
	}
	fmt.Printf("환경: %s\n", cfg.Env)

	if !cfg.HasPrivateKey() {
		log.Fatal("개인키가 설정되지 않음. .env 파일의 PRIVATE_KEY를 확인하세요.")
	}

	// 2. 이더리움 클라이언트 연결
	client, err := ethclient.Dial(cfg.RpcURL)
	if err != nil {
		log.Fatalf("이더리움 연결 실패: %v", err)
	}
	defer client.Close()
	fmt.Println("이더리움 네트워크 연결 성공!")

	// 3. 개인키 로드
	privateKey, err := crypto.HexToECDSA(cfg.PrivateKey)
	if err != nil {
		log.Fatalf("개인키 파싱 실패: %v", err)
	}

	publicKey := privateKey.Public().(*ecdsa.PublicKey)
	fromAddress := crypto.PubkeyToAddress(*publicKey)
	fmt.Printf("배포 주소: %s\n", fromAddress.Hex())

	// 4. 잔액 확인
	balance, err := client.BalanceAt(ctx, fromAddress, nil)
	if err != nil {
		log.Fatalf("잔액 조회 실패: %v", err)
	}
	ethBalance := new(big.Float).Quo(new(big.Float).SetInt(balance), big.NewFloat(1e18))
	fmt.Printf("잔액: %s ETH\n", ethBalance.String())

	if balance.Cmp(big.NewInt(0)) == 0 {
		fmt.Println("\n테스트 ETH가 없습니다. Faucet에서 받아주세요:")
		fmt.Println("https://cloud.google.com/application/web3/faucet/ethereum/sepolia")
		fmt.Printf("주소: %s\n", fromAddress.Hex())
		return
	}

	// 5. 트랜잭션 옵션 설정
	nonce, err := client.PendingNonceAt(ctx, fromAddress)
	if err != nil {
		log.Fatalf("nonce 조회 실패: %v", err)
	}

	gasPrice, err := client.SuggestGasPrice(ctx)
	if err != nil {
		log.Fatalf("가스 가격 조회 실패: %v", err)
	}

	chainID := big.NewInt(cfg.ChainID)

	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err != nil {
		log.Fatalf("Transactor 생성 실패: %v", err)
	}
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)
	auth.GasLimit = uint64(300000)
	auth.GasPrice = gasPrice

	// 6. 컨트랙트 배포
	fmt.Println("\n=== Counter 컨트랙트 배포 중... ===")
	initialCount := big.NewInt(0) // 초기값 0

	address, tx, _, err := contracts.DeployCounter(auth, client, initialCount)
	if err != nil {
		log.Fatalf("컨트랙트 배포 실패: %v", err)
	}

	fmt.Printf("트랜잭션 전송됨: %s\n", tx.Hash().Hex())
	fmt.Printf("컨트랙트 주소 (예상): %s\n", address.Hex())
	fmt.Printf("\nEtherscan에서 확인:\n")
	fmt.Printf("https://sepolia.etherscan.io/tx/%s\n", tx.Hash().Hex())

	// 7. 트랜잭션 확인 대기
	fmt.Println("\n트랜잭션 확인 대기 중...")
	receipt, err := bind.WaitMined(ctx, client, tx)
	if err != nil {
		log.Fatalf("트랜잭션 확인 실패: %v", err)
	}

	if receipt.Status == 1 {
		fmt.Println("\n=== 배포 성공! ===")
		fmt.Printf("컨트랙트 주소: %s\n", receipt.ContractAddress.Hex())
		fmt.Printf("가스 사용량: %d\n", receipt.GasUsed)
		fmt.Printf("블록 번호: %d\n", receipt.BlockNumber.Uint64())
	} else {
		fmt.Println("배포 실패!")
	}
}

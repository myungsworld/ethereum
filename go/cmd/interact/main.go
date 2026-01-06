package main

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"

	"github.com/myungsworld/ethereum/internal/config"
	"github.com/myungsworld/ethereum/internal/infrastructure/contracts"
)

// =============================================================================
// 배포된 컨트랙트 주소
// 이 주소는 블록체인에 영구적으로 저장된 Counter 컨트랙트의 위치
// =============================================================================
const contractAddress = "0x6CE93619B4A2C81ABf0FB86701493D2d6903B8D9"

func main() {
	ctx := context.Background()

	fmt.Println("=============================================================================")
	fmt.Println("                    스마트 컨트랙트 상호작용 예제")
	fmt.Println("=============================================================================")
	fmt.Println()

	// =========================================================================
	// STEP 1: 설정 로드
	// .env 파일에서 RPC URL, Chain ID, Private Key를 읽어옴
	// =========================================================================
	fmt.Println("[STEP 1] 설정 로드")
	fmt.Println("─────────────────────────────────────────────────────────────────────────────")

	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("설정 로드 실패: %v", err)
	}

	fmt.Printf("  환경(ENV):     %s\n", cfg.Env)
	fmt.Printf("  RPC URL:       %s\n", cfg.RpcURL)
	fmt.Printf("  Chain ID:      %d\n", cfg.ChainID)
	fmt.Printf("  개인키 설정:   %v\n", cfg.HasPrivateKey())
	fmt.Println()

	// =========================================================================
	// STEP 2: 이더리움 노드 연결
	// ethclient.Dial()로 JSON-RPC를 통해 이더리움 노드에 연결
	// 이 연결을 통해 블록체인 데이터를 읽고 트랜잭션을 전송
	// =========================================================================
	fmt.Println("[STEP 2] 이더리움 노드 연결")
	fmt.Println("─────────────────────────────────────────────────────────────────────────────")

	client, err := ethclient.Dial(cfg.RpcURL)
	if err != nil {
		log.Fatalf("이더리움 연결 실패: %v", err)
	}
	defer client.Close()

	// 연결 확인 - 최신 블록 번호 조회
	blockNumber, err := client.BlockNumber(ctx)
	if err != nil {
		log.Fatalf("블록 번호 조회 실패: %v", err)
	}

	fmt.Printf("  연결 상태:     성공\n")
	fmt.Printf("  최신 블록:     %d\n", blockNumber)
	fmt.Println()

	// =========================================================================
	// STEP 3: 컨트랙트 인스턴스 생성
	//
	// contracts.NewCounter()는 abigen이 자동 생성한 함수
	// - 첫 번째 인자: 컨트랙트 주소 (블록체인에서 컨트랙트가 위치한 곳)
	// - 두 번째 인자: ethclient (블록체인과 통신하는 클라이언트)
	//
	// 반환되는 'counter' 객체를 통해 컨트랙트의 모든 함수를 호출할 수 있음
	// =========================================================================
	fmt.Println("[STEP 3] 컨트랙트 인스턴스 생성")
	fmt.Println("─────────────────────────────────────────────────────────────────────────────")

	address := common.HexToAddress(contractAddress)
	fmt.Printf("  컨트랙트 주소: %s\n", address.Hex())

	// counter는 *contracts.Counter 타입
	// 이 객체에는 GetCount(), Owner(), Increment(), Decrement(), SetCount() 메서드가 있음
	counter, err := contracts.NewCounter(address, client)
	if err != nil {
		log.Fatalf("컨트랙트 로드 실패: %v", err)
	}
	fmt.Printf("  인스턴스 생성: 성공\n")
	fmt.Printf("  타입:          *contracts.Counter\n")
	fmt.Println()

	fmt.Println("  Counter 컨트랙트 구조:")
	fmt.Println("  ┌─────────────────────────────────────────────────────────────┐")
	fmt.Println("  │ 상태 변수 (Storage)                                         │")
	fmt.Println("  │   - count: uint256 (카운터 값)                               │")
	fmt.Println("  │   - owner: address (컨트랙트 배포자)                         │")
	fmt.Println("  ├─────────────────────────────────────────────────────────────┤")
	fmt.Println("  │ 읽기 함수 (View) - 가스 무료                                 │")
	fmt.Println("  │   - getCount() → uint256                                    │")
	fmt.Println("  │   - owner() → address                                       │")
	fmt.Println("  ├─────────────────────────────────────────────────────────────┤")
	fmt.Println("  │ 쓰기 함수 (State Change) - 가스 필요                        │")
	fmt.Println("  │   - increment(): count를 1 증가                             │")
	fmt.Println("  │   - decrement(): count를 1 감소                             │")
	fmt.Println("  │   - setCount(n): count를 n으로 설정 (owner만)               │")
	fmt.Println("  └─────────────────────────────────────────────────────────────┘")
	fmt.Println()

	// =========================================================================
	// STEP 4: 읽기 작업 (Call)
	//
	// Call vs Transaction:
	// - Call: 블록체인 상태를 읽기만 함, 가스 불필요, 즉시 결과 반환
	// - Transaction: 블록체인 상태를 변경함, 가스 필요, 채굴 후 확정
	//
	// bind.CallOpts: Call 옵션 설정
	// - Context: 타임아웃, 취소 등 컨텍스트
	// - Pending: true면 pending 상태 포함 (기본: false)
	// - BlockNumber: 특정 블록 기준으로 조회 (기본: latest)
	// =========================================================================
	fmt.Println("[STEP 4] 읽기 작업 (Call) - 가스 불필요")
	fmt.Println("─────────────────────────────────────────────────────────────────────────────")

	callOpts := &bind.CallOpts{
		Context: ctx,
		Pending: false, // pending 상태 제외, 확정된 상태만 조회
	}

	fmt.Println("  CallOpts 설정:")
	fmt.Printf("    - Context: %v\n", callOpts.Context != nil)
	fmt.Printf("    - Pending: %v\n", callOpts.Pending)
	fmt.Println()

	// getCount() 호출
	// Solidity: function getCount() public view returns (uint256)
	// Go에서는 *big.Int로 반환됨
	fmt.Println("  [4-1] counter.GetCount(callOpts) 호출")
	count, err := counter.GetCount(callOpts)
	if err != nil {
		log.Fatalf("  getCount 실패: %v", err)
	}
	fmt.Printf("    → 반환값: %d (타입: *big.Int)\n", count)
	fmt.Printf("    → 의미: 현재 카운터 값\n")
	fmt.Println()

	// owner() 호출
	// Solidity: address public owner;
	// public 변수는 자동으로 getter 함수가 생성됨
	fmt.Println("  [4-2] counter.Owner(callOpts) 호출")
	owner, err := counter.Owner(callOpts)
	if err != nil {
		log.Fatalf("  owner 실패: %v", err)
	}
	fmt.Printf("    → 반환값: %s\n", owner.Hex())
	fmt.Printf("    → 의미: 컨트랙트 배포자 주소\n")
	fmt.Println()

	// =========================================================================
	// STEP 5: 쓰기 작업 (Transaction)
	//
	// 트랜잭션을 보내려면:
	// 1. 개인키로 서명해야 함
	// 2. 가스비를 지불해야 함
	// 3. nonce를 설정해야 함 (이중 지불 방지)
	//
	// bind.TransactOpts: 트랜잭션 옵션 설정
	// - Nonce: 트랜잭션 순서 번호
	// - GasLimit: 최대 가스 사용량
	// - GasPrice: 가스 당 가격 (wei)
	// - Value: 전송할 ETH (wei)
	// =========================================================================
	if !cfg.HasPrivateKey() {
		fmt.Println("[STEP 5] 쓰기 작업 - 건너뜀 (개인키 없음)")
		fmt.Println("─────────────────────────────────────────────────────────────────────────────")
		fmt.Println("  개인키가 없어서 쓰기 작업을 수행할 수 없습니다.")
		fmt.Println("  .env 파일에 PRIVATE_KEY를 설정하세요.")
		return
	}

	fmt.Println("[STEP 5] 쓰기 작업 (Transaction) - 가스 필요")
	fmt.Println("─────────────────────────────────────────────────────────────────────────────")

	// 개인키 파싱
	fmt.Println("  [5-1] 개인키 → 공개키 → 주소 파생")
	privateKey, err := crypto.HexToECDSA(cfg.PrivateKey)
	if err != nil {
		log.Fatalf("  개인키 파싱 실패: %v", err)
	}
	publicKey := privateKey.Public().(*ecdsa.PublicKey)
	fromAddress := crypto.PubkeyToAddress(*publicKey)

	fmt.Printf("    Private Key: %s...(숨김)\n", cfg.PrivateKey[:8])
	fmt.Printf("    Public Key:  ECDSA secp256k1\n")
	fmt.Printf("    Address:     %s\n", fromAddress.Hex())
	fmt.Println()

	// 트랜잭션 파라미터 조회
	fmt.Println("  [5-2] 트랜잭션 파라미터 조회")
	nonce, err := client.PendingNonceAt(ctx, fromAddress)
	if err != nil {
		log.Fatalf("  nonce 조회 실패: %v", err)
	}
	fmt.Printf("    Nonce: %d (이 주소에서 보낸 트랜잭션 수)\n", nonce)

	gasPrice, err := client.SuggestGasPrice(ctx)
	if err != nil {
		log.Fatalf("  가스 가격 조회 실패: %v", err)
	}
	gasPriceGwei := new(big.Float).Quo(new(big.Float).SetInt(gasPrice), big.NewFloat(1e9))
	fmt.Printf("    Gas Price: %s Gwei (네트워크 추천 가격)\n", gasPriceGwei.String())

	balance, _ := client.BalanceAt(ctx, fromAddress, nil)
	balanceEth := new(big.Float).Quo(new(big.Float).SetInt(balance), big.NewFloat(1e18))
	fmt.Printf("    Balance: %s ETH\n", balanceEth.String())
	fmt.Println()

	// TransactOpts 생성
	fmt.Println("  [5-3] TransactOpts 생성")
	chainID := big.NewInt(cfg.ChainID)
	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err != nil {
		log.Fatalf("  Transactor 생성 실패: %v", err)
	}
	auth.Nonce = big.NewInt(int64(nonce))
	auth.GasLimit = uint64(100000) // 가스 리밋 (최대 사용량)
	auth.GasPrice = gasPrice
	auth.Value = big.NewInt(0) // 전송할 ETH 없음

	fmt.Printf("    auth.Nonce:    %d\n", auth.Nonce)
	fmt.Printf("    auth.GasLimit: %d\n", auth.GasLimit)
	fmt.Printf("    auth.GasPrice: %s wei\n", auth.GasPrice.String())
	fmt.Printf("    auth.Value:    %d wei (ETH 전송 없음)\n", auth.Value)
	fmt.Println()

	// increment() 호출
	fmt.Println("  [5-4] counter.Increment(auth) 호출")
	fmt.Println("    → Solidity: function increment() public")
	fmt.Println("    → 동작: count += 1; emit CountChanged(count, msg.sender);")
	fmt.Println()

	tx, err := counter.Increment(auth)
	if err != nil {
		log.Fatalf("  increment 실패: %v", err)
	}

	fmt.Println("  [5-5] 트랜잭션 전송됨")
	fmt.Printf("    Tx Hash:  %s\n", tx.Hash().Hex())
	fmt.Printf("    To:       %s (컨트랙트)\n", tx.To().Hex())
	fmt.Printf("    Gas:      %d\n", tx.Gas())
	fmt.Printf("    GasPrice: %s wei\n", tx.GasPrice().String())
	fmt.Printf("    Data:     %x (함수 호출 데이터)\n", tx.Data()[:4]) // 함수 selector만
	fmt.Println()

	// 트랜잭션 확인 대기
	fmt.Println("  [5-6] 트랜잭션 채굴 대기 (블록에 포함될 때까지)")
	fmt.Println("    → 블록 생성까지 약 12초 소요...")

	receipt, err := bind.WaitMined(ctx, client, tx)
	if err != nil {
		log.Fatalf("  트랜잭션 확인 실패: %v", err)
	}

	fmt.Println()
	fmt.Println("  [5-7] 트랜잭션 영수증 (Receipt)")
	fmt.Printf("    Status:       %d (1=성공, 0=실패)\n", receipt.Status)
	fmt.Printf("    Block Number: %d\n", receipt.BlockNumber.Uint64())
	fmt.Printf("    Gas Used:     %d\n", receipt.GasUsed)
	fmt.Printf("    Logs:         %d개의 이벤트\n", len(receipt.Logs))

	// 가스비 계산
	gasCost := new(big.Int).Mul(gasPrice, big.NewInt(int64(receipt.GasUsed)))
	gasCostEth := new(big.Float).Quo(new(big.Float).SetInt(gasCost), big.NewFloat(1e18))
	fmt.Printf("    실제 가스비:  %s ETH\n", gasCostEth.String())
	fmt.Println()

	// =========================================================================
	// STEP 6: 결과 확인
	// =========================================================================
	fmt.Println("[STEP 6] 결과 확인")
	fmt.Println("─────────────────────────────────────────────────────────────────────────────")

	newCount, _ := counter.GetCount(callOpts)
	fmt.Printf("  이전 카운트: %d\n", count)
	fmt.Printf("  현재 카운트: %d\n", newCount)
	fmt.Printf("  변화량:      +%d\n", new(big.Int).Sub(newCount, count))
	fmt.Println()

	fmt.Println("=============================================================================")
	fmt.Println("                              완료!")
	fmt.Println("=============================================================================")
	fmt.Println()
	fmt.Println("Etherscan에서 확인:")
	fmt.Printf("  트랜잭션: https://sepolia.etherscan.io/tx/%s\n", tx.Hash().Hex())
	fmt.Printf("  컨트랙트: https://sepolia.etherscan.io/address/%s\n", contractAddress)
}

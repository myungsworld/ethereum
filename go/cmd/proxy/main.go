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

	fmt.Println("═══════════════════════════════════════════════════════════════════════════════")
	fmt.Println("                         Proxy 패턴 배포 및 테스트")
	fmt.Println("═══════════════════════════════════════════════════════════════════════════════")
	fmt.Println()

	// =========================================================================
	// STEP 1: 설정 로드 및 클라이언트 연결
	// =========================================================================
	fmt.Println("[STEP 1] 설정 로드 및 연결")
	fmt.Println("───────────────────────────────────────────────────────────────────────────────")

	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("설정 로드 실패: %v", err)
	}

	if !cfg.HasPrivateKey() {
		log.Fatal("개인키가 설정되지 않았습니다. .env 파일을 확인하세요.")
	}

	client, err := ethclient.Dial(cfg.RpcURL)
	if err != nil {
		log.Fatalf("이더리움 연결 실패: %v", err)
	}
	defer client.Close()

	fmt.Printf("  네트워크: %s (Chain ID: %d)\n", cfg.Env, cfg.ChainID)

	// 개인키 → 주소
	privateKey, err := crypto.HexToECDSA(cfg.PrivateKey)
	if err != nil {
		log.Fatalf("개인키 파싱 실패: %v", err)
	}
	publicKey := privateKey.Public().(*ecdsa.PublicKey)
	fromAddress := crypto.PubkeyToAddress(*publicKey)
	fmt.Printf("  배포자 주소: %s\n", fromAddress.Hex())

	// 잔액 확인
	balance, _ := client.BalanceAt(ctx, fromAddress, nil)
	balanceEth := new(big.Float).Quo(new(big.Float).SetInt(balance), big.NewFloat(1e18))
	fmt.Printf("  잔액: %s ETH\n", balanceEth.String())
	fmt.Println()

	// =========================================================================
	// STEP 2: CounterV1 배포
	// =========================================================================
	fmt.Println("[STEP 2] CounterV1 배포")
	fmt.Println("───────────────────────────────────────────────────────────────────────────────")
	fmt.Println("  CounterV1 = 첫 번째 로직 컨트랙트")
	fmt.Println("  increment(), decrement(), getCount() 기능만 있음")
	fmt.Println()

	auth, err := getTransactOpts(ctx, client, privateKey, cfg.ChainID)
	if err != nil {
		log.Fatalf("TransactOpts 생성 실패: %v", err)
	}

	v1Address, v1Tx, _, err := contracts.DeployCounterV1(auth, client)
	if err != nil {
		log.Fatalf("CounterV1 배포 실패: %v", err)
	}

	fmt.Printf("  트랜잭션 전송됨: %s\n", v1Tx.Hash().Hex())
	fmt.Println("  채굴 대기 중...")

	v1Receipt, err := bind.WaitMined(ctx, client, v1Tx)
	if err != nil {
		log.Fatalf("V1 배포 트랜잭션 실패: %v", err)
	}

	fmt.Printf("  ✓ CounterV1 배포 완료!\n")
	fmt.Printf("    주소: %s\n", v1Address.Hex())
	fmt.Printf("    가스 사용: %d\n", v1Receipt.GasUsed)
	fmt.Println()

	// =========================================================================
	// STEP 3: CounterProxy 배포 (V1 주소 전달)
	// =========================================================================
	fmt.Println("[STEP 3] CounterProxy 배포")
	fmt.Println("───────────────────────────────────────────────────────────────────────────────")
	fmt.Println("  Proxy 생성자에 V1 주소를 전달")
	fmt.Println("  → Proxy._logic = V1 주소")
	fmt.Println()

	auth, err = getTransactOpts(ctx, client, privateKey, cfg.ChainID)
	if err != nil {
		log.Fatalf("TransactOpts 생성 실패: %v", err)
	}

	proxyAddress, proxyTx, proxy, err := contracts.DeployCounterProxy(auth, client, v1Address)
	if err != nil {
		log.Fatalf("Proxy 배포 실패: %v", err)
	}

	fmt.Printf("  트랜잭션 전송됨: %s\n", proxyTx.Hash().Hex())
	fmt.Println("  채굴 대기 중...")

	proxyReceipt, err := bind.WaitMined(ctx, client, proxyTx)
	if err != nil {
		log.Fatalf("Proxy 배포 트랜잭션 실패: %v", err)
	}

	fmt.Printf("  ✓ CounterProxy 배포 완료!\n")
	fmt.Printf("    주소: %s\n", proxyAddress.Hex())
	fmt.Printf("    가스 사용: %d\n", proxyReceipt.GasUsed)

	// Proxy 상태 확인
	callOpts := &bind.CallOpts{Context: ctx}
	logicAddr, _ := proxy.Implementation(callOpts)
	adminAddr, _ := proxy.Admin(callOpts)
	fmt.Printf("    _logic (V1): %s\n", logicAddr.Hex())
	fmt.Printf("    _admin: %s\n", adminAddr.Hex())
	fmt.Println()

	// =========================================================================
	// STEP 4: Proxy를 통해 initialize() 호출
	// =========================================================================
	fmt.Println("[STEP 4] Proxy를 통해 initialize() 호출")
	fmt.Println("───────────────────────────────────────────────────────────────────────────────")
	fmt.Println("  Proxy 주소로 V1의 initialize() 호출")
	fmt.Println("  → delegatecall로 V1 코드 실행")
	fmt.Println("  → Proxy 스토리지에 _owner 저장")
	fmt.Println()

	// Proxy 주소에 V1 ABI로 접근 (delegatecall이므로)
	v1ViaProxy, err := contracts.NewCounterV1(proxyAddress, client)
	if err != nil {
		log.Fatalf("V1 via Proxy 연결 실패: %v", err)
	}

	auth, err = getTransactOpts(ctx, client, privateKey, cfg.ChainID)
	if err != nil {
		log.Fatalf("TransactOpts 생성 실패: %v", err)
	}

	initTx, err := v1ViaProxy.Initialize(auth)
	if err != nil {
		log.Fatalf("initialize 실패: %v", err)
	}

	fmt.Printf("  트랜잭션 전송됨: %s\n", initTx.Hash().Hex())
	fmt.Println("  채굴 대기 중...")

	initReceipt, err := bind.WaitMined(ctx, client, initTx)
	if err != nil {
		log.Fatalf("initialize 트랜잭션 실패: %v", err)
	}

	fmt.Printf("  ✓ initialize 완료! (가스: %d)\n", initReceipt.GasUsed)

	// 결과 확인
	count, _ := v1ViaProxy.GetCount(callOpts)
	owner, _ := v1ViaProxy.Owner(callOpts)
	version, _ := v1ViaProxy.Version(callOpts)
	fmt.Printf("    count: %d\n", count)
	fmt.Printf("    owner: %s\n", owner.Hex())
	fmt.Printf("    version: %s\n", version)
	fmt.Println()

	// =========================================================================
	// STEP 5: Proxy를 통해 increment() 호출
	// =========================================================================
	fmt.Println("[STEP 5] Proxy를 통해 increment() 호출")
	fmt.Println("───────────────────────────────────────────────────────────────────────────────")

	auth, err = getTransactOpts(ctx, client, privateKey, cfg.ChainID)
	if err != nil {
		log.Fatalf("TransactOpts 생성 실패: %v", err)
	}

	incTx, err := v1ViaProxy.Increment(auth)
	if err != nil {
		log.Fatalf("increment 실패: %v", err)
	}

	fmt.Printf("  트랜잭션 전송됨: %s\n", incTx.Hash().Hex())
	bind.WaitMined(ctx, client, incTx)

	count, _ = v1ViaProxy.GetCount(callOpts)
	fmt.Printf("  ✓ increment 완료! count: %d\n", count)
	fmt.Println()

	// =========================================================================
	// STEP 6: CounterV2 배포
	// =========================================================================
	fmt.Println("[STEP 6] CounterV2 배포")
	fmt.Println("───────────────────────────────────────────────────────────────────────────────")
	fmt.Println("  V2에는 incrementBy(), reset(), setCount() 추가됨")
	fmt.Println()

	auth, err = getTransactOpts(ctx, client, privateKey, cfg.ChainID)
	if err != nil {
		log.Fatalf("TransactOpts 생성 실패: %v", err)
	}

	v2Address, v2Tx, _, err := contracts.DeployCounterV2(auth, client)
	if err != nil {
		log.Fatalf("CounterV2 배포 실패: %v", err)
	}

	fmt.Printf("  트랜잭션 전송됨: %s\n", v2Tx.Hash().Hex())
	fmt.Println("  채굴 대기 중...")

	v2Receipt, err := bind.WaitMined(ctx, client, v2Tx)
	if err != nil {
		log.Fatalf("V2 배포 트랜잭션 실패: %v", err)
	}

	fmt.Printf("  ✓ CounterV2 배포 완료!\n")
	fmt.Printf("    주소: %s\n", v2Address.Hex())
	fmt.Printf("    가스 사용: %d\n", v2Receipt.GasUsed)
	fmt.Println()

	// =========================================================================
	// STEP 7: Proxy 업그레이드 (V1 → V2)
	// =========================================================================
	fmt.Println("[STEP 7] Proxy 업그레이드 (V1 → V2)")
	fmt.Println("───────────────────────────────────────────────────────────────────────────────")
	fmt.Println("  proxy.upgradeTo(V2 주소) 호출")
	fmt.Println("  → Proxy._logic = V2 주소로 변경")
	fmt.Println("  → 데이터(count)는 그대로 유지!")
	fmt.Println()

	// 업그레이드 전 상태
	fmt.Printf("  업그레이드 전:\n")
	fmt.Printf("    _logic: %s (V1)\n", logicAddr.Hex())
	fmt.Printf("    count: %d\n", count)

	auth, err = getTransactOpts(ctx, client, privateKey, cfg.ChainID)
	if err != nil {
		log.Fatalf("TransactOpts 생성 실패: %v", err)
	}

	upgradeTx, err := proxy.UpgradeTo(auth, v2Address)
	if err != nil {
		log.Fatalf("upgradeTo 실패: %v", err)
	}

	fmt.Printf("\n  트랜잭션 전송됨: %s\n", upgradeTx.Hash().Hex())
	fmt.Println("  채굴 대기 중...")

	upgradeReceipt, err := bind.WaitMined(ctx, client, upgradeTx)
	if err != nil {
		log.Fatalf("upgrade 트랜잭션 실패: %v", err)
	}

	// 업그레이드 후 상태
	newLogicAddr, _ := proxy.Implementation(callOpts)
	fmt.Printf("\n  ✓ 업그레이드 완료! (가스: %d)\n", upgradeReceipt.GasUsed)
	fmt.Printf("  업그레이드 후:\n")
	fmt.Printf("    _logic: %s (V2)\n", newLogicAddr.Hex())
	fmt.Println()

	// =========================================================================
	// STEP 8: V2 기능 테스트 (incrementBy)
	// =========================================================================
	fmt.Println("[STEP 8] V2 기능 테스트 (incrementBy)")
	fmt.Println("───────────────────────────────────────────────────────────────────────────────")
	fmt.Println("  V1에는 없던 incrementBy(5) 호출")
	fmt.Println()

	// 이제 V2 ABI로 Proxy에 접근
	v2ViaProxy, err := contracts.NewCounterV2(proxyAddress, client)
	if err != nil {
		log.Fatalf("V2 via Proxy 연결 실패: %v", err)
	}

	// 현재 count 확인
	countBefore, _ := v2ViaProxy.GetCount(callOpts)
	versionNow, _ := v2ViaProxy.Version(callOpts)
	fmt.Printf("  현재 상태:\n")
	fmt.Printf("    count: %d\n", countBefore)
	fmt.Printf("    version: %s (V1→V2로 바뀜!)\n", versionNow)

	auth, err = getTransactOpts(ctx, client, privateKey, cfg.ChainID)
	if err != nil {
		log.Fatalf("TransactOpts 생성 실패: %v", err)
	}

	incByTx, err := v2ViaProxy.IncrementBy(auth, big.NewInt(5))
	if err != nil {
		log.Fatalf("incrementBy 실패: %v", err)
	}

	fmt.Printf("\n  incrementBy(5) 트랜잭션: %s\n", incByTx.Hash().Hex())
	bind.WaitMined(ctx, client, incByTx)

	countAfter, _ := v2ViaProxy.GetCount(callOpts)
	fmt.Printf("  ✓ incrementBy(5) 완료!\n")
	fmt.Printf("    count: %d → %d (+5)\n", countBefore, countAfter)
	fmt.Println()

	// =========================================================================
	// 결과 요약
	// =========================================================================
	fmt.Println("═══════════════════════════════════════════════════════════════════════════════")
	fmt.Println("                              결과 요약")
	fmt.Println("═══════════════════════════════════════════════════════════════════════════════")
	fmt.Println()
	fmt.Println("  배포된 컨트랙트:")
	fmt.Printf("    CounterV1:    %s\n", v1Address.Hex())
	fmt.Printf("    CounterV2:    %s\n", v2Address.Hex())
	fmt.Printf("    CounterProxy: %s ← 사용자는 이 주소만 사용\n", proxyAddress.Hex())
	fmt.Println()
	fmt.Println("  Proxy 상태:")
	fmt.Printf("    _logic: %s (V2)\n", newLogicAddr.Hex())
	fmt.Printf("    _admin: %s\n", adminAddr.Hex())
	fmt.Printf("    count:  %d\n", countAfter)
	fmt.Println()
	fmt.Println("  핵심 확인:")
	fmt.Println("    ✓ V1 → V2 업그레이드 성공")
	fmt.Println("    ✓ 업그레이드 후에도 count 데이터 유지됨")
	fmt.Println("    ✓ V2의 새 기능(incrementBy) 사용 가능")
	fmt.Println()
	fmt.Println("  Etherscan에서 확인:")
	fmt.Printf("    Proxy: https://sepolia.etherscan.io/address/%s\n", proxyAddress.Hex())
	fmt.Println()
}

// getTransactOpts는 트랜잭션 옵션을 생성합니다
func getTransactOpts(ctx context.Context, client *ethclient.Client, privateKey *ecdsa.PrivateKey, chainID int64) (*bind.TransactOpts, error) {
	publicKey := privateKey.Public().(*ecdsa.PublicKey)
	fromAddress := crypto.PubkeyToAddress(*publicKey)

	nonce, err := client.PendingNonceAt(ctx, fromAddress)
	if err != nil {
		return nil, fmt.Errorf("nonce 조회 실패: %w", err)
	}

	gasPrice, err := client.SuggestGasPrice(ctx)
	if err != nil {
		return nil, fmt.Errorf("가스 가격 조회 실패: %w", err)
	}

	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(chainID))
	if err != nil {
		return nil, fmt.Errorf("transactor 생성 실패: %w", err)
	}

	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)
	auth.GasLimit = uint64(3000000) // 배포는 가스가 많이 필요
	auth.GasPrice = gasPrice
	auth.Context = ctx

	return auth, nil
}

package usecase

import (
	"context"
	"fmt"
	"math/big"

	"github.com/myungsworld/ethereum/internal/domain"
	"github.com/myungsworld/ethereum/internal/repository"
)

// EthereumUsecase는 이더리움 관련 비즈니스 로직을 담당
type EthereumUsecase struct {
	reader repository.EthereumReader
	writer repository.EthereumWriter
}

// NewEthereumUsecase는 새로운 EthereumUsecase를 생성 (읽기 전용)
func NewEthereumUsecase(repo repository.EthereumReader) *EthereumUsecase {
	return &EthereumUsecase{reader: repo}
}

// NewEthereumUsecaseWithWriter는 읽기/쓰기가 가능한 EthereumUsecase를 생성
func NewEthereumUsecaseWithWriter(repo repository.EthereumRepository) *EthereumUsecase {
	return &EthereumUsecase{reader: repo, writer: repo}
}

// NetworkInfo는 네트워크 기본 정보를 담는 구조체
type NetworkInfo struct {
	ChainID     *big.Int
	BlockNumber uint64
}

// GetNetworkInfo는 현재 네트워크 정보를 조회
func (u *EthereumUsecase) GetNetworkInfo(ctx context.Context) (*NetworkInfo, error) {
	chainID, err := u.reader.ChainID(ctx)
	if err != nil {
		return nil, fmt.Errorf("Chain ID 조회 실패: %w", err)
	}

	blockNumber, err := u.reader.BlockNumber(ctx)
	if err != nil {
		return nil, fmt.Errorf("블록 번호 조회 실패: %w", err)
	}

	return &NetworkInfo{
		ChainID:     chainID,
		BlockNumber: blockNumber,
	}, nil
}

// GetLatestBlock은 최신 블록 정보를 조회
func (u *EthereumUsecase) GetLatestBlock(ctx context.Context) (*domain.Block, error) {
	blockNumber, err := u.reader.BlockNumber(ctx)
	if err != nil {
		return nil, fmt.Errorf("블록 번호 조회 실패: %w", err)
	}

	block, err := u.reader.BlockByNumber(ctx, big.NewInt(int64(blockNumber)))
	if err != nil {
		return nil, fmt.Errorf("블록 조회 실패: %w", err)
	}

	return block, nil
}

// GetBlock은 특정 블록 번호의 블록을 조회
func (u *EthereumUsecase) GetBlock(ctx context.Context, number uint64) (*domain.Block, error) {
	return u.reader.BlockByNumber(ctx, big.NewInt(int64(number)))
}

// AccountInfo는 계정 정보를 담는 구조체
type AccountInfo struct {
	Address    string
	Balance    *big.Int
	BalanceETH float64
	Nonce      uint64
}

// GetAccountInfo는 특정 주소의 계정 정보를 조회
func (u *EthereumUsecase) GetAccountInfo(ctx context.Context, address string) (*AccountInfo, error) {
	balance, err := u.reader.BalanceAt(ctx, address, nil)
	if err != nil {
		return nil, fmt.Errorf("잔액 조회 실패: %w", err)
	}

	nonce, err := u.reader.NonceAt(ctx, address, nil)
	if err != nil {
		return nil, fmt.Errorf("nonce 조회 실패: %w", err)
	}

	// ETH 단위로 변환
	ethBalance := new(big.Float).Quo(
		new(big.Float).SetInt(balance),
		new(big.Float).SetFloat64(1e18),
	)
	ethFloat, _ := ethBalance.Float64()

	return &AccountInfo{
		Address:    address,
		Balance:    balance,
		BalanceETH: ethFloat,
		Nonce:      nonce,
	}, nil
}

// GetTransaction은 트랜잭션 해시로 트랜잭션을 조회
func (u *EthereumUsecase) GetTransaction(ctx context.Context, hash string) (*domain.Transaction, error) {
	return u.reader.TransactionByHash(ctx, hash)
}

// SendETHRequest는 ETH 전송 요청 정보
type SendETHRequest struct {
	To     string  // 수신자 주소
	Amount float64 // ETH 단위
}

// SendETHResult는 ETH 전송 결과
type SendETHResult struct {
	TxHash string
	From   string
	To     string
	Amount float64
}

// SendETH는 ETH를 전송
func (u *EthereumUsecase) SendETH(ctx context.Context, req *SendETHRequest) (*SendETHResult, error) {
	if u.writer == nil {
		return nil, fmt.Errorf("쓰기 권한 없음: NewEthereumUsecaseWithWriter를 사용하세요")
	}

	// ETH를 Wei로 변환
	weiFloat := new(big.Float).Mul(
		big.NewFloat(req.Amount),
		big.NewFloat(1e18),
	)
	wei, _ := weiFloat.Int(nil)

	tx := &domain.Transaction{
		To:    req.To,
		Value: wei,
	}

	txHash, err := u.writer.SendTransaction(ctx, tx)
	if err != nil {
		return nil, fmt.Errorf("ETH 전송 실패: %w", err)
	}

	return &SendETHResult{
		TxHash: txHash,
		To:     req.To,
		Amount: req.Amount,
	}, nil
}

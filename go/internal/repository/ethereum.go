package repository

import (
	"context"
	"math/big"

	"github.com/myungsworld/ethereum/internal/domain"
)

// EthereumReader는 이더리움 블록체인 읽기 작업 인터페이스
type EthereumReader interface {
	// 네트워크 정보
	ChainID(ctx context.Context) (*big.Int, error)
	BlockNumber(ctx context.Context) (uint64, error)

	// 블록 조회
	BlockByNumber(ctx context.Context, number *big.Int) (*domain.Block, error)
	BlockByHash(ctx context.Context, hash string) (*domain.Block, error)

	// 트랜잭션 조회
	TransactionByHash(ctx context.Context, hash string) (*domain.Transaction, error)

	// 계정 조회
	BalanceAt(ctx context.Context, address string, blockNumber *big.Int) (*big.Int, error)
	NonceAt(ctx context.Context, address string, blockNumber *big.Int) (uint64, error)
}

// EthereumWriter는 이더리움 블록체인 쓰기 작업 인터페이스
type EthereumWriter interface {
	// 트랜잭션 전송
	SendTransaction(ctx context.Context, tx *domain.Transaction) (string, error)
}

// EthereumRepository는 이더리움 블록체인과의 모든 상호작용 인터페이스
type EthereumRepository interface {
	EthereumReader
	EthereumWriter
	Close()
}

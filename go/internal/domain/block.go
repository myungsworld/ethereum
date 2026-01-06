package domain

import (
	"math/big"
	"time"
)

// Block은 이더리움 블록을 나타내는 도메인 엔티티
type Block struct {
	Number       uint64
	Hash         string
	ParentHash   string
	Timestamp    time.Time
	Transactions []Transaction
	GasUsed      uint64
	GasLimit     uint64
	BaseFee      *big.Int
}

// TransactionCount는 블록 내 트랜잭션 수를 반환
func (b *Block) TransactionCount() int {
	return len(b.Transactions)
}

// BaseFeeGwei는 Base Fee를 Gwei 단위로 반환
func (b *Block) BaseFeeGwei() float64 {
	if b.BaseFee == nil {
		return 0
	}
	gwei := new(big.Float).Quo(
		new(big.Float).SetInt(b.BaseFee),
		new(big.Float).SetFloat64(1e9),
	)
	f, _ := gwei.Float64()
	return f
}

// GasUsagePercent는 가스 사용률을 백분율로 반환
func (b *Block) GasUsagePercent() float64 {
	if b.GasLimit == 0 {
		return 0
	}
	return float64(b.GasUsed) / float64(b.GasLimit) * 100
}

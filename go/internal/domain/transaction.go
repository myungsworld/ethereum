package domain

import "math/big"

// TxType은 트랜잭션 타입을 나타냄
type TxType uint8

const (
	TxTypeLegacy     TxType = 0 // Legacy 트랜잭션 (EIP-155)
	TxTypeAccessList TxType = 1 // Access List 트랜잭션 (EIP-2930)
	TxTypeDynamicFee TxType = 2 // Dynamic Fee 트랜잭션 (EIP-1559)
	TxTypeBlob       TxType = 3 // Blob 트랜잭션 (EIP-4844)
)

// Transaction은 이더리움 트랜잭션을 나타내는 도메인 엔티티
type Transaction struct {
	Hash     string
	From     string
	To       string // 빈 문자열이면 컨트랙트 생성
	Nonce    uint64
	Value    *big.Int
	GasLimit uint64
	GasPrice *big.Int
	Data     []byte
	Type     TxType
}

// ValueETH는 Value를 ETH 단위로 반환
func (t *Transaction) ValueETH() float64 {
	if t.Value == nil {
		return 0
	}
	eth := new(big.Float).Quo(
		new(big.Float).SetInt(t.Value),
		new(big.Float).SetFloat64(1e18),
	)
	f, _ := eth.Float64()
	return f
}

// GasPriceGwei는 Gas Price를 Gwei 단위로 반환
func (t *Transaction) GasPriceGwei() float64 {
	if t.GasPrice == nil {
		return 0
	}
	gwei := new(big.Float).Quo(
		new(big.Float).SetInt(t.GasPrice),
		new(big.Float).SetFloat64(1e9),
	)
	f, _ := gwei.Float64()
	return f
}

// IsContractCreation은 컨트랙트 생성 트랜잭션인지 확인
func (t *Transaction) IsContractCreation() bool {
	return t.To == ""
}

// TypeString은 트랜잭션 타입을 문자열로 반환
func (t *Transaction) TypeString() string {
	switch t.Type {
	case TxTypeLegacy:
		return "Legacy"
	case TxTypeAccessList:
		return "AccessList"
	case TxTypeDynamicFee:
		return "EIP-1559"
	case TxTypeBlob:
		return "Blob"
	default:
		return "Unknown"
	}
}

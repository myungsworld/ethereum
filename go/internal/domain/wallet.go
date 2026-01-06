package domain

import "math/big"

// Wallet은 이더리움 지갑을 나타내는 도메인 엔티티
type Wallet struct {
	Address    string
	PrivateKey string // 실제 운영에서는 절대 평문 저장하지 않음
	PublicKey  string
}

// Account는 이더리움 계정 상태를 나타냄
type Account struct {
	Address string
	Balance *big.Int
	Nonce   uint64
}

// BalanceETH는 잔액을 ETH 단위로 반환
func (a *Account) BalanceETH() float64 {
	if a.Balance == nil {
		return 0
	}
	eth := new(big.Float).Quo(
		new(big.Float).SetInt(a.Balance),
		new(big.Float).SetFloat64(1e18),
	)
	f, _ := eth.Float64()
	return f
}

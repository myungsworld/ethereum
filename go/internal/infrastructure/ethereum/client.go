package ethereum

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"

	"github.com/myungsworld/ethereum/internal/domain"
	"github.com/myungsworld/ethereum/internal/repository"
)

// Client는 go-ethereum ethclient를 래핑한 구현체
type Client struct {
	client     *ethclient.Client
	privateKey *ecdsa.PrivateKey
	address    common.Address
	chainID    *big.Int
}

// 컴파일 타임에 인터페이스 구현 확인
var _ repository.EthereumRepository = (*Client)(nil)

// NewClient는 새로운 이더리움 클라이언트를 생성
func NewClient(rpcURL string) (*Client, error) {
	client, err := ethclient.Dial(rpcURL)
	if err != nil {
		return nil, fmt.Errorf("이더리움 연결 실패: %w", err)
	}
	return &Client{client: client}, nil
}

// NewClientWithWallet는 지갑이 포함된 클라이언트를 생성
func NewClientWithWallet(rpcURL string, privateKeyHex string, chainID int64) (*Client, error) {
	client, err := ethclient.Dial(rpcURL)
	if err != nil {
		return nil, fmt.Errorf("이더리움 연결 실패: %w", err)
	}

	// 개인키 파싱
	privateKey, err := crypto.HexToECDSA(privateKeyHex)
	if err != nil {
		return nil, fmt.Errorf("개인키 파싱 실패: %w", err)
	}

	// 공개키에서 주소 파생
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return nil, fmt.Errorf("공개키 변환 실패")
	}
	address := crypto.PubkeyToAddress(*publicKeyECDSA)

	return &Client{
		client:     client,
		privateKey: privateKey,
		address:    address,
		chainID:    big.NewInt(chainID),
	}, nil
}

// Address는 클라이언트의 지갑 주소를 반환
func (c *Client) Address() string {
	return c.address.Hex()
}

// HasWallet은 지갑이 설정되어 있는지 확인
func (c *Client) HasWallet() bool {
	return c.privateKey != nil
}

// Close는 클라이언트 연결을 종료
func (c *Client) Close() {
	c.client.Close()
}

// ChainID는 현재 연결된 체인의 ID를 반환
func (c *Client) ChainID(ctx context.Context) (*big.Int, error) {
	return c.client.ChainID(ctx)
}

// BlockNumber는 최신 블록 번호를 반환
func (c *Client) BlockNumber(ctx context.Context) (uint64, error) {
	return c.client.BlockNumber(ctx)
}

// BlockByNumber는 블록 번호로 블록을 조회
func (c *Client) BlockByNumber(ctx context.Context, number *big.Int) (*domain.Block, error) {
	block, err := c.client.BlockByNumber(ctx, number)
	if err != nil {
		return nil, fmt.Errorf("블록 조회 실패: %w", err)
	}

	// 도메인 모델로 변환
	txs := make([]domain.Transaction, 0, len(block.Transactions()))
	for _, tx := range block.Transactions() {
		to := ""
		if tx.To() != nil {
			to = tx.To().Hex()
		}

		// From 주소 복구
		from := ""
		sender, err := c.client.TransactionSender(ctx, tx, block.Hash(), 0)
		if err == nil {
			from = sender.Hex()
		}

		txs = append(txs, domain.Transaction{
			Hash:     tx.Hash().Hex(),
			From:     from,
			To:       to,
			Nonce:    tx.Nonce(),
			Value:    tx.Value(),
			GasLimit: tx.Gas(),
			GasPrice: tx.GasPrice(),
			Data:     tx.Data(),
			Type:     domain.TxType(tx.Type()),
		})
	}

	return &domain.Block{
		Number:       block.Number().Uint64(),
		Hash:         block.Hash().Hex(),
		ParentHash:   block.ParentHash().Hex(),
		Timestamp:    time.Unix(int64(block.Time()), 0),
		Transactions: txs,
		GasUsed:      block.GasUsed(),
		GasLimit:     block.GasLimit(),
		BaseFee:      block.BaseFee(),
	}, nil
}

// BlockByHash는 블록 해시로 블록을 조회
func (c *Client) BlockByHash(ctx context.Context, hash string) (*domain.Block, error) {
	block, err := c.client.BlockByHash(ctx, common.HexToHash(hash))
	if err != nil {
		return nil, fmt.Errorf("블록 조회 실패: %w", err)
	}

	return &domain.Block{
		Number:     block.Number().Uint64(),
		Hash:       block.Hash().Hex(),
		ParentHash: block.ParentHash().Hex(),
		Timestamp:  time.Unix(int64(block.Time()), 0),
		GasUsed:    block.GasUsed(),
		GasLimit:   block.GasLimit(),
		BaseFee:    block.BaseFee(),
	}, nil
}

// TransactionByHash는 트랜잭션 해시로 트랜잭션을 조회
func (c *Client) TransactionByHash(ctx context.Context, hash string) (*domain.Transaction, error) {
	tx, _, err := c.client.TransactionByHash(ctx, common.HexToHash(hash))
	if err != nil {
		return nil, fmt.Errorf("트랜잭션 조회 실패: %w", err)
	}

	to := ""
	if tx.To() != nil {
		to = tx.To().Hex()
	}

	return &domain.Transaction{
		Hash:     tx.Hash().Hex(),
		To:       to,
		Nonce:    tx.Nonce(),
		Value:    tx.Value(),
		GasLimit: tx.Gas(),
		GasPrice: tx.GasPrice(),
		Data:     tx.Data(),
		Type:     domain.TxType(tx.Type()),
	}, nil
}

// BalanceAt는 특정 주소의 잔액을 조회
func (c *Client) BalanceAt(ctx context.Context, address string, blockNumber *big.Int) (*big.Int, error) {
	addr := common.HexToAddress(address)
	return c.client.BalanceAt(ctx, addr, blockNumber)
}

// NonceAt는 특정 주소의 nonce를 조회
func (c *Client) NonceAt(ctx context.Context, address string, blockNumber *big.Int) (uint64, error) {
	addr := common.HexToAddress(address)
	return c.client.NonceAt(ctx, addr, blockNumber)
}

// SendTransaction은 트랜잭션을 전송
func (c *Client) SendTransaction(ctx context.Context, tx *domain.Transaction) (string, error) {
	if c.privateKey == nil {
		return "", fmt.Errorf("지갑이 설정되지 않음: NewClientWithWallet을 사용하세요")
	}

	toAddress := common.HexToAddress(tx.To)

	// nonce 조회
	nonce, err := c.client.PendingNonceAt(ctx, c.address)
	if err != nil {
		return "", fmt.Errorf("nonce 조회 실패: %w", err)
	}

	// 가스 가격 조회
	gasPrice, err := c.client.SuggestGasPrice(ctx)
	if err != nil {
		return "", fmt.Errorf("가스 가격 조회 실패: %w", err)
	}

	// 가스 리밋 (ETH 전송은 21000)
	gasLimit := uint64(21000)
	if len(tx.Data) > 0 {
		gasLimit = tx.GasLimit
	}

	// 트랜잭션 생성
	ethTx := types.NewTransaction(
		nonce,
		toAddress,
		tx.Value,
		gasLimit,
		gasPrice,
		tx.Data,
	)

	// 트랜잭션 서명
	signedTx, err := types.SignTx(ethTx, types.NewEIP155Signer(c.chainID), c.privateKey)
	if err != nil {
		return "", fmt.Errorf("트랜잭션 서명 실패: %w", err)
	}

	// 트랜잭션 전송
	err = c.client.SendTransaction(ctx, signedTx)
	if err != nil {
		return "", fmt.Errorf("트랜잭션 전송 실패: %w", err)
	}

	return signedTx.Hash().Hex(), nil
}

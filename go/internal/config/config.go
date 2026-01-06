package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Environment string

const (
	Development Environment = "development"
	Production  Environment = "production"
)

type Config struct {
	Env        Environment
	RpcURL     string
	ChainID    int64
	PrivateKey string
}

// Load는 .env 파일을 읽어 환경에 맞는 설정을 반환
func Load() (*Config, error) {
	// .env 파일 로드 (없어도 에러 아님 - 환경변수로 대체 가능)
	_ = godotenv.Load()

	env := getEnv("ENV", "development")

	var rpcURL string
	var chainID int64

	switch Environment(env) {
	case Production:
		rpcURL = getEnv("RPC_URL_PRODUCTION", "https://eth.llamarpc.com")
		chainID = 1
	default:
		rpcURL = getEnv("RPC_URL_DEVELOPMENT", "https://ethereum-sepolia.publicnode.com")
		chainID = 11155111
	}

	privateKey := getEnv("PRIVATE_KEY", "")

	return &Config{
		Env:        Environment(env),
		RpcURL:     rpcURL,
		ChainID:    chainID,
		PrivateKey: privateKey,
	}, nil
}

// getEnv는 환경변수를 가져오고, 없으면 기본값 반환
func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

// IsDevelopment는 개발 환경인지 확인
func (c *Config) IsDevelopment() bool {
	return c.Env == Development
}

// IsProduction는 프로덕션 환경인지 확인
func (c *Config) IsProduction() bool {
	return c.Env == Production
}

// String은 설정 정보를 문자열로 반환
func (c *Config) String() string {
	return fmt.Sprintf("Env: %s, RPC: %s, ChainID: %d", c.Env, c.RpcURL, c.ChainID)
}

// HasPrivateKey는 개인키가 설정되어 있는지 확인
func (c *Config) HasPrivateKey() bool {
	return c.PrivateKey != "" && c.PrivateKey != "your_private_key_here"
}

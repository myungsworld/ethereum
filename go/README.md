# Go Ethereum Client

Go 언어를 사용한 이더리움 클라이언트 (클린 아키텍처)

## 구조

```
go/
├── .env                                # 환경 설정
├── .env.example                        # 환경 설정 예시
├── cmd/
│   └── main.go                         # 진입점
│
└── internal/
    ├── config/                         # 환경 설정 로드
    │   └── config.go
    │
    ├── domain/                         # 엔티티 (비즈니스 핵심)
    │   ├── block.go
    │   ├── transaction.go
    │   └── wallet.go
    │
    ├── repository/                     # 인터페이스 정의
    │   └── ethereum.go
    │
    ├── usecase/                        # 비즈니스 로직
    │   └── ethereum.go
    │
    └── infrastructure/
        ├── ethereum/                   # ethclient 구현
        │   └── client.go
        └── contracts/                  # 컨트랙트 Go 바인딩 (abigen)
```

## 환경 설정

```bash
cp .env.example .env
```

`.env` 파일:
```bash
ENV=development                          # development | production
PRIVATE_KEY=your_private_key_here        # 개인키 (0x 없이)
```

| 환경 | 네트워크 | Chain ID |
|------|---------|----------|
| development | Sepolia 테스트넷 | 11155111 |
| production | 이더리움 메인넷 | 1 |

## 실행

```bash
go run cmd/main.go
```

## 기능

- 블록/트랜잭션 조회
- 계정 잔액 조회
- ETH 전송
- 스마트 컨트랙트 배포 및 호출

## 아키텍처

```
┌─────────────────────────────────────────────────────┐
│                      main.go                        │
└─────────────────────────────────────────────────────┘
                         │
                         ▼
┌─────────────────────────────────────────────────────┐
│                     Usecase                         │
│              (비즈니스 로직)                          │
└─────────────────────────────────────────────────────┘
                         │
                         ▼
┌─────────────────────────────────────────────────────┐
│              Repository (인터페이스)                  │
└─────────────────────────────────────────────────────┘
                         ▲
                         │
┌─────────────────────────────────────────────────────┐
│                 Infrastructure                       │
│           (ethclient, contracts)                    │
└─────────────────────────────────────────────────────┘
                         │
                         ▼
┌─────────────────────────────────────────────────────┐
│                     Domain                          │
│          (Block, Transaction, Wallet)               │
└─────────────────────────────────────────────────────┘
```

의존성은 항상 안쪽(Domain)으로만 향합니다.

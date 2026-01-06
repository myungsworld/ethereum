# Ethereum Development

블록체인 개발자를 위한 학습 & 프로젝트 저장소

## 프로젝트 구조

```
ethereum/
├── docs/
│   └── ROADMAP.md              # 학습 로드맵
│
├── go/                         # Go 애플리케이션 (클린 아키텍처)
│   ├── .env                    # 환경 설정
│   ├── cmd/                    # 진입점
│   └── internal/
│       ├── config/             # 환경 설정 로드
│       ├── domain/             # 비즈니스 엔티티
│       ├── repository/         # 인터페이스 정의
│       ├── usecase/            # 비즈니스 로직
│       └── infrastructure/
│           ├── ethereum/       # ethclient 구현
│           └── contracts/      # 컨트랙트 Go 바인딩
│
├── solidity/                   # Solidity 스마트 컨트랙트
│   ├── contracts/              # 컨트랙트 소스 (.sol)
│   ├── build/                  # 컴파일 결과 (ABI, BIN)
│   └── test/                   # 테스트
│
├── scripts/                    # 빌드/배포 스크립트
│
├── projects/                   # 종합 프로젝트 (예정)
├── rust/                       # Rust 프로젝트 (예정)
└── typescript/                 # TypeScript 프로젝트 (예정)
```

## 기술 스택

| 도구 | 버전 | 용도 |
|------|------|------|
| Go | 1.25+ | 애플리케이션 개발 |
| go-ethereum | 1.16+ | 이더리움 클라이언트 |
| solc | 0.8.33 | Solidity 컴파일러 |
| abigen | 1.16+ | Go 바인딩 생성 |

## 시작하기

### 1. 환경 설정

```bash
cd go
cp .env.example .env
# .env 파일에서 PRIVATE_KEY 설정
```

### 2. Go 애플리케이션 실행

```bash
cd go
go mod download
go run cmd/main.go
```

### 3. 스마트 컨트랙트 컴파일

```bash
./scripts/compile.sh
```

## 환경 설정

`.env` 파일에서 환경을 전환합니다:

```bash
# development: Sepolia 테스트넷
# production: 이더리움 메인넷
ENV=development
```

## 학습 로드맵

자세한 내용: [docs/ROADMAP.md](docs/ROADMAP.md)

| Phase | 주제 | 상태 |
|-------|------|------|
| 1 | 이더리움 기초 (블록, 트랜잭션, 지갑) | 진행중 |
| 2 | 스마트 컨트랙트 (Solidity, 배포, 호출) | 진행중 |
| 3 | go-ethereum 심화 | 예정 |
| 4 | 실무 (인덱서, 이벤트 리스너) | 예정 |
| 5 | L2, MEV | 예정 |
| 6 | 대체 체인 (Solana, Aptos) | 예정 |

## 참고 자료

- [Ethereum Docs](https://ethereum.org/developers)
- [go-ethereum](https://github.com/ethereum/go-ethereum)
- [Solidity Docs](https://docs.soliditylang.org/)
- [Ethereum Development with Go](https://goethereumbook.org/)

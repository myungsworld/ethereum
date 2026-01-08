# 진행 상황 (Claude 세션용)

> **새 세션 시작 시 이 파일을 먼저 읽어주세요**

## 현재 상태

**마지막 업데이트**: 2026-01-06

### 완료된 작업

| 항목 | 상태 | 설명 |
|------|------|------|
| 클린 아키텍처 구조 | ✅ 완료 | domain, repository, usecase, infrastructure 분리 |
| 환경 설정 (.env) | ✅ 완료 | development(Sepolia) / production(Mainnet) 전환 |
| ETH 전송 기능 | ✅ 완료 | SendTransaction 구현 |
| Counter 컨트랙트 | ✅ 완료 | 배포됨: `0x6CE93619B4A2C81ABf0FB86701493D2d6903B8D9` (Sepolia) |
| 컨트랙트 상호작용 | ✅ 완료 | cmd/interact/main.go (상세 디버깅 포함) |

### 완료된 작업 (최근)

**Proxy 패턴 구현** - ✅ 100% 완료

작성된 파일:
- `solidity/contracts/proxy/CounterStorage.sol` ✅
- `solidity/contracts/proxy/CounterV1.sol` ✅
- `solidity/contracts/proxy/CounterV2.sol` ✅
- `solidity/contracts/proxy/CounterProxy.sol` ✅
- `go/cmd/proxy/main.go` ✅ (배포 및 테스트 코드)
- `go/deployments/sepolia.json` ✅ (배포 정보 기록)

완료된 작업:
1. ✅ CounterV2.sol 작성 (incrementBy, reset, setCount 추가)
2. ✅ CounterProxy.sol 작성 (delegatecall, upgradeTo 구현)
3. ✅ 컴파일 및 Go 바인딩 생성
4. ✅ Sepolia 테스트넷 배포 및 업그레이드 테스트

### 진행 중인 작업

---

## 프로젝트 구조

```
ethereum/
├── docs/
│   ├── ROADMAP.md          # 학습 로드맵
│   └── PROGRESS.md         # 이 파일 (세션 진행 상황)
│
├── go/                     # Go 애플리케이션
│   ├── .env                # 환경 설정 (PRIVATE_KEY 포함)
│   ├── cmd/
│   │   ├── main.go         # 기본 예제
│   │   ├── deploy/         # 컨트랙트 배포
│   │   └── interact/       # 컨트랙트 상호작용
│   └── internal/
│       ├── config/         # 환경 설정 로드
│       ├── domain/         # 엔티티
│       ├── repository/     # 인터페이스
│       ├── usecase/        # 비즈니스 로직
│       └── infrastructure/
│           ├── ethereum/   # ethclient 구현
│           └── contracts/  # Go 바인딩 (abigen)
│
├── solidity/               # Solidity 컨트랙트
│   ├── contracts/
│   │   ├── Counter.sol     # 기본 카운터 (배포됨)
│   │   └── proxy/          # Proxy 패턴 (진행 중)
│   └── build/              # 컴파일 결과
│
└── scripts/
    └── compile.sh          # 컴파일 스크립트
```

---

## 주요 정보

### 지갑
- **주소**: `0xdFcC6De7FC6ebfcbE0f465A2fC14917e74FCB5f6`
- **네트워크**: Sepolia (Chain ID: 11155111)
- **개인키**: `.env` 파일에 저장됨

### 배포된 컨트랙트
| 컨트랙트 | 주소 | 네트워크 |
|----------|------|---------|
| Counter | `0x6CE93619B4A2C81ABf0FB86701493D2d6903B8D9` | Sepolia |
| CounterV1 | `0x401e2e87Cd6B0D419Fc40f6192D6Ae1F8f9E53a3` | Sepolia |
| CounterV2 | `0xB032FaaB7504d2143A766f5CD562895175448000` | Sepolia |
| **CounterProxy** | `0xE61c5e06B5656Eb618011140807b21ffE9d6D4E4` | Sepolia |

### 실행 방법
```bash
cd /Users/myungsworld/go/src/ethereum/go

# 기본 실행
go run cmd/main.go

# 컨트랙트 배포
go run cmd/deploy/main.go

# 컨트랙트 상호작용
go run cmd/interact/main.go

# 컨트랙트 컴파일
../scripts/compile.sh
```

---

## 학습 노트

### 이해한 개념
1. **ABI**: 이더리움 표준, 언어 중립적 인터페이스
2. **컨트랙트 배포**: Bytecode를 블록체인에 저장하는 트랜잭션
3. **Call vs Transaction**: 읽기(무료) vs 쓰기(가스 필요)
4. **Proxy 패턴**: 컨트랙트 업그레이드를 위한 우회 방법
5. **geth는 표준이 아님**: Go로 작성된 이더리움 구현체일 뿐

### 다음에 학습할 것
- Proxy 패턴 구현 완료
- delegatecall 동작 원리
- 스토리지 슬롯 충돌 문제

---

## 다음 세션에서 할 일

```
1. CounterV2.sol 작성 완료
2. CounterProxy.sol 작성
3. 컴파일 및 테스트
4. (선택) 이벤트 구독 기능
```

# Blockchain Developer Roadmap

백엔드 개발자를 위한 블록체인 개발자 전환 로드맵

---

## 개요

```
Phase 1-3: 이더리움 & EVM 기초
Phase 4-5: 실무 & 심화
Phase 6: 대체 체인 (Solana, Aptos 등)
```

---

## Phase 1: 이더리움 기초 (Foundation)

**언어:** Go
**위치:** `go/`

### 1.1 이더리움 핵심 개념

**이해해야 할 것들:**
- **Account**: EOA(Externally Owned Account) vs Contract Account
- **Transaction**: 트랜잭션 구조 (nonce, gasPrice, gasLimit, to, value, data)
- **Block**: 블록 구조, 블록해시, 부모해시
- **State**: World State, Account State, Storage
- **Gas**: 가스 개념, EIP-1559 (baseFee, maxFeePerGas, maxPriorityFeePerGas)

**실습:**
- [ ] ethclient로 메인넷 연결
- [ ] 최신 블록 조회
- [ ] 특정 주소 잔액 조회
- [ ] 트랜잭션 조회

### 1.2 지갑 (Wallet)

**이해해야 할 것들:**
- 개인키(Private Key) → 공개키(Public Key) → 주소(Address) 파생 과정
- ECDSA (secp256k1) 암호화
- Keystore 파일 구조
- HD Wallet (BIP-32/39/44)

**실습:**
- [ ] 새 지갑 생성
- [ ] Keystore 파일로 저장/불러오기
- [ ] 니모닉으로 지갑 복구

### 1.3 트랜잭션 (Transaction)

**이해해야 할 것들:**
- Legacy Transaction vs EIP-1559 Transaction
- 트랜잭션 서명 과정
- RLP 인코딩
- 트랜잭션 해시 계산

**실습:**
- [ ] ETH 전송 트랜잭션 생성
- [ ] 트랜잭션 서명
- [ ] 트랜잭션 전송 (테스트넷)
- [ ] 트랜잭션 영수증 조회

---

## Phase 2: 스마트 컨트랙트 (Smart Contracts)

**언어:** Solidity
**위치:** `solidity/`

### 2.1 Solidity 기초

**이해해야 할 것들:**
- 데이터 타입 (uint, address, mapping, struct)
- 함수 가시성 (public, private, external, internal)
- Storage vs Memory vs Calldata
- 이벤트 (Events)
- Modifier

**실습:**
- [ ] 간단한 Counter 컨트랙트 작성
- [ ] ERC-20 토큰 컨트랙트 분석
- [ ] 테스트 작성 (Foundry)

### 2.2 컨트랙트 개발 환경

**도구:**
- **Foundry** - 현대적 Solidity 개발 환경 (추천)
- `forge` - 빌드/테스트
- `cast` - CLI 상호작용
- `anvil` - 로컬 노드

**실습:**
- [ ] Foundry 프로젝트 초기화
- [ ] 컨트랙트 컴파일 & 테스트
- [ ] 테스트넷 배포
- [ ] Go에서 컨트랙트 호출 (abigen)

### 2.3 이벤트 & 로그

**이해해야 할 것들:**
- 이벤트 구조 (topics, data)
- 로그 필터링
- 블룸 필터 (Bloom Filter)

**실습:**
- [ ] 특정 컨트랙트의 이벤트 구독
- [ ] 과거 이벤트 조회 (eth_getLogs)
- [ ] 이벤트 파싱

### 2.4 DeFi 기초

**이해해야 할 것들:**
- ERC-20, ERC-721, ERC-1155
- Uniswap V2/V3 구조
- Lending Protocol (Aave, Compound)
- Oracle (Chainlink)

**실습:**
- [ ] ERC-20 토큰 발행
- [ ] Uniswap에서 스왑 실행
- [ ] Flash Loan 이해

---

## Phase 3: go-ethereum 심화 (Deep Dive)

**언어:** Go
**위치:** `go/`

### 3.1 geth 소스코드 분석

**핵심 패키지:**
```
go-ethereum/
├── core/           # 블록체인 핵심 로직
│   ├── types/      # Block, Transaction, Receipt 등 타입 정의
│   ├── state/      # State 관리
│   ├── vm/         # EVM 구현
│   └── rawdb/      # 저수준 DB 접근
├── consensus/      # 합의 알고리즘 (PoS/PoW)
├── trie/           # Merkle Patricia Trie 구현
├── p2p/            # P2P 네트워크 레이어
├── eth/            # 이더리움 프로토콜
└── accounts/       # 계정 관리
```

**학습 순서:**
1. `core/types/` - 기본 타입 이해
2. `core/state/` - 상태 관리 방식
3. `trie/` - 데이터 구조
4. `core/vm/` - EVM 실행 흐름
5. `consensus/` - 블록 검증

### 3.2 로컬 노드 운영

**실습:**
- [ ] geth 빌드
- [ ] 개발용 로컬 체인 실행
- [ ] IPC/RPC 연결
- [ ] 동기화 모드 이해 (full, snap, light)

### 3.3 데이터 구조

**이해해야 할 것들:**
- Merkle Patricia Trie
- RLP (Recursive Length Prefix) 인코딩
- State Trie, Storage Trie, Receipt Trie

---

## Phase 4: 실무 역량 (Production Skills)

**언어:** Go, TypeScript
**위치:** `projects/`

### 4.1 블록 인덱서

**구현해볼 것:**
- [ ] 블록 순회 및 저장
- [ ] 트랜잭션 파싱 및 분류
- [ ] ERC-20 전송 추적
- [ ] 데이터베이스 설계 (PostgreSQL/ClickHouse)

**위치:** `projects/indexer/`

### 4.2 이벤트 리스너

**구현해볼 것:**
- [ ] 실시간 이벤트 구독 (WebSocket)
- [ ] 재시작 시 누락 이벤트 처리
- [ ] 리오그(Reorg) 대응

### 4.3 멀티체인 지원

**EVM 호환 체인:**
- Polygon
- Arbitrum
- Optimism
- Base

**차이점:**
- Chain ID
- 가스 계산 방식
- 블록 생성 시간
- L2 특수 트랜잭션 (Arbitrum Retryable, Optimism L1→L2 메시지)

### 4.4 가스 최적화

**알아야 할 것:**
- 가스 추정 (eth_estimateGas)
- EIP-1559 가스 가격 전략
- 트랜잭션 대기열 관리
- Flashbots / MEV 보호

---

## Phase 5: 심화 주제 (Advanced)

### 5.1 EIP 분석

**중요 EIP:**
- **EIP-1559**: 가스 수수료 개편
- **EIP-4844**: Proto-danksharding (Blob 트랜잭션)
- **EIP-4337**: Account Abstraction
- **EIP-712**: 타입화된 구조적 데이터 서명

### 5.2 Layer 2

**이해해야 할 것:**
- Optimistic Rollup vs ZK Rollup
- 챌린지 기간 (Optimistic)
- 증명 시스템 (ZK)
- L1 ↔ L2 브릿지

**실습:**
- [ ] Arbitrum/Optimism에 컨트랙트 배포
- [ ] L1→L2 메시지 전송

### 5.3 MEV (Maximal Extractable Value)

**이해해야 할 것:**
- 프론트러닝, 백러닝, 샌드위치
- Flashbots
- Block Builder / Searcher
- PBS (Proposer-Builder Separation)

---

## Phase 6: 대체 체인 (Alternative Chains)

### 6.1 Solana

**언어:** Rust
**위치:** `rust/solana/`

**특징:**
- 높은 TPS (65,000+)
- Proof of History (PoH)
- 병렬 트랜잭션 처리 (Sealevel)

**학습 순서:**
1. Rust 기초
2. Solana 아키텍처 이해
3. Anchor 프레임워크
4. Program 작성 & 배포

**실습:**
- [ ] Rust 기초 학습
- [ ] Solana CLI 사용
- [ ] 간단한 Program 작성
- [ ] SPL Token 발행

**자료:**
- [Solana Cookbook](https://solanacookbook.com/)
- [Anchor Book](https://book.anchor-lang.com/)

### 6.2 Aptos / Sui

**언어:** Move
**위치:** `move/`

**특징:**
- Move 언어 (자원 지향 프로그래밍)
- 병렬 실행 (Block-STM)
- 객체 중심 모델 (Sui)

**학습 순서:**
1. Move 언어 기초
2. Aptos/Sui 아키텍처
3. 모듈 작성 & 배포

**실습:**
- [ ] Move 기초 학습
- [ ] Aptos CLI 사용
- [ ] 간단한 모듈 배포

**자료:**
- [Aptos Learn](https://aptos.dev/en/build/get-started)
- [Move Book](https://move-book.com/)

### 6.3 Cosmos

**언어:** Go
**위치:** `go/cosmos/`

**특징:**
- Tendermint 합의
- IBC (Inter-Blockchain Communication)
- Cosmos SDK로 체인 구축

**실습:**
- [ ] Cosmos SDK 기초
- [ ] 간단한 체인 모듈 작성
- [ ] IBC 이해

---

## 학습 자료

### 이더리움 & EVM
- [Ethereum Docs](https://ethereum.org/developers/docs)
- [go-ethereum](https://github.com/ethereum/go-ethereum)
- [Ethereum Development with Go](https://goethereumbook.org/)
- [Foundry Book](https://book.getfoundry.sh/)
- [Solidity by Example](https://solidity-by-example.org/)

### DeFi
- [DeFi Developer Road Map](https://github.com/OffcierCia/DeFi-Developer-Road-Map)
- [Uniswap V2 Core](https://github.com/Uniswap/v2-core)
- [Uniswap V3 Book](https://uniswapv3book.com/)

### 심화
- [EIP Repository](https://eips.ethereum.org/)
- [Flashbots Docs](https://docs.flashbots.net/)
- [L2Beat](https://l2beat.com/)

### 대체 체인
- [Solana Docs](https://docs.solana.com/)
- [Aptos Docs](https://aptos.dev/)
- [Cosmos SDK Docs](https://docs.cosmos.network/)

---

## 체크리스트

### Phase 1: 이더리움 기초
- [ ] ethclient로 블록/트랜잭션 조회
- [ ] 지갑 생성 및 서명
- [ ] ETH 전송 트랜잭션 실행

### Phase 2: 스마트 컨트랙트
- [ ] Solidity 컨트랙트 작성
- [ ] Foundry로 테스트
- [ ] 테스트넷 배포

### Phase 3: go-ethereum 심화
- [ ] geth 소스코드 분석
- [ ] 로컬 노드 운영

### Phase 4: 실무
- [ ] 블록 인덱서 구현
- [ ] 이벤트 리스너 구현
- [ ] 멀티체인 지원

### Phase 5: 심화
- [ ] L2 컨트랙트 배포
- [ ] MEV 이해

### Phase 6: 대체 체인
- [ ] Solana Program 작성
- [ ] Move 모듈 작성

---

## 진행 상황

| Phase | 주제 | 상태 | 시작일 | 완료일 |
|-------|------|------|--------|--------|
| 1 | 이더리움 기초 | 진행 예정 | - | - |
| 2 | 스마트 컨트랙트 | - | - | - |
| 3 | go-ethereum 심화 | - | - | - |
| 4 | 실무 역량 | - | - | - |
| 5 | 심화 주제 | - | - | - |
| 6 | 대체 체인 | - | - | - |

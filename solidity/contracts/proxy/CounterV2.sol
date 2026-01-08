// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import "./CounterStorage.sol";

/// @title CounterV2
/// @notice 카운터 로직 버전 2 - V1에서 기능 추가
/// @dev V1과 동일한 스토리지 레이아웃을 유지해야 함 (CounterStorage 상속)
contract CounterV2 is CounterStorage {
    // =========================================================================
    // 이벤트 (Event)
    // =========================================================================
    // Go 비교: log.Printf() 또는 Kafka publish와 비슷
    // - 블록체인에 영구 기록됨
    // - 오프체인에서 구독 가능 (eth_subscribe)
    // - 상태 변수 저장보다 가스가 저렴함

    event CountChanged(uint256 newCount, address changedBy);
    event CountReset(address resetBy);
    event Initialized(address owner);

    // =========================================================================
    // 초기화 함수
    // =========================================================================
    // Proxy 패턴에서는 constructor 대신 initialize 사용
    // 이유: constructor는 배포 시 한 번만 실행되는데,
    //       Proxy는 Logic 컨트랙트의 코드만 빌려쓰므로 constructor가 실행 안 됨

    /// @notice 초기화 함수 (한 번만 호출 가능)
    /// @dev _owner == address(0)이면 아직 초기화 안 된 것
    function initialize() external {
        // require: 조건이 false면 트랜잭션 취소 (revert)
        // Go 비교: if _owner != address(0) { return errors.New("Already initialized") }
        require(_owner == address(0), "Already initialized");

        // msg.sender: 이 함수를 호출한 주소
        // Go 비교: ctx.Value("caller").(common.Address)
        _owner = msg.sender;
        _count = 0;

        // emit: 이벤트 발생 (로그 기록)
        emit Initialized(msg.sender);
    }

    // =========================================================================
    // 읽기 함수 (View/Pure)
    // =========================================================================
    // view: 상태를 읽기만 함 → 외부 호출 시 가스 무료
    // pure: 상태에 접근조차 안 함 → 외부 호출 시 가스 무료

    /// @notice 현재 카운트 값 반환
    /// @return 현재 카운트 (uint256)
    function getCount() external view returns (uint256) {
        // external: 외부에서만 호출 가능 (가스 효율적)
        // view: 상태 읽기만 함
        // returns (uint256): 반환 타입
        return _count;
    }

    /// @notice 소유자 주소 반환
    function owner() external view returns (address) {
        return _owner;
    }

    /// @notice 버전 문자열 반환
    /// @dev pure는 상태 변수를 전혀 사용 안 함
    function version() external pure returns (string memory) {
        // pure: _count나 _owner 같은 상태 변수를 안 씀
        // memory: 문자열은 메모리에 저장 (storage에 저장 안 함)
        return "V2";  // V1과 다른 점: 여기가 "V2"
    }

    // =========================================================================
    // 쓰기 함수 (상태 변경)
    // =========================================================================
    // view/pure가 없으면 상태 변경 가능 → 가스 필요

    /// @notice 카운트를 1 증가
    function increment() external {
        _count += 1;
        emit CountChanged(_count, msg.sender);
    }

    /// @notice 카운트를 1 감소
    function decrement() external {
        require(_count > 0, "Counter: cannot go below zero");
        _count -= 1;
        emit CountChanged(_count, msg.sender);
    }

    // =========================================================================
    // V2에서 추가된 새 기능들
    // =========================================================================

    /// @notice 카운트를 지정한 양만큼 증가
    /// @param amount 증가시킬 양
    /// @dev V1에는 없던 기능 - 업그레이드의 핵심!
    function incrementBy(uint256 amount) external {
        // uint256: Go의 *big.Int와 비슷 (256비트 부호 없는 정수)
        require(amount > 0, "Counter: amount must be greater than zero");
        _count += amount;
        emit CountChanged(_count, msg.sender);
    }

    /// @notice 카운트를 0으로 리셋 (소유자만 가능)
    /// @dev onlyOwner 패턴: msg.sender가 _owner인지 확인
    function reset() external {
        // 권한 체크: 소유자만 리셋 가능
        require(msg.sender == _owner, "Counter: only owner can reset");
        _count = 0;
        emit CountReset(msg.sender);
    }

    /// @notice 카운트를 특정 값으로 설정 (소유자만 가능)
    /// @param newCount 새로운 카운트 값
    function setCount(uint256 newCount) external {
        require(msg.sender == _owner, "Counter: only owner can set count");
        _count = newCount;
        emit CountChanged(_count, msg.sender);
    }
}

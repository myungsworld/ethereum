// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import "./CounterStorage.sol";

/// @title CounterProxy
/// @notice 업그레이드 가능한 Proxy 컨트랙트
/// @dev delegatecall을 사용해 Logic 컨트랙트의 코드를 실행하되,
///      스토리지는 이 Proxy에 저장됨
contract CounterProxy is CounterStorage {
    // =========================================================================
    // 스토리지 변수
    // =========================================================================
    // 중요: CounterStorage를 상속받았으므로 슬롯 0, 1은 이미 사용 중
    //   슬롯 0: _count (CounterStorage)
    //   슬롯 1: _owner (CounterStorage)
    //   슬롯 2: _logic (여기서 추가)
    //   슬롯 3: _admin (여기서 추가)
    //
    // 순서를 바꾸면 데이터가 꼬임! (스토리지 충돌)

    /// @dev Logic 컨트랙트 주소 (V1 또는 V2)
    address internal _logic;

    /// @dev Proxy 관리자 (업그레이드 권한)
    address internal _admin;

    // =========================================================================
    // 이벤트
    // =========================================================================

    /// @dev Logic 컨트랙트가 업그레이드됨
    event Upgraded(address indexed oldLogic, address indexed newLogic);

    /// @dev 관리자가 변경됨
    event AdminChanged(address indexed oldAdmin, address indexed newAdmin);

    // =========================================================================
    // 생성자
    // =========================================================================
    // Proxy는 constructor 사용 가능 (Logic과 달리 직접 배포되니까)

    /// @param logic_ 초기 Logic 컨트랙트 주소 (V1)
    constructor(address logic_) {
        require(logic_ != address(0), "Proxy: logic address cannot be zero");

        _logic = logic_;
        _admin = msg.sender;  // 배포자가 관리자

        emit Upgraded(address(0), logic_);
    }

    // =========================================================================
    // 관리자 함수
    // =========================================================================

    /// @notice 현재 Logic 컨트랙트 주소 조회
    function implementation() external view returns (address) {
        return _logic;
    }

    /// @notice 현재 관리자 주소 조회
    function admin() external view returns (address) {
        return _admin;
    }

    /// @notice Logic 컨트랙트 업그레이드 (관리자만)
    /// @param newLogic 새로운 Logic 컨트랙트 주소
    /// @dev 이게 업그레이드의 핵심! 주소만 바꾸면 새 코드 사용
    function upgradeTo(address newLogic) external {
        // 권한 체크
        require(msg.sender == _admin, "Proxy: only admin can upgrade");
        require(newLogic != address(0), "Proxy: new logic cannot be zero");
        require(newLogic != _logic, "Proxy: same logic address");

        address oldLogic = _logic;
        _logic = newLogic;

        emit Upgraded(oldLogic, newLogic);
    }

    /// @notice 관리자 변경 (현재 관리자만)
    /// @param newAdmin 새로운 관리자 주소
    function changeAdmin(address newAdmin) external {
        require(msg.sender == _admin, "Proxy: only admin can change admin");
        require(newAdmin != address(0), "Proxy: new admin cannot be zero");

        address oldAdmin = _admin;
        _admin = newAdmin;

        emit AdminChanged(oldAdmin, newAdmin);
    }

    // =========================================================================
    // fallback & receive
    // =========================================================================
    // 이 부분이 Proxy의 핵심!
    //
    // Go 비교:
    //   HTTP 라우터에서 매칭되는 핸들러가 없을 때 기본 핸들러로 가는 것과 비슷
    //   mux.HandleFunc("/", defaultHandler)
    //
    // 동작 원리:
    //   1. 사용자가 proxy.increment() 호출
    //   2. Proxy에 increment() 함수 없음
    //   3. fallback() 실행됨
    //   4. fallback()이 Logic으로 delegatecall
    //   5. Logic의 increment() 코드 실행, 스토리지는 Proxy 것 사용

    /// @dev ETH를 받을 때 실행 (msg.data가 비어있을 때)
    receive() external payable {
        // 단순 ETH 전송도 Logic으로 전달
        _delegate(_logic);
    }

    /// @dev 정의되지 않은 함수 호출 시 실행
    /// @notice 모든 함수 호출을 Logic 컨트랙트로 전달
    fallback() external payable {
        // 핵심: Logic으로 delegatecall
        _delegate(_logic);
    }

    // =========================================================================
    // 내부 함수: delegatecall 실행
    // =========================================================================

    /// @dev Logic 컨트랙트로 delegatecall 수행
    /// @param logic Logic 컨트랙트 주소
    function _delegate(address logic) internal {
        // assembly: 저수준 EVM 코드 (인라인 어셈블리)
        // Go 비교: unsafe 패키지나 cgo와 비슷한 느낌
        //
        // 왜 assembly를 쓰나?
        // - delegatecall의 결과를 그대로 반환해야 함
        // - Solidity 고수준 문법으로는 불가능
        // - 가스 효율성

        assembly {
            // ─────────────────────────────────────────────────────
            // 1. calldata 복사
            // ─────────────────────────────────────────────────────
            // calldatasize(): 입력 데이터 크기
            // calldatacopy(to, from, size): calldata를 메모리로 복사
            //
            // 예: increment() 호출 시
            //     calldata = 0xd09de08a (함수 selector 4바이트)

            calldatacopy(
                0,              // 메모리 위치 0에 복사
                0,              // calldata 시작 위치
                calldatasize()  // 전체 calldata 크기
            )

            // ─────────────────────────────────────────────────────
            // 2. delegatecall 실행
            // ─────────────────────────────────────────────────────
            // delegatecall(gas, to, inOffset, inSize, outOffset, outSize)
            //
            // 핵심 차이점:
            //   call:         Logic의 스토리지 사용
            //   delegatecall: Proxy의 스토리지 사용 (코드만 빌려옴)
            //
            // msg.sender도 원래 호출자 그대로 유지됨!

            let result := delegatecall(
                gas(),          // 남은 가스 전부 전달
                logic,          // Logic 컨트랙트 주소
                0,              // 입력 데이터 메모리 시작
                calldatasize(), // 입력 데이터 크기
                0,              // 출력 저장할 메모리 위치 (나중에 복사)
                0               // 출력 크기 (모름, 나중에 확인)
            )

            // ─────────────────────────────────────────────────────
            // 3. 반환 데이터 복사
            // ─────────────────────────────────────────────────────
            // returndatasize(): 반환 데이터 크기
            // returndatacopy(to, from, size): 반환 데이터를 메모리로 복사

            returndatacopy(
                0,                // 메모리 위치 0에 복사
                0,                // 반환 데이터 시작 위치
                returndatasize()  // 반환 데이터 크기
            )

            // ─────────────────────────────────────────────────────
            // 4. 결과에 따라 반환 또는 revert
            // ─────────────────────────────────────────────────────
            // result: 1이면 성공, 0이면 실패

            switch result
            case 0 {
                // 실패: revert하고 에러 메시지 반환
                revert(0, returndatasize())
            }
            default {
                // 성공: 결과 반환
                return(0, returndatasize())
            }
        }
    }
}

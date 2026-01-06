// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import "./CounterStorage.sol";

/// @title CounterV1
/// @notice 카운터 로직 버전 1 - 기본 기능만
contract CounterV1 is CounterStorage {
    // =============================================================
    // 이벤트
    // =============================================================
    event CountChanged(uint256 newCount, address changedBy);
    event Initialized(address owner);

    // =============================================================
    // 초기화 (constructor 대신 사용 - Proxy 패턴의 특징)
    // =============================================================

    /// @notice 초기화 함수 (한 번만 호출 가능)
    /// @dev Proxy에서는 constructor가 아닌 initialize 사용
    function initialize() external {
        require(_owner == address(0), "Already initialized");
        _owner = msg.sender;
        _count = 0;
        emit Initialized(msg.sender);
    }

    // =============================================================
    // 읽기 함수
    // =============================================================

    function getCount() external view returns (uint256) {
        return _count;
    }

    function owner() external view returns (address) {
        return _owner;
    }

    function version() external pure returns (string memory) {
        return "V1";
    }

    // =============================================================
    // 쓰기 함수
    // =============================================================

    function increment() external {
        _count += 1;
        emit CountChanged(_count, msg.sender);
    }

    function decrement() external {
        require(_count > 0, "Counter: cannot go below zero");
        _count -= 1;
        emit CountChanged(_count, msg.sender);
    }
}

// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

/// @title Counter
/// @notice 간단한 카운터 컨트랙트 - 학습용
contract Counter {
    uint256 private count;
    address public owner;

    event CountChanged(uint256 newCount, address changedBy);

    constructor(uint256 _initialCount) {
        count = _initialCount;
        owner = msg.sender;
    }

    /// @notice 현재 카운트 값을 반환
    function getCount() public view returns (uint256) {
        return count;
    }

    /// @notice 카운트를 1 증가
    function increment() public {
        count += 1;
        emit CountChanged(count, msg.sender);
    }

    /// @notice 카운트를 1 감소
    function decrement() public {
        require(count > 0, "Counter: cannot decrement below zero");
        count -= 1;
        emit CountChanged(count, msg.sender);
    }

    /// @notice 카운트를 특정 값으로 설정 (owner만 가능)
    function setCount(uint256 _newCount) public {
        require(msg.sender == owner, "Counter: only owner can set count");
        count = _newCount;
        emit CountChanged(count, msg.sender);
    }
}

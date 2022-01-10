# Ethereum

## [Go](https://github.com/myungsworld/ethereum/tree/main/go)

---

## [Solidity](https://github.com/myungsworld/ethereum/tree/main/solidity)

---

### Deploy

![스크린샷 2022-01-07 오후 1 21 11](https://user-images.githubusercontent.com/56465854/148492075-a0f15cd3-110d-4428-80d3-5048e390bde6.png)

- 솔리디티 스마트 컨트랙트를 작성한 후 [Solidity Compiler](https://github.com/ethereum/solidity) 를 통한 .bin 과 .abi 파일 생성
- 바이너리 파일 (.bin) 은 어더리움 네트워크(EVM) 에 배포
- abi 는 (API 와 비슷한 개념) abigen 을 통해 고 파일로 변경이 가능함

#### solc and abigen

```shell
solc --bin --abi contract.sol -o build
abigen --bin=build/contract.bin --abi=build/contract.abi --pkg=contract --out=gen/contract.go
```

- solc 명령어로 bin 과 abi 파일 빌드 후
- abigen으로 고언어로 핸들링 할 수 있는 contract.go 파일 생성
- 오너의 월렛주소, 프라이빗키 , 네트워크 및 컨트랙트 정보 (chainID , nonce , GasPrice , GasLimit) 등을 가지고
- 선택한 네트워크에 스마트 컨트랙트를 배포한다.
- [abigen을 통해서 생성된 go 파일 살펴보기](https://github.com/myungsworld/ethereum/blob/main/go/gen/todo.go)

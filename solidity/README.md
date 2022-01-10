# Solidity

## 개발 환경 구축

### [Remix IDE](https://remix.ethereum.org)
- Remix 는 Brower-Solidity 라고 하는 웹 전용 Solidity 개발 환경이다.
- 별도의 설치가 필요하지 않지만 오프라인으로 작업하고 싶으면 [여길](https://remix-ide.readthedocs.io/en/latest/remixd.html#remixd-installation) 참고

---

## 개념

### Smart Contract

- 스마트 컨트랙트는 특정한 주소를 가진다. ( 하나의 암호화폐 지갑이다.)
- 암호화폐 지갑의 형태로 존재하고 balance를 가지고 있으면서 트랜잭션을 보낼수 있다.
- 이더리움 네트워크에 배포되게 되면  프로그램의 형태로 실행되게 된다.
- 사람이 컨트롤 할 수 없고 모든사람이 이 스마트 컨트랙트를 사용할 수 있다.
- 사용하게 되면 스마트 컨트랙트 안에 있는 특정한 함수를 실행하게 된다.
- Solidity 언어로 개발되며 Java와 JavaScript 와 유사하며 객체지향이고 정적타입의 언어이다.


### Deploy

![스크린샷 2022-01-07 오후 1 21 11](https://user-images.githubusercontent.com/56465854/148492075-a0f15cd3-110d-4428-80d3-5048e390bde6.png)

- 솔리디티 스마트 컨트랙트를 작성한 후 [Solidity Compiler](https://github.com/ethereum/solidity) 를 통한 .bin 과 .abi 파일 생성
- 바이너리 파일 (.bin) 은 어더리움 네트워크(EVM) 에 배포
- abi 는 (API 와 비슷한 개념) abigen 을 통해 고 파일로 변경이 가능함 

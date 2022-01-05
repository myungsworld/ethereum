#이더리움 네트워크 튜토리얼

## 1. [마지막 블록 조회](https://github.com/myungsworld/ethereum/blob/main/tutorial/1.lastest_block_number.go)

## 2. [공개주소로 지갑 잔액 확인](https://github.com/myungsworld/ethereum/blob/main/tutorial/2.get_ether_from_address.go)

## 3. [지갑 생성](https://github.com/myungsworld/ethereum/blob/main/tutorial/3.generate_wallet.go)

- Private Key 생성 -> Public Key 생성 -> 이더 지갑 생성 

## 4. [암호 지갑 생성]()

![스크린샷 2022-01-05 오후 1 00 05](https://user-images.githubusercontent.com/56465854/148158568-8d554bfc-54ba-4766-a513-1414ee32d8e8.png)

- 암호화된 지갑의 비밀번호를 잃어버리면 대칭키 알고리즘에 따라서
- PrivateKey 와 PublicKey를 가져올수 없으므로 그 지갑은 사용하지 못하는 지갑이 된다.

## 5. [테스트넷을 통한 이더 가져오기]()

- 이더리움 네트워크의 이해와 연습을 위한 테스트 네트워크가 있다.
- ex) Gorli, Kovan , Rinkeby, Ropsten 
- 각각의 네트워크에 만들어진 지갑으로는 다른 지갑과는 별개이므로 이더를 전송하고 사용할수 없다.
- Infura 에서 각각의 테스트 네트워크 엔드포인트를 지원한다.
- [Kovan](https://kovan.etherscan.io/address/0xb347b9f5b56b431b2cf4e1d90a5995f7519ca792#writeContract) 여기서 테스트를 진행했다. 가입후 address 를 채팅에 넣으면 일정량의 이더를 지급받는다.
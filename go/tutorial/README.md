#이더리움 네트워크 튜토리얼

## 1. [마지막 블록 조회](https://github.com/myungsworld/ethereum/blob/main/go/tutorial/1.lastest_block_number.go)



## 2. [공개주소로 지갑 잔액 확인](https://github.com/myungsworld/ethereum/blob/main/go/tutorial/2.get_ether_from_address.go)



## 3. [지갑 생성](https://github.com/myungsworld/ethereum/blob/main/go/tutorial/3.generate_wallet.go)

- Private Key 생성 -> Public Key 생성 -> 이더 지갑 생성 



## 4. [암호 지갑 생성](https://github.com/myungsworld/ethereum/blob/main/go/tutorial/4.generate_ehtereum_keystore_wallet.go)

![스크린샷 2022-01-05 오후 1 00 05](https://user-images.githubusercontent.com/56465854/148158568-8d554bfc-54ba-4766-a513-1414ee32d8e8.png)

- 암호화된 지갑의 비밀번호를 잃어버리면 대칭키 알고리즘에 따라서
- PrivateKey 와 PublicKey를 가져올수 없으므로 그 지갑은 사용하지 못하는 지갑이 된다.



## 5. [테스트넷을 통한 이더 가져오기](https://github.com/myungsworld/ethereum/blob/main/go/tutorial/5.get_ether_from_test_network.go)

- 이더리움 네트워크의 이해와 연습을 위한 테스트 네트워크가 있다.
- ex) Gorli, Kovan , Rinkeby, Ropsten 
- 각각의 네트워크에 만들어진 지갑으로는 다른 지갑과는 별개이므로 이더를 전송하고 사용할수 없다.
- Infura 에서 각각의 테스트 네트워크 엔드포인트를 지원한다.
- [Kovan](https://gitter.im/kovan-testnet/faucet) 여기서 테스트를 진행했다. 가입후 address 를 채팅에 넣으면 일정량의 이더를 지급받는다.



## 6. [다른계좌로 이더 전송(트랜잭션)](https://github.com/myungsworld/ethereum/blob/main/go/tutorial/6.make_transaction.go)

![스크린샷 2022-01-05 오후 3 08 53](https://user-images.githubusercontent.com/56465854/148168594-afaaae30-dca1-47cc-ab7e-ebcb68e747d2.png)



- from : 보내는 주소 , to : 받는 주소
- value : 보내는 이더량 -> 웨이(Wei) 단위 (이더의 10의 -18승)
- gasLimit : 이더리움 네트워크에서 트랜잭션을 만들기위한 고정값
- gasPrice : 가스 프라이스를 너무 낮게 잡으면 트랜잭션 실패 확률이 올라감 , 너무 높게 잡으면 비용이 많이 듬
- nonce : 트랜잭션 발생 횟수 , nonce를 제대로 기입하지 않으면 에러를 반환한다.

5개를 설정한뒤 프라이빗키로 트랜잭션을 암호화 시킨후 요청을 보낸다.   
해커가 이걸 암호화 시키려고 시도하면 이더리움 네트워크에서 자동으로 이 트랜잭션을 거절시킨다.   
gasLimit 과 gasPrice 를 곱한 이더가 수수료가 된다.

위의 예제로는 21000 * 200 = 4,200,000 Gwei 즉, 0.0042 ETH 가 수수료로 지불한다.   
이 수수료는 마이닝을 하는 사람에게 간다.

## 7. [내 스마트컨트랙트 배포하기](https://github.com/myungsworld/ethereum/blob/main/go/tutorial/7.deploy_smart_contract.go)

- abigen 으로 생성한 go 파일안에 있는 Deploy 함수로 이더리움 네트워크에 쉽게 스마트 컨트랙트를 배포할수 있다.
- 프라이빗키 , chainID , gasPrice 등 기본적인 배포에 대한 정보들이 필요하다.

## 8. [배포된 컨트랙트 핸들링하기](https://github.com/myungsworld/ethereum/blob/main/go/tutorial/8.interact_my_deployed_contract.go)

- 고파일에 있는 함수들로 컨트랙트를 고언어로 핸들링이 가능하다.
- solidity로 작성된 함수들이 고 함수들로 변경이 되며 권한에 따라 읽기,쓰기,수정,삭제가 가능하다.



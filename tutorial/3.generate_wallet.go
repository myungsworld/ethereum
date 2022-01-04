package tutorial

import (
	"fmt"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
)

func GenerateWallet() {

	// 프라이빗키 생성
	privateKey , err := crypto.GenerateKey()
	if err != nil {
		fmt.Println("err key generate :", err)
	}

	// 퍼블릭키 생성
	publicKey := &privateKey.PublicKey

	// 이더리움 지갑 생성
	wallet := crypto.PubkeyToAddress(privateKey.PublicKey)


	privateData := crypto.FromECDSA(privateKey)
	fmt.Println("프라이빗 키 :", hexutil.Encode(privateData))

	publicData := crypto.FromECDSAPub(publicKey)
	fmt.Println("퍼블릭 키 : ", hexutil.Encode(publicData))

	fmt.Println("이더리움 지갑주소 :", wallet.Hex())
}

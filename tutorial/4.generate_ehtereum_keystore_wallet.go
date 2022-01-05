package tutorial

import (
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"io/ioutil"
)

func GenerateKeystore() {

	password := "password"

	////wallet 에 keystore 생성하는 부분
	//key := keystore.NewKeyStore("./wallet", keystore.StandardScryptN, keystore.StandardScryptP)

	//a, err := key.NewAccount(password)
	//if err != nil {
	//	fmt.Println(err)
	//	panic(err)
	//}
	//fmt.Println(a)

	// 파일을 읽고 password로 프라이빗키, 퍼블릭키, 주소 가져오기
	b , err := ioutil.ReadFile("./wallet/UTC--2022-01-05T04-43-29.938823000Z--369c0ab74788d901a563192e956674aa173d6d0d")
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	key , err := keystore.DecryptKey(b,password)
	if err != nil {
		panic(err)
	}

	privateData := crypto.FromECDSA(key.PrivateKey)
	fmt.Println("파일에서 비밀번호로 가져온 프라이빗키 :", hexutil.Encode(privateData))
	publicData := crypto.FromECDSAPub(&key.PrivateKey.PublicKey)
	fmt.Println("프라이빗키로 가져온 퍼블릭키:", hexutil.Encode(publicData))
	addr := crypto.PubkeyToAddress(key.PrivateKey.PublicKey).Hex()
	fmt.Println("퍼블릭키로 가져온 지갑주소:", addr)
}
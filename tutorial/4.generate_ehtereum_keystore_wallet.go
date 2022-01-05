package tutorial

import "github.com/ethereum/go-ethereum/accounts/keystore"

func GenerateKeystore() {
	key := keystore.NewKeyStore("./wallet", keystore.StandardScryptN, keystore.StandardScryptP)

}
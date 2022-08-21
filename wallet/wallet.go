package wallet

import (
	"crypto/ecdsa"
	"os"
)

type wallet struct {
	privateKey *ecdsa.PrivateKey
}

var w *wallet

func hasWalletFile() bool {
	_, err := os.Stat("baami.wallet")
	return !os.IsNotExist(err)
}

// signature, priavte key, public key 에 대서만
func Wallet() *wallet {
	if w == nil {
		// has a wallet alread?
		if hasWalletFile() {
			// yes -> restore from file

		} else {
			// no -> create private key, save to file

		}
	}
	return w
}

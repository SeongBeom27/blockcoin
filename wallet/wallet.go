package wallet

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/x509"
	"os"

	"github.com/baaami/blockcoin/utils"
)

type wallet struct {
	privateKey *ecdsa.PrivateKey
	address    string // hexa public key
}

const (
	fileName string = "baami.wallet"
)

var w *wallet

func hasWalletFile() bool {
	_, err := os.Stat(fileName)
	return !os.IsNotExist(err)
}

func createPrivKey() *ecdsa.PrivateKey {
	privKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	utils.HandleErr(err)
	return privKey
}

func persistKey(key *ecdsa.PrivateKey) {
	bytes, err := x509.MarshalECPrivateKey(key)
	utils.HandleErr(err)
	err = os.WriteFile(fileName, bytes, 0644)
	utils.HandleErr(err)
}

func restoreKey() *ecdsa.PrivateKey {
	keyAsBytes, err := os.ReadFile(fileName)
	utils.HandleErr(err)
	key, err := x509.ParseECPrivateKey(keyAsBytes)
	utils.HandleErr(err)
	return key
}

// signature, priavte key, public key 에 대서만
func Wallet() *wallet {
	if w == nil {
		// has a wallet alread?
		w = &wallet{}

		if hasWalletFile() {
			// yes -> restore from file
			w.privateKey = restoreKey()
		} else {
			// no -> create private key, save to file

			// 1. create key
			key := createPrivKey()
			// 2. save key to file syste
			persistKey(key)
			w.privateKey = key
		}
	}
	return w
}

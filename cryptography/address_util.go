package cryptography

import (
	"crypto/sha256"
	"ericshu.io/aelf-chain-test/aelf"
	"github.com/btcsuite/btcutil/base58"
)

func GenerateAddressFromPublicKey(publicKey []byte) (address aelf.Address) {
	hash1 := CalculateHash(publicKey)
	hash2 :=  CalculateHash(hash1)
	address = aelf.Address{
		Value: hash2,
	}

	return
}

func GetAddressFormattedFromPublicKey(publicKey []byte) string {
	return base58.Encode(publicKey)
}

func GetAddressFormatted(address aelf.Address) string {
	return base58.Encode(address.Value)
}

func CalculateHash(bytes []byte) []byte {
	sha  := sha256.New()
	sha.Write(bytes)
	return sha.Sum(nil)
}

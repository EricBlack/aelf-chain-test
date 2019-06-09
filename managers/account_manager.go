package managers

import (
	"ericshu.io/aelf-chain-test/aelf"
	"ericshu.io/aelf-chain-test/cryptography"
	"github.com/haltingstate/secp256k1-go"
)

type AddressInfo struct {
	PublicKey []byte
	PrivateKey  []byte
	Address aelf.Address
	Account string
}

var AddressCollection map[string]AddressInfo

func init(){
	AddressCollection = make(map[string]AddressInfo)
}

func GenerateNewAddress() (addressInfo AddressInfo) {
	pubKey, priKey := secp256k1.GenerateKeyPair()
	address := cryptography.GenerateAddressFromPublicKey(pubKey)
	addressString := cryptography.GetAddressFormatted(address)
	addressInfo = AddressInfo{
		PublicKey: pubKey,
		PrivateKey: priKey,
		Address: address,
		Account: addressString,
	}

	AddressCollection[addressString] = addressInfo
	return
}

func GetAddressByString(account string) (address AddressInfo) {
	if val, ok := AddressCollection[account]; ok {
		address = val
	} else {
		address = AddressInfo{}
	}

	return
}

func (address AddressInfo) SignTransaction(transaction aelf.Transaction) (aelf.Transaction) {
	data := GetSignatureData(transaction)
	from := cryptography.GetAddressFormattedFromPublicKey(transaction.From.Value)
	addressInfo := GetAddressByString(from)
	transaction.Signature = secp256k1.Sign(data, addressInfo.PrivateKey)

	return transaction
}
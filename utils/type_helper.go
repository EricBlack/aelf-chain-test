package utils

import "encoding/hex"

func FromHexString(hexString string) (bytes []byte, err error) {
	bytes, err = hex.DecodeString(hexString)

	return
}

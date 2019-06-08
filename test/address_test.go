package test

import (
	"ericshu.io/aelf-chain-test/managers"
	"log"
	"testing"
)

func Test_GenerateAddress(t *testing.T){
	address := managers.GenerateNewAddress()
	log.Println("Address:", address.Account)
}

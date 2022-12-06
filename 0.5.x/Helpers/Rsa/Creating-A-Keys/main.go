package main

import (
	"fmt"
	helper_rsa "github.com/langwan/langgo/helpers/crypto/rsa"
)

func main() {
	publicKeyString, privateKeyString := helper_rsa.CreateKeyX509PKCS1(2048)
	fmt.Printf("publicKeyString = %v\n", publicKeyString)
	fmt.Printf("privateKeyString = %v\n", privateKeyString)
	publicKey, err := helper_rsa.PublicKeyFromX509PKCS1(publicKeyString)
	if err != nil {
		panic(err)
	}
	privateKey, err := helper_rsa.PrivateKeyFromX509PKCS1(privateKeyString)
	if err != nil {
		panic(err)
	}
	fmt.Printf("publicKey = %v\n", publicKey)
	fmt.Printf("privateKey = %v\n", privateKey)
}

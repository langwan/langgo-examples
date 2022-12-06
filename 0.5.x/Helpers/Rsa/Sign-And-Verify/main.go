package main

import (
	"fmt"
	helper_rsa "github.com/langwan/langgo/helpers/crypto/rsa"
)

func main() {
	bits := 2048
	publicKeyString, privateKeyString := helper_rsa.CreateKeyX509PKCS1(bits)
	data := "0123456789"
	publicKey, err := helper_rsa.PublicKeyFromX509PKCS1(publicKeyString)
	if err != nil {
		panic(err)
	}
	privateKey, err := helper_rsa.PrivateKeyFromX509PKCS1(privateKeyString)
	if err != nil {
		panic(err)
	}

	sign, err := helper_rsa.Sign(privateKey, []byte(data))
	if err != nil {
		panic(err)
	}
	err = helper_rsa.Verify(publicKey, sign, []byte(data))
	if err != nil {
		panic(err)
	}
	fmt.Println("verify ok")
}

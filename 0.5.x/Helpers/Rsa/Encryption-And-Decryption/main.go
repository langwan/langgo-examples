package main

import (
	"fmt"
	helper_rsa "github.com/langwan/langgo/helpers/crypto/rsa"
)

func main() {
	bits := 2048
	data := "0123456789"

	publicKeyString, privateKeyString := helper_rsa.CreateKeyX509PKCS1(bits)
	publicKey, err := helper_rsa.PublicKeyFromX509PKCS1(publicKeyString)
	if err != nil {
		panic(err)
	}
	privateKey, err := helper_rsa.PrivateKeyFromX509PKCS1(privateKeyString)
	if err != nil {
		panic(err)
	}
	encrypt, err := helper_rsa.Encrypt(publicKey, []byte(data))
	if err != nil {
		panic(err)
	}
	fmt.Printf("encrypt=%v\n", encrypt)
	decrypt, err := helper_rsa.Decrypt(privateKey, encrypt)
	if err != nil {
		panic(err)
	}
	fmt.Printf("decrypt=%s\n", string(decrypt))
}

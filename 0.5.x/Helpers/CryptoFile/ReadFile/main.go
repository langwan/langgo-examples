package main

import (
	helper_crypto_file "github.com/langwan/langgo/helpers/crypto/file"
	"os"
)

func main() {
	key := []byte("01234567890123456789012345678901")
	data, err := os.ReadFile("./icon.jpg")
	if err != nil {
		return
	}
	helper_crypto_file.WriteFile(key, "./dst.jpg.en", data, 0666)

	de, err := helper_crypto_file.ReadFile(key, "./dst.jpg.en")
	if err != nil {
		return
	}
	os.WriteFile("./dst.jpg", de, 0666)
}

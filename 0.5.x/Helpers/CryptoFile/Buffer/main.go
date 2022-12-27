package main

import (
	"fmt"
	helper_crypto_file "github.com/langwan/langgo/helpers/crypto/file"
	"io"
	"os"
	"time"
)

func main() {
	encryptedFile("01234567890123456789012345678901", "./icon.jpg", "./dst.jpg", "./dst.jpg.en")
	//encryptedFile("01234567890123456789012345678901", "./sample.mp4", "./dst.mp4", "./dst.mp4.en")
}

func encryptedFile(key string, src string, dst string, en string) {

	st := time.Now()

	defer func() {
		fmt.Println("runtime", time.Since(st))
	}()

	os.Remove(en)
	os.Remove(dst)

	cf := helper_crypto_file.File{
		Secret: []byte(key),
	}
	fd, err := os.Open(src)
	defer fd.Close()
	if err != nil {
		panic(err)
	}
	err = cf.Create(en)
	if err != nil {
		panic(err)
	}
	buf := make([]byte, helper_crypto_file.DefaultBufferSize)
	rn := 0
	wn := 0
	i := 0
	for {
		n, err := fd.ReadAt(buf, int64(rn))

		i++
		rn += n
		if n == 0 {
			break
		} else if err != nil && err != io.EOF {

			panic(err)
		}
		n, err = cf.WriteAt(buf[:n], int64(wn))

		if err != nil {
			panic(err)
		}
		wn += n
	}

	cf.Close()

	cf = helper_crypto_file.File{
		Secret: []byte("01234567890123456789012345678901"),
	}
	cf.Open(en)
	defer cf.Close()

	fd, err = os.Create(dst)
	if err != nil {
		panic(err)
	}
	defer fd.Close()
	buf = make([]byte, helper_crypto_file.DefaultBlockSize)
	rn = 0
	wn = 0

	i = 0
	for {
		de, n, err := cf.ReadAt(buf, int64(rn))

		i++
		rn += n
		if n == 0 {
			break
		} else if err != nil && err != io.EOF {

			panic(err)
		}
		n, err = fd.WriteAt(de, int64(wn))

		if err != nil {
			panic(err)
		}
		wn += n
	}
}

package main

import (
	"fmt"
	helper_size "github.com/langwan/langgo/helpers/size"
)

func main() {
	size := helper_size.BytesSize(1024 * 1024)
	fmt.Printf("1024 BytesSize is %s\n", size)
	size = helper_size.HumanSize(1000 * 1000)
	fmt.Printf("1024 HumanSize is %s\n", size)
	if val, err := helper_size.RAMInBytes("1Kib"); err == nil {
		fmt.Printf("1Mib RAMInBytes is %d\n", val)
	}
	if val, err := helper_size.FromHumanSize("1K"); err == nil {
		fmt.Printf("1Mib FromHumanSize is %d\n", val)
	}
}

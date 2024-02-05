package main

import (
	"crypto/md5"
	"fmt"
	"io"
	"strconv"
)

func main() {

	input := "bgvyzdsv"

	for i := 0; ; i++ {
		hash := md5.New()
		io.WriteString(hash, input)
		io.WriteString(hash, strconv.Itoa(i))
		x := hash.Sum(nil)

		// Check if the first 2.5 bytes are 0
		if x[0] == 0 && x[1] == 0 && x[2] < 16 {
			fmt.Printf("Total: %d %x\n", i, x)
			return
		}

		if i > 1000000 {
			fmt.Println("No result found")
			return
		}
	}

}

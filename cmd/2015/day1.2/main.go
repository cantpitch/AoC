package main

import (
	"fmt"
	"log"
	"os"
	"pkg/file"
)

func main() {
	args := os.Args[1:] // Skip program name

	f := file.NewFile(args[0])
	parens, _ := f.Next()
	floor := 0
	idx := 0
	for i := 0; i < len(parens); i++ {
		idx++
		r := parens[i]
		if r == '(' {
			floor++
		} else if r == ')' {
			floor--
		} else {
			log.Fatalf("Unexpected character: %c", r)
		}

		if floor < 0 {
			break
		}
	}

	fmt.Printf("Basement Index: %d\n", idx)
}

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
	for i := 0; i < len(parens); i++ {
		r := parens[i]
		if r == '(' {
			floor++
		} else if r == ')' {
			floor--
		} else {
			log.Fatalf("Unexpected character: %c", r)
		}
	}

	fmt.Printf("Final Floor: %d\n", floor)
}

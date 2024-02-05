package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	args := os.Args[1:] // Skip program name

	file, err := os.Open(args[0])
	if err != nil {
		panic(err)
	}

	defer func() {
		if err := file.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	scanner := bufio.NewScanner(bufio.NewReader(file))

	total := 0
	world := make(map[string]int)
	x, y := 0, 0

	for scanner.Scan() {
		dirs := scanner.Text()
		for _, r := range dirs {
			key := fmt.Sprintf("%d,%d", x, y)
			world[key]++
			if world[key] == 1 {
				total++
			}

			switch r {
			case '^':
				y++
			case 'v':
				y--
			case '>':
				x++
			case '<':
				x--
			}

		}
	}

	fmt.Printf("Total: %d\n", total)
}

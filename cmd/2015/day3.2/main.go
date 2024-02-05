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
	x := make([]int, 2)
	y := make([]int, 2)

	i := 0
	for scanner.Scan() {
		dirs := scanner.Text()
		for _, r := range dirs {
			key := fmt.Sprintf("%d,%d", x[i], y[i])
			world[key]++
			if world[key] == 1 {
				total++
			}

			switch r {
			case '^':
				y[i]++
			case 'v':
				y[i]--
			case '>':
				x[i]++
			case '<':
				x[i]--
			}

			i = (i + 1) % 2
		}
	}

	fmt.Printf("Total: %d\n", total)
}

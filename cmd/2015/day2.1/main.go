package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
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
	for scanner.Scan() {
		p := scanner.Text()
		dims := strings.Split(p, "x")
		l, _ := strconv.Atoi(dims[0])
		w, _ := strconv.Atoi(dims[1])
		h, _ := strconv.Atoi(dims[2])
		ft := 2*(l*w+w*h+h*l) + min(l*w, w*h, h*l)
		total += ft
	}

	fmt.Printf("Total: %d\n", total)
}

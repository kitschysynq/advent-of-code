package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	var (
		acc uint64
		max uint64
	)

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		val := scanner.Text()
		switch val {
		case "":
			if acc > max {
				max = acc
			}
			acc = 0
		default:
			n, err := strconv.ParseUint(val, 10, 64)
			if err != nil {
				log.Fatal(err)
			}
			acc += n
		}
	}
	if acc > max {
		max = acc
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Max: %d\n", max)
}

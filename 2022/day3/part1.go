package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var itemTypes = " abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func priority(b byte) int { return strings.IndexByte(itemTypes, b) }

func main() {
	var total int
	seen := make([]bool, 53)
	dups := make([]bool, 53)

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		val := scanner.Bytes()
		for _, c := range val[:len(val)/2] {
			seen[priority(c)] = true
		}
		for _, c := range val[len(val)/2:] {
			if p := priority(c); seen[p] {
				dups[p] = true
			}
		}
		for i, c := range dups {
			if c {
				total += i
			}
		}
		seen = make([]bool, 53)
		dups = make([]bool, 53)
	}

	fmt.Println(total)
}

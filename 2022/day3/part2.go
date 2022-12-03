package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const groupSize = 3

var itemTypes = " abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func priority(b byte) int { return strings.IndexByte(itemTypes, b) }

func main() {
	var total int
	seen := make([]bool, 53)
	has := make([]int, 53)

	var idx int
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		idx++
		val := scanner.Bytes()
		for _, c := range val {
			if p := priority(c); !seen[p] {
				has[p]++
				seen[p] = true
			}
		}
		seen = make([]bool, 53)
		if idx%groupSize != 0 {
			continue
		}
		for i, c := range has {
			if c == groupSize {
				total += i
			}
		}
		has = make([]int, 53)
	}

	fmt.Println(total)
}

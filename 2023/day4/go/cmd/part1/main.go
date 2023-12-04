package main

import (
	"bufio"
	"day4"
	"fmt"
	"os"
)

func main() {
	var acc int
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		m := day4.ParseCard(scanner.Text()).Matches()
		if len(m) > 0 {
			acc += 1 << (len(m) - 1)
		}
	}
	fmt.Println(acc)
}

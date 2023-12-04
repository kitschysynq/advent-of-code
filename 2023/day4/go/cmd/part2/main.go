package main

import (
	"bufio"
	"day4"
	"fmt"
	"os"
)

func main() {
	var acc int
	line := 1
	var wins []int
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		acc += len(wins)
		m := day4.ParseCard(scanner.Text()).Matches()
		acc++
		line++
		if len(m) > 0 {
			cnt := len(wins) + 1 // Get additional cards for previously won copies + the original card
			for i := 0; i < cnt; i++ {
				wins = append(wins, len(m))
			}
		}
		wins = prune(wins)
	}
	fmt.Println(acc)
}

func prune(w []int) []int {
	var x []int
	for _, v := range w {
		v--
		if v >= 0 {
			x = append(x, v)
		}
	}
	return x
}

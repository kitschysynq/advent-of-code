package main

import (
	"day6"
	"fmt"
	"os"
)

func main() {
	r := day6.ParseRaces(os.Stdin)
	acc := 1
	for i := range r.Time {
		min, max := day6.Roots(r.Time[i], r.Dist[i])
		ways := (max - min) + 1
		acc *= ways
		//fmt.Printf("%d-%d=%d\t| %d\n", max, min, ways, acc)
	}
	fmt.Println(acc)
}

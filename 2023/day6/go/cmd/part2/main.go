package main

import (
	"day6"
	"fmt"
	"os"
)

func main() {
	r := day6.ParseRaces2(os.Stdin)
	min, max := day6.Roots(r.Time[0], r.Dist[0])
	ways := (max - min) + 1
	fmt.Println(ways)
}

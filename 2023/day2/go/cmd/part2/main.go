package main

import (
	"bufio"
	"day2"
	"encoding/json"
	"fmt"
	"log"
	"os"
)

func main() {
	var acc int

	scanner := bufio.NewScanner(os.Stdin)
	enc := json.NewEncoder(os.Stdout)
	enc.SetIndent("", "  ")
	for scanner.Scan() {
		g, err := day2.ParseGame(scanner.Text())
		if err != nil {
			log.Println("error parsing game: %s", err.Error())
			continue
		}
		max := make(map[string]int)
		for _, r := range g.Rounds {
			for k, v := range r {
				if v > max[k] {
					max[k] = v
				}
			}
		}
		mul := 1
		for _, v := range max {
			mul *= v
		}
		acc += mul
	}
	fmt.Println(acc)
}

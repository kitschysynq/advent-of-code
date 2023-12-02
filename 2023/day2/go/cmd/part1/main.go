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
	bag := day2.Bag{
		Cubes: map[string]int{
			"red":   12,
			"green": 13,
			"blue":  14,
		},
	}

	var acc int

	scanner := bufio.NewScanner(os.Stdin)
	enc := json.NewEncoder(os.Stdout)
	enc.SetIndent("", "  ")
	for scanner.Scan() {
		possible := true
		g, err := day2.ParseGame(scanner.Text())
		if err != nil {
			log.Println("error parsing game: %s", err.Error())
			continue
		}
		for _, r := range g.Rounds {
			if !bag.Possible(r) {
				//enc.Encode(r)
			}
			possible = possible && bag.Possible(r)
		}
		//fmt.Printf("%s\n\t%q\n", scanner.Text(), gRE.FindAllStringSubmatch(scanner.Text(), -1))
		//_ = enc.Encode(g)
		if possible {
			acc += g.ID
		}
	}
	fmt.Println(acc)
}

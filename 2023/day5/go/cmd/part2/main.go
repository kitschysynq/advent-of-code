package main

import (
	"bufio"
	"day5"
	"errors"
	"fmt"
	"io"
	"log"
	"math"
	"os"
	"strings"
)

func main() {
	b := bufio.NewReader(os.Stdin)
	seedStr, err := b.ReadString('\n')
	if err != nil {
		log.Fatalf("error reading seeds: %s", err.Error())
	}
	seeds, err := parseSeeds(seedStr[:len(seedStr)-1])
	if err != nil {
		log.Fatalf("invalid seed spec: %s", err.Error())
	}

	p := day5.NewParser(b)
	m, err := p.Parse()
	if err != nil {
		log.Fatalf("error parsing input: %s", err.Error())
	}

	acc := math.MaxInt
	fmt.Printf("processing %d seed ranges", len(seeds))
	for _, i := range seeds {
		fmt.Print(".")
		for j := i[0]; j < i[0]+i[1]; j++ {
			l := mapSeedAll(m, j)
			//fmt.Printf("Seed %d maps to location %d\n", i, l)
			if l < acc {
				acc = l
			}
		}
	}
	fmt.Println()
	fmt.Println(acc)
}

func parseSeeds(seedStr string) ([][2]int, error) {
	if !strings.HasPrefix(seedStr, "seeds: ") {
		return nil, errors.New("incorrect prefix")
	}

	var seeds [][2]int
	r := strings.NewReader(seedStr[7:])
	for {
		var seed [2]int
		_, err := fmt.Fscanf(r, "%d %d", &seed[0], &seed[1])
		if err != nil {
			if err == io.EOF {
				return seeds, nil
			}
			return nil, fmt.Errorf("invalid seed num at pos %d: %s", len(seeds)+1, err.Error())
		}
		seeds = append(seeds, seed)
	}
}

func mapSeedAll(m *day5.Map, i int) int {
	k, n, ok := "seed", i, true
	for ok {
		var k1 string
		var n1 int
		k1, n1, ok = m.To(k, n)
		if ok {
			k, n = k1, n1
		}
	}
	return n
}

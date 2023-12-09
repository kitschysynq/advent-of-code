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
	for _, i := range seeds {
		l := mapSeedAll(m, i)
		fmt.Printf("Seed %d maps to location %d\n", i, l)
		if l < acc {
			acc = l
		}
	}
	fmt.Println(acc)
}

func parseSeeds(seedStr string) ([]int, error) {
	if !strings.HasPrefix(seedStr, "seeds: ") {
		return nil, errors.New("incorrect prefix")
	}

	var seeds []int
	r := strings.NewReader(seedStr[7:])
	for {
		var seed int
		_, err := fmt.Fscanf(r, "%d", &seed)
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

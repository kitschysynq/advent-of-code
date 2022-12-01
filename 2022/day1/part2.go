package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
)

type Elf struct {
	idx int
	cal uint64
}

func (e Elf) String() string { return fmt.Sprintf("(idx: %d, cal %d)", e.idx, e.cal) }

type Elves []Elf

func (e Elves) Len() int           { return len(e) }
func (e Elves) Less(i, j int) bool { return e[i].cal < e[j].cal }
func (e Elves) Swap(i, j int)      { e[i], e[j] = e[j], e[i] }

func main() {
	var (
		idx   int
		acc   uint64
		elves Elves
	)

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		val := scanner.Text()
		switch val {
		case "":
			elves = append(elves, Elf{idx, acc})
			idx++
			acc = 0
		default:
			n, err := strconv.ParseUint(val, 10, 64)
			if err != nil {
				log.Fatal(err)
			}
			acc += n
		}
	}
	elves = append(elves, Elf{idx, acc})
	acc = 0

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	sort.Sort(sort.Reverse(elves))

	for _, e := range elves[0:3] {
		acc += e.cal
	}

	fmt.Printf("Top 3 elves are carrying %d cal in total\n", acc)
}

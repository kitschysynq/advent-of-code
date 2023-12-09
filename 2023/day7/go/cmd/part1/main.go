package main

import (
	"bufio"
	"day7"
	"fmt"
	"os"
	"slices"
)

type entry struct {
	Hand string
	Bid  int
}

func Compare(a, b entry) int {
	return day7.Compare(a.Hand, b.Hand)
}

func main() {
	//hands := []entry{
	//	{"32T3K", 765},
	//	{"T55J5", 684},
	//	{"KK677", 28},
	//	{"KTJJT", 220},
	//	{"QQQJA", 483},
	//}
	var hands []entry
	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		var e entry
		_, err := fmt.Sscanf(s.Text(), "%s %d", &e.Hand, &e.Bid)
		if err != nil {
			panic(err)
		}
		hands = append(hands, e)
	}
	slices.SortFunc(hands, Compare)
	var acc int
	for i, hand := range hands {
		acc += hand.Bid * (i + 1)
	}
	fmt.Println(acc)
}

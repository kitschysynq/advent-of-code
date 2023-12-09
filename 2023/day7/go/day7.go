package day7

import (
	"slices"
	"strings"
)

var Cards = "23456789TJQKA"
var CardsWild = "J23456789TQKA"

type HandType int

const (
	Invalid HandType = iota
	HighCard
	OnePair
	TwoPair
	ThreeOfAKind
	FullHouse
	FourOfAKind
	FiveOfAKind
)

var htString = map[HandType]string{
	HighCard:     "High Card",
	OnePair:      "One Pair",
	TwoPair:      "Two Pair",
	ThreeOfAKind: "Three of a Kind",
	FullHouse:    "Full House",
	FourOfAKind:  "Four of a Kind",
	FiveOfAKind:  "Five of a Kind",
}

func (h HandType) String() string {
	return htString[h]
}

func TypeOf(hand string) HandType {
	chars := make(map[rune]int)
	for _, c := range hand {
		chars[c]++
	}
	var max int
	for _, v := range chars {
		if v > max {
			max = v
		}
	}
	switch l := len(chars); l {
	case 5:
		return HighCard
	case 4:
		return OnePair
	case 3:
		if max == 3 {
			return ThreeOfAKind
		}
		return TwoPair
	case 2:
		if max == 4 {
			return FourOfAKind
		}
		return FullHouse
	case 1:
		return FiveOfAKind
	}
	return Invalid
}

func TypeOfWild(hand string) HandType {
	h := slices.MaxFunc(Options(hand), Compare)
	return TypeOf(h)
}

func Compare(a, b string) int {
	return compare(a, b, TypeOf, Strength)
}

func Strength(a byte) int {
	return strings.Index(Cards, string(a))
}

func StrengthWild(a byte) int {
	return strings.Index(CardsWild, string(a))
}

func CompareWild(a, b string) int {
	return compare(a, b, TypeOfWild, StrengthWild)
}

func compare(a, b string, typeOf func(string) HandType, strength func(byte) int) int {
	tA, tB := typeOf(a), typeOf(b)
	if tA > tB {
		return 1
	}
	if tA < tB {
		return -1
	}
	for i := range a {
		idxA, idxB := strength(a[i]), strength(b[i])
		if idxA > idxB {
			return 1
		}
		if idxA < idxB {
			return -1
		}
	}
	return 0
}

func Options(hand string) []string {
	if !strings.Contains(hand, "J") {
		return []string{hand}
	}
	chars := make(map[rune]int)
	for _, c := range hand {
		chars[c]++
	}
	var max int
	for _, v := range chars {
		if v > max {
			max = v
		}
	}
	delete(chars, 'J')
	if len(chars) == 0 {
		// If there are only J in the hand,
		// then just return 5 aces, since we
		// know we're only interested in the
		// max anyway
		return []string{"AAAAA"}
	}
	opts := make(chan string, 3125)
	opts <- hand
	var ret []string
	for len(opts) > 0 {
		h := <-opts
		if !strings.Contains(h, "J") {
			ret = append(ret, h)
			continue
		}
		for c := range chars {
			opts <- strings.Replace(h, "J", string(c), 1)
		}
	}
	return ret
}

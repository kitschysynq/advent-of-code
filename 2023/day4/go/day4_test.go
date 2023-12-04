package day4

import (
	"fmt"
	"reflect"
	"testing"
)

func TestParseCard(t *testing.T) {
	tests := []struct {
		in     string
		parsed *Card
	}{
		{
			in: "Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53",
			parsed: &Card{
				ID:   1,
				Win:  []int{17, 41, 48, 83, 86},
				Have: []int{6, 9, 17, 31, 48, 53, 83, 86},
			},
		},
		{
			in: "Card 2: 13 32 20 16 61 | 61 30 68 82 17 32 24 19",
			parsed: &Card{
				ID:   2,
				Win:  []int{13, 16, 20, 32, 61},
				Have: []int{17, 19, 24, 30, 32, 61, 68, 82},
			},
		},
		{
			in: "Card 3:  1 21 53 59 44 | 69 82 63 72 16 21 14  1",
			parsed: &Card{
				ID:   3,
				Win:  []int{1, 21, 44, 53, 59},
				Have: []int{1, 14, 16, 21, 63, 69, 72, 82},
			},
		},
		{
			in: "Card 4: 41 92 73 84 69 | 59 84 76 51 58  5 54 83",
			parsed: &Card{
				ID:   4,
				Win:  []int{41, 69, 73, 84, 92},
				Have: []int{5, 51, 54, 58, 59, 76, 83, 84},
			},
		},
		{
			in: "Card 5: 87 83 26 28 32 | 88 30 70 12 93 22 82 36",
			parsed: &Card{
				ID:   5,
				Win:  []int{26, 28, 32, 83, 87},
				Have: []int{12, 22, 30, 36, 70, 82, 88, 93},
			},
		},
		{
			in: "Card 6: 31 18 13 56 72 | 74 77 10 23 35 67 36 11",
			parsed: &Card{
				ID:   6,
				Win:  []int{13, 18, 31, 56, 72},
				Have: []int{10, 11, 23, 35, 36, 67, 74, 77},
			},
		},
	}

	for i, test := range tests {
		t.Run(fmt.Sprintf("Line %d", i+1), func(t *testing.T) {
			if got, want := ParseCard(test.in), test.parsed; !reflect.DeepEqual(got, want) {
				t.Errorf("parsed card incorrectly;\n\t got: %v\n\twant: %v", got, want)
			}
		})
	}
}

func TestMatches(t *testing.T) {
	tests := []struct {
		card  *Card
		match []int
	}{
		{
			card: &Card{
				Win:  []int{17, 41, 48, 83, 86},
				Have: []int{6, 9, 17, 31, 48, 53, 83, 86},
			},
			match: []int{17, 48, 83, 86},
		},
		{
			card: &Card{
				Win:  []int{13, 16, 20, 32, 61},
				Have: []int{17, 19, 24, 30, 32, 61, 68, 82},
			},
			match: []int{32, 61},
		},
		{
			card: &Card{
				Win:  []int{1, 21, 44, 53, 59},
				Have: []int{1, 14, 16, 21, 63, 69, 72, 82},
			},
			match: []int{1, 21},
		},
		{
			card: &Card{
				Win:  []int{41, 69, 73, 84, 92},
				Have: []int{5, 51, 54, 58, 59, 76, 83, 84},
			},
			match: []int{84},
		},
		{
			card: &Card{
				Win:  []int{26, 28, 32, 83, 87},
				Have: []int{12, 22, 30, 36, 70, 82, 88, 93},
			},
			match: nil,
		},
		{
			card: &Card{
				Win:  []int{13, 18, 31, 56, 72},
				Have: []int{10, 11, 23, 35, 36, 67, 74, 77},
			},
			match: nil,
		},
	}

	for i, test := range tests {
		t.Run(fmt.Sprintf("Card %d", i+1), func(t *testing.T) {
			if got, want := test.card.Matches(), test.match; !reflect.DeepEqual(got, want) {
				t.Errorf("found incorrect matches\n\t got: %v\n\twant: %v", got, want)
			}
		})
	}
}

package day4

import (
	"log"
	"slices"
	"strconv"
	"strings"
)

type Card struct {
	ID   int
	Win  []int
	Have []int
}

func ParseCard(str string) *Card {
	var c Card

	l := NewLexer(strings.NewReader(str))

	_, t, s := l.Lex()
	if t != CARD {
		log.Fatalf("expected CARD found %q", s)
	}
	v, err := strconv.Atoi(s)
	if err != nil {
		log.Fatalf("invalid card ID %q", s)
	}
	c.ID = v

	// Read winning numbers up to the pipe separator
	for {
		_, t, s = l.Lex()
		if t == PIPE {
			break
		}
		if t != NUM {
			log.Fatalf("expected NUM found %q", s)
		}
		v, err := strconv.Atoi(s)
		if err != nil {
			log.Fatalf("invalid numeric literal %q", s)
		}
		c.Win = append(c.Win, v)
	}

	// Read found numbers until the end
	for {
		_, t, s = l.Lex()
		if t == EOF {
			break
		}
		if t != NUM {
			log.Fatalf("expected NUM found %q", t.String())
		}
		v, err := strconv.Atoi(s)
		if err != nil {
			log.Fatalf("invalid numeric literal %q", s)
		}
		c.Have = append(c.Have, v)
	}

	slices.Sort(c.Win)
	slices.Sort(c.Have)

	return &c
}

func (c *Card) Matches() []int {
	var i, j int
	var m []int
	for i < len(c.Win) && j < len(c.Have) {
		if c.Win[i] < c.Have[j] {
			i++
		} else if c.Win[i] > c.Have[j] {
			j++
		} else {
			m = append(m, c.Win[i])
			i++
			j++
		}
	}
	return m
}

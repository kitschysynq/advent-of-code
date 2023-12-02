package day2

import (
	"fmt"
	"strings"
)

type Bag struct {
	Cubes map[string]int
}

func (b *Bag) Possible(cubes map[string]int) bool {
	possible := true
	for k, v := range cubes {
		possible = possible && v <= b.Cubes[k]
	}
	return possible
}

type Game struct {
	ID     int
	Rounds []map[string]int
	Raw    string
}

func ParseGame(str string) (*Game, error) {
	p := strings.Split(str, ":")
	if len(p) != 2 {
		return nil, fmt.Errorf("invalid game string %q", str)
	}

	id, err := ParseID(p[0])
	if err != nil {
		return nil, err
	}

	rounds, err := ParseRounds(p[1])
	if err != nil {
		return nil, err
	}

	return &Game{
		ID:     id,
		Rounds: rounds,
		Raw:    str,
	}, nil
}

func ParseID(str string) (int, error) {
	var id int
	n, err := fmt.Sscanf(str, "Game %d", &id)
	if err != nil {
		return 0, fmt.Errorf("error in game id: %s", err.Error())
	}
	if n != 1 {
		return 0, fmt.Errorf("count error in game id: got %d want 1", n)
	}
	return id, nil
}

func ParseRounds(str string) ([]map[string]int, error) {
	var r []map[string]int
	rounds := strings.Split(str, "; ")
	for _, round := range rounds {
		round, err := ParseRound(round)
		if err != nil {
			return nil, fmt.Errorf("error parsing rounds: %s", err.Error())
		}
		r = append(r, round)
	}
	return r, nil
}

func ParseRound(str string) (map[string]int, error) {
	c := make(map[string]int)
	colors := strings.Split(str, ", ")
	for _, col := range colors {
		cnt, colStr, err := ParseColor(col)
		if err != nil {
			return nil, fmt.Errorf("error parsing color: %s", err.Error())
		}
		c[colStr] = cnt
	}
	return c, nil
}

func ParseColor(str string) (int, string, error) {
	var cnt int
	var col string
	n, err := fmt.Sscanf(str, "%d %s", &cnt, &col)
	if err != nil {
		return 0, "", fmt.Errorf("error parsing color: %s", err.Error())
	}
	if n != 2 {
		return 0, "", fmt.Errorf("invalid color definition: %q", str)
	}
	return cnt, col, nil
}

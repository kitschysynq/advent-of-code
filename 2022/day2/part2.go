// You can edit this code!
// Click here and start typing.
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

const (
	rock = iota
	paper
	scissors
	end
)

var (
	plays   = "ABC"
	results = "XYZ"
)

func decodePlay(in byte) (int, error) {
	out := strings.IndexByte(plays, in)
	if out == -1 {
		return 0, fmt.Errorf("invalid byte %s", string(in))
	}
	return out, nil
}

func decodeResult(in byte) (int, error) {
	out := strings.IndexByte(results, in)
	if out == -1 {
		return 0, fmt.Errorf("invalid result %s", string(in))
	}
	return out - 1, nil
}

func main() {
	var total int
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		round := scanner.Text()
		plays := strings.Split(round, " ")
		if len(plays) != 2 {
			log.Fatalf("invalid round: %q", round)
		}

		left, err := decodePlay(plays[0][0])
		if err != nil {
			log.Fatalf("invalid opponent play: %q", string(plays[0][0]))
		}

		result, err := decodeResult(plays[1][0])
		if err != nil {
			log.Fatalf("invalid suggested result: %q", string(plays[1][0]))
		}

		right := play(left, result)
		total += scoreShape(right) + scoreRound(result)
	}
	fmt.Println(total)
}

func play(opp, res int) int    { return (opp + res + end) % end }
func scoreShape(shape int) int { return shape + 1 }
func scoreRound(res int) int   { return (res + 1) * 3 }

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

var chars = "ABCXYZ"

func decodePlay(in byte) (int, error) {
	out := strings.IndexByte(chars, in)
	if out == -1 {
		return 0, fmt.Errorf("invalid byte %s", string(in))
	}
	return out % end, nil
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

		right, err := decodePlay(plays[1][0])
		if err != nil {
			log.Fatalf("invalid suggested play: %q", string(plays[1][0]))
		}

		result := beats(right, left)
		total += scoreShape(right) + scoreRound(result)
	}
	fmt.Println(total)
}

// returns 1 if left beats right, 0 if draw, -1 if left loses to right
func beats(left, right int) int {
	// draw if equal
	if right == left {
		return 0
	}
	// always operate on positive difference
	if left < right {
		left += end
	}
	return -1 + (2 * ((left - right) % 2))
}

func scoreShape(shape int) int { return shape + 1 }
func scoreRound(res int) int   { return (res + 1) * 3 }

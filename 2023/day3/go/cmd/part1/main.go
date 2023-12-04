package main

import (
	"day3"
	"fmt"
	"os"
	"strconv"
)

type Num struct {
	points []int
	str    string
}

func main() {
	var (
		boxen []Num
		syms  [][]int
	)

	l := day3.NewLexer(os.Stdin)
	for {
		p, t, s := l.Lex()

		if t == day3.NUM {
			boxen = append(boxen, Num{
				[]int{
					p.Column - 1,
					p.Column + len(s),
					p.Line - 1,
					p.Line + 1,
				},
				s,
			})
			continue
		}

		if t == day3.SYM {
			syms = append(syms, []int{p.Column, p.Line})
		}

		if t == day3.EOF {
			break
		}
	}

	inc := make(map[int]struct{})
	for _, sym := range syms {
		for i, box := range boxen {
			if contains(box.points, sym) {
				inc[i] = struct{}{}
			}
		}
	}

	var acc int
	for i := range inc {
		v, err := strconv.Atoi(boxen[i].str)
		if err != nil {
			panic(err)
		}
		acc += v
	}

	fmt.Println(acc)

}

func contains(box, point []int) bool {
	return box[0] <= point[0] &&
		box[1] >= point[0] &&
		box[2] <= point[1] &&
		box[3] >= point[1]
}

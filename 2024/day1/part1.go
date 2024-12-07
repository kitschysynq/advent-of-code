package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	l := make([]int, 0)
	r := make([]int, 0)

	f, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	s := bufio.NewScanner(f)

	for s.Scan() {
		var lin, rin int
		fmt.Sscanf(s.Text(), "%d %d", &lin, &rin)
		l = append(l, lin)
		r = append(r, rin)
	}

	sort.Ints(l)
	sort.Ints(r)

	var d int
	for i := range l {
		d += dist(l[i], r[i])
	}

	fmt.Println(d)
}

func dist(l, r int) int {
	if l < r {
		return r - l
	}

	return l - r
}

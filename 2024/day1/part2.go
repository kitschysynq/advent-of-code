package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	l := make([]int64, 0)
	r := make(map[int64]int64)

	f, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	s := bufio.NewScanner(f)

	for s.Scan() {
		var lin, rin int64
		fmt.Sscanf(s.Text(), "%d %d", &lin, &rin)
		l = append(l, lin)
		r[rin]++
	}

	var d int64
	for _, n := range l {
		d += n * r[n]
	}

	fmt.Println(d)
}

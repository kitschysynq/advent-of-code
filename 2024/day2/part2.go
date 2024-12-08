package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	s := bufio.NewScanner(f)

	var c int
	for s.Scan() {
		inStr := strings.Fields(s.Text())
		in := conv(inStr)
		if safe(in) || anySafe(in) {
			c++
		}
	}

	fmt.Println(c)
}

func conv(in []string) []int64 {
	var ints []int64
	for _, s := range in {
		cur, _ := strconv.ParseInt(s, 10, 64)
		ints = append(ints, cur)
	}
	return ints
}

func safe(in []int64) bool {
	return safeSkip(in, -1)
}

func anySafe(in []int64) bool {
	for i := range in {
		fmt.Printf("skipping %d\n", i)
		if safeSkip(in, i) {
			return true
		}
	}
	return false
}

func safeSkip(in []int64, skip int) bool {
	if len(in) < 2 {
		return false
	}

	if skip < 0 {
		skip = len(in) + 1
	}

	var dir bool
	for i := range len(in) - 2 {
		cur, next := in[i], in[i+1]
		if i >= skip {
			cur = in[i+1]
		}
		if i+1 >= skip {
			if i+1 == len(in)-1 {
				break
			}
			next = in[i+2]
		}

		delta := next - cur
		a, s := abs(delta)
		if a < 1 || a > 3 {
			return false
		}
		if i == 0 {
			dir = s
		} else if dir != s {
			return false
		}
	}
	return true
}

func abs(a int64) (int64, bool) {
	if a >= 0 {
		return a, true
	}
	return -a, false
}

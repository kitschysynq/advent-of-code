package main

import (
	"bufio"
	"fmt"
	"log"
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
		if safe(inStr) {
			c++
		}
	}

	fmt.Println(c)
}

func safe(in []string) bool {
	if len(in) < 2 {
		return false
	}

	var ints []int64
	for _, s := range in {
		cur, err := strconv.ParseInt(s, 10, 64)
		if err != nil {
			log.Printf("error parsing: %s", err.Error())
			return false
		}
		ints = append(ints, cur)
	}

	var dir bool
	for i := range ints[1:] {
		ints[i] = ints[i+1] - ints[i]
		a, s := abs(ints[i])
		if a < 1 || a > 3 {
			return false
		}
		if i == 0 {
			dir = s
		} else {
			if dir != s {
				return false
			}
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

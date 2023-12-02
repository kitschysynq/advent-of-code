package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

var (
	find    = regexp.MustCompile(`one|two|three|four|five|six|seven|eight|nine|\d`)
	replace = map[string]string{
		"one":   "1",
		"two":   "2",
		"three": "3",
		"four":  "4",
		"five":  "5",
		"six":   "6",
		"seven": "7",
		"eight": "8",
		"nine":  "9",
	}
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var sum, num int
	var err error
	for scanner.Scan() {
		t := scanner.Text()
		d := digits(t)
		if len(d) == 0 {
			continue
		}
		f, l := d[0], d[len(d)-1]
		if r, ok := replace[f]; ok {
			f = r
		}
		if r, ok := replace[l]; ok {
			l = r
		}
		num, err = strconv.Atoi(f + l)
		if err != nil {
			panic(err)
		}
		sum += num
	}
	fmt.Println(sum)
}

func digits(s string) []string {
	var d []string
	for idx := 0; idx < len(s); {
		st := s[idx:]
		l := find.FindStringIndex(st)
		if l == nil {
			return d
		}
		d0 := st[l[0]:l[1]]
		d = append(d, d0)
		idx += l[0] + 1
	}
	return d
}

package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

var digits = regexp.MustCompile(`\d`)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var sum, num int
	var err error
	for scanner.Scan() {
		d := digits.FindAllString(scanner.Text(), -1)
		if len(d) == 0 {
			continue
		}
		num, err = strconv.Atoi(d[0] + d[len(d)-1])
		if err != nil {
			panic(err)
		}
		sum += num
	}
	fmt.Println(sum)
}

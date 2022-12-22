package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

var nbr = regexp.MustCompile(`^(\w{4}): (\d+)`)
var exp = regexp.MustCompile(`^(\w{4}): (\w{4}) ([+-/*]) (\w{4})`)

var reg = make(map[string]Resolver)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		str := scanner.Text()

		name, res := parseResolver(str)
		reg[name] = res

	}
	fmt.Println(reg["root"].resolve())
}

func parseResolver(str string) (string, Resolver) {
	if m := nbr.FindStringSubmatch(str); m != nil {
		n, err := strconv.Atoi(m[len(m)-1])
		if err != nil {
			log.Fatalf("error parsing static resolver: %s", err.Error())
		}
		return m[1], staticResolver(n)
	}

	if m := exp.FindStringSubmatch(str); m != nil {
		return m[1], opResolver{
			a:  m[len(m)-3],
			b:  m[len(m)-1],
			op: ops[m[len(m)-2]],
		}
	}

	log.Fatalf("error parsing line: %s", str)
	return "", nil
}

type Resolver interface {
	resolve() int
}

type staticResolver int

func (s staticResolver) resolve() int { return int(s) }

type opResolver struct {
	a, b string
	op   op
}

func (o opResolver) resolve() int { return o.op(reg[o.a].resolve(), reg[o.b].resolve()) }

type op func(int, int) int

var ops = map[string]op{
	"+": func(a, b int) int { return a + b },
	"-": func(a, b int) int { return a - b },
	"/": func(a, b int) int { return a / b },
	"*": func(a, b int) int { return a * b },
}

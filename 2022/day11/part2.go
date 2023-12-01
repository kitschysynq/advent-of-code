package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"math/big"
	"os"
	"strconv"
	"strings"

	"github.com/Knetic/govaluate"
)

type Monkey struct {
	ID           int
	Items        []*big.Int
	testFn       func(big.Int) bool
	opFn         func(big.Int) big.Int
	TrueTgt      int
	FalseTgt     int
	itemsHandled int
}

func main() {
	var monkeys []*Monkey
	buf := bufio.NewReader(os.Stdin)

	for {
		if _, err := buf.Peek(1); err == io.EOF {
			break
		}
		monkeys = append(monkeys, parse(buf))
	}

	for i := 0; i < 1; i++ {
		for _, monkey := range monkeys {
			for _, item := range monkey.Items {
				monkey.itemsHandled++
				item = monkey.opFn(item)
				tgt := monkey.FalseTgt
				if monkey.testFn(item) {
					tgt = monkey.TrueTgt
				}
				monkeys[tgt].Items = append(monkeys[tgt].Items, item)
			}
			monkey.Items = monkey.Items[:0]
		}
	}

	var max, pen int
	for i, m := range monkeys {
		fmt.Printf("Monkey %d: handled %d items and has %v\n", i, m.itemsHandled, m.Items)
		if m.itemsHandled > max {
			pen, max = max, m.itemsHandled
			continue
		}
		if m.itemsHandled > pen {
			pen = m.itemsHandled
		}
	}

	b, m, p := big.NewInt(0), big.NewInt(int64(max)), big.NewInt(int64(pen))
	b.Mul(m, p)
	fmt.Printf("%s %s %s\n", m, p, b)
}

func parse(r *bufio.Reader) *Monkey {
	var m Monkey
	if n, err := fmt.Fscanf(r, "Monkey %d:\n", &m.ID); err != nil || n != 1 {
		log.Fatalf("parsed %d items for ID; got error %q", n, err.Error())
	}

	itemsStr, err := r.ReadString('\n')
	if err != nil {
		log.Fatalf("error reading starting items: %q", err.Error())
	}

	m.Items = parseItems(strings.TrimRight(itemsStr, "\n"))

	opStr, err := r.ReadString('\n')
	if err != nil {
		log.Fatalf("error reading operation: %q", err.Error())
	}

	m.opFn = parseOpFn(strings.TrimRight(opStr, "\n"))

	testStr, err := r.ReadString('\n')
	if err != nil {
		log.Fatalf("error reading test: %q", err.Error())
	}

	m.testFn = parseTestFn(strings.TrimRight(testStr, "\n"))

	if n, err := fmt.Fscanf(r, "    If true: throw to monkey %d\n", &m.TrueTgt); err != nil || n != 1 {
		log.Fatalf("parsed %d items for true target; got error %q", n, err.Error())
	}

	if n, err := fmt.Fscanf(r, "    If false: throw to monkey %d\n", &m.FalseTgt); err != nil || n != 1 {
		log.Fatalf("parsed %d items for false target; got error %q", n, err.Error())
	}

	blank, err := r.ReadString('\n')
	if err != nil && err != io.EOF {
		log.Fatalf("error skipping blank line: %q", err.Error())
	}

	if strings.TrimRight(blank, "\n") != "" {
		log.Fatalf("expected blank line, found %q", blank)
	}

	return &m
}

func parseItems(str string) []int {
	if !strings.HasPrefix(str, "  Starting items: ") {
		log.Fatalf("expected items line, found %q", str)
	}

	itemStrs := strings.Split(str[18:], ", ")
	var items []int
	for _, itemStr := range itemStrs {
		n, err := strconv.Atoi(itemStr)
		if err != nil {
			log.Fatalf("error parsing item %q: %q", itemStr, err.Error())
		}
		items = append(items, n)
	}
	return items
}

func parseOpFn(str string) func(int) int {
	if !strings.HasPrefix(str, "  Operation: new = ") {
		log.Fatalf("expected operation line, found %q", str)
	}

	expr, err := govaluate.NewEvaluableExpression(str[19:])
	if err != nil {
		log.Fatalf("error parsing op: %q", err.Error())
	}

	return func(i int) int {
		val, err := expr.Evaluate(map[string]interface{}{"old": i})
		if err != nil {
			log.Fatalf("error evaluating function: %q", err.Error())
		}

		valf64, ok := val.(float64)
		if !ok {
			log.Fatalf("error casting %T as float64", val)
		}

		return int(valf64)
	}
}

func parseTestFn(str string) func(int) bool {
	var d int64
	if n, err := fmt.Sscanf(str, "  Test: divisible by %d", &d); err != nil || n != 1 {
		log.Fatalf("parsed %d items; got error %q", n, err.Error())
	}
	j := big.NewInt(d)
	return func(i big.Int) bool {
		b := new(big.Int)
		b.Mod(i, j)
		return b.Cmp(big.NewInt(0)) == 0
	}
}

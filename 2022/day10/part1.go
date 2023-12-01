package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

type CPU struct {
	cycle int
	reg   int

	watchFn func(*CPU)
}

func NewCPU() *CPU {
	return &CPU{
		reg: 1,
	}
}

func (c *CPU) watch(fn func(*CPU)) {
	c.watchFn = fn
}

func (c *CPU) noop() { c.tick() }

func (c *CPU) addx(n int) {
	c.tick()
	c.tick()
	c.reg += n
}

func (c *CPU) tick() {
	c.cycle++
	if c.watchFn != nil {
		c.watchFn(c)
	}
}

var inst = regexp.MustCompile(`(noop|addx)( (-?\d+))?`)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	var (
		cnt    = 20
		step   = 40
		signal int
	)

	c := NewCPU()
	c.watch(func(c *CPU) {
		if c.cycle == cnt {
			signal += cnt * c.reg
			cnt += step
		}
	})

	for scanner.Scan() {
		parts := inst.FindStringSubmatch(scanner.Text())
		//fmt.Printf("%+q\n", parts)
		switch parts[1] {
		case "noop":
			c.noop()
		case "addx":
			n, err := strconv.Atoi(parts[3])
			if err != nil {
				log.Fatalf("error parsing operand: %s", err.Error())
			}
			c.addx(n)
		default:
			log.Fatalf("illegal instruction: %s", parts[1])
		}
	}
	fmt.Println(signal)
}

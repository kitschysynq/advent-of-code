package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
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

type Display struct {
	x, y   int
	pixels [][]byte
}

func NewDisplay(x, y int) *Display {
	pix := make([][]byte, y)
	for i := range pix {
		pix[i] = make([]byte, x)
	}
	return &Display{pixels: pix}
}

func (d *Display) draw(b byte) {
	d.pixels[d.y][d.x] = b
	d.x++
	if d.x >= len(d.pixels[0]) {
		d.x -= len(d.pixels[0])
		d.y++
	}
}

func (d *Display) String() string {
	b := new(strings.Builder)
	for _, row := range d.pixels {
		fmt.Fprintln(b, string(row))
	}
	return b.String()
}

var inst = regexp.MustCompile(`(noop|addx)( (-?\d+))?`)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	c := NewCPU()
	d := NewDisplay(40, 6)

	c.watch(func(c *CPU) {
		if c.reg >= d.x-1 && c.reg <= d.x+1 {
			d.draw('#')
			return
		}
		d.draw('.')
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
	fmt.Println(d.String())
}

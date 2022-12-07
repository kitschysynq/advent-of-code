package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
)

var (
	crate = regexp.MustCompile(`\[[A-Z]\] ?|    ?`)
	stack = regexp.MustCompile(` \d  ?`)
)

var stacks [][]byte

func main() {
	r := bufio.NewReader(os.Stdin)
	for {
		val, err := r.ReadString('\n')
		if err != nil {
			log.Fatalf("error reading stacks: %s", err.Error())
		}
		// check for labels first
		stackLabels := stack.FindAllString(val, -1)
		if stackLabels != nil {
			if len(stackLabels) != len(stacks) {
				log.Fatalf("found %d stacks but only %d labels", len(stacks), len(stackLabels))
			}
			break
		}
		items := crate.FindAllString(val, -1)
		if items == nil {
			log.Fatalf("expected crates, found %q", val)
		}
		processCrates(items)
	}

	_, _ = r.ReadString('\n')

	for {
		var cnt, fr, to int
		n, err := fmt.Fscanf(r, "move %d from %d to %d\n", &cnt, &fr, &to)
		if err != nil {
			break
		}
		if n != 3 {
			continue
		}

		fr--
		to--

		for i := 0; i < cnt; i++ {
			var c byte
			c, stacks[fr] = stacks[fr][0], stacks[fr][1:]
			stacks[to] = append([]byte{c}, stacks[to]...)
		}
	}

	for _, stack := range stacks {
		if len(stack) == 0 {
			fmt.Print("-")
		} else {
			fmt.Printf("%s", string(stack[0]))
		}
	}
	fmt.Println()
}

func processCrates(items []string) {
	if stacks == nil {
		stacks = make([][]byte, len(items))
	}

	for i, item := range items {
		if item[0] == '[' {
			stacks[i] = append(stacks[i], item[1])
		}
	}
}

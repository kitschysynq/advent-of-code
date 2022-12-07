package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	r := bufio.NewReader(os.Stdin)

	var buf []byte

	var idx int
	for len(buf) < 4 || hasDup(buf) {
		idx++
		b, err := r.ReadByte()
		if err == io.EOF || b == '\n' {
			log.Fatalf("no seq found by index %d", idx)
		}
		if err != nil {
			log.Fatalf("error reading byte: %s", err.Error())
		}

		if len(buf) < 4 {
			buf = append(buf, b)
			continue
		}

		buf = append(buf[1:], b)
	}

	fmt.Println(idx)
}

func atoIdx(b byte) int { return int(b) - 97 }
func hasDup(buf []byte) bool {
	for i, b1 := range buf[:len(buf)-1] {
		for _, b2 := range buf[i+1:] {
			if b1 == b2 {
				return true
			}
		}
	}
	return false
}

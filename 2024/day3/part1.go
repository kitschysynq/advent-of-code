package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
	"strconv"
)

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}

	buf := bufio.NewReader(f)

	var acc int64
	for {
		b, err := buf.ReadByte()
		if err != nil {
			if err == io.EOF {
				break
			}
			panic(err)
		}

		if b == 'm' {
			acc += parseMul(buf)
		}
	}

	fmt.Println(acc)
}

func parseMul(buf *bufio.Reader) int64 {
	needle := []byte("ul(") // exclude 'm' as it's already been matched
	var next []byte
	next, err := buf.Peek(3)
	if err != nil {
		return 0
	}
	if !bytes.Equal(next, needle) {
		return 0
	}
	buf.Discard(3)

	x, err := parseInt(buf)
	if err != nil {
		return 0
	}
	//log.Printf("x: %d", x)

	comma, err := buf.ReadByte()
	if err != nil || comma != ',' {
		buf.UnreadByte()
		return 0
	}

	y, err := parseInt(buf)
	if err != nil {
		return 0
	}
	//log.Printf("y: %d", y)

	rparen, err := buf.ReadByte()
	if err != nil || rparen != ')' {
		return 0
	}

	return x * y
}

func parseInt(buf *bufio.Reader) (int64, error) {
	var val []byte
	for {
		b, err := buf.ReadByte()
		if err != nil {
			return 0, err
		}

		if !isDigit(b) {
			buf.UnreadByte()
			break
		}
		val = append(val, b)
	}

	return strconv.ParseInt(string(val), 10, 32)
}

func isDigit(b byte) bool {
	return bytes.IndexByte([]byte("0123456789"), b) != -1
}

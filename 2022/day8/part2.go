package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	var (
		grid   []byte
		rowLen int
	)

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		row := scanner.Bytes()
		if rowLen == 0 {
			rowLen = len(row)
		} else if rowLen != len(row) {
			log.Fatalf("all input rows must be the same length")
		}
		grid = append(grid, row...)
	}

	var max int
	for idx := range grid {
		score := 1
		score *= scoreW(grid, rowLen, idx)
		score *= scoreE(grid, rowLen, idx)
		score *= scoreN(grid, rowLen, idx)
		score *= scoreS(grid, rowLen, idx)
		if score > max {
			max = score
		}
	}

	fmt.Println(max)
}

func scoreW(grid []byte, rowLen, idx int) int {
	var score int
	if idx%rowLen == 0 {
		return 0
	}
	for i := idx - 1; (i+1)%rowLen > 0; i-- {
		score++
		if grid[i] >= grid[idx] {
			return score
		}
	}
	return score
}

func scoreE(grid []byte, rowLen, idx int) int {
	var score int
	if (idx+1)%rowLen == 0 {
		return 0
	}
	for i := idx + 1; i%rowLen > 0; i++ {
		score++
		if grid[i] >= grid[idx] {
			return score
		}
	}
	return score
}

func scoreN(grid []byte, rowLen, idx int) int {
	var score int
	if idx-rowLen < 0 {
		return 0
	}
	for i := idx - rowLen; i >= 0; i -= rowLen {
		score++
		if grid[i] >= grid[idx] {
			return score
		}
	}
	return score
}

func scoreS(grid []byte, rowLen, idx int) int {
	var score int
	if idx+rowLen >= len(grid) {
		return 0
	}
	for i := idx + rowLen; i < len(grid); i += rowLen {
		score++
		if grid[i] >= grid[idx] {
			return score
		}
	}
	return score
}

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
		vis    []bool
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

		var (
			max byte
			pos int
		)
		rowVis := make([]bool, len(row))
		for i, b := range row {
			if b > max {
				rowVis[i] = true
				max = b
				pos = i
			}
		}
		max = 0
		for i := len(row) - 1; i > pos; i-- {
			if row[i] > max {
				rowVis[i] = true
				max = row[i]
			}
		}
		vis = append(vis, rowVis...)
	}

	for i := 0; i < rowLen; i++ {
		var (
			max byte
			pos int
		)
		for j := 0; j < len(grid)/rowLen; j++ {
			idx := (j * rowLen) + i
			if grid[idx] > max {
				vis[idx] = true
				max = grid[idx]
				pos = j
			}
		}
		max = 0
		for j := len(grid)/rowLen - 1; j > pos; j-- {
			idx := (j * rowLen) + i
			if grid[idx] > max || j == len(grid)/rowLen-1 {
				vis[idx] = true
				max = grid[idx]
			}
		}
	}

	count := 0
	for _, v := range vis {
		if v {
			count++
		}
	}
	fmt.Println(count)
}

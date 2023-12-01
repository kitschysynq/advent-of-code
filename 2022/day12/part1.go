package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/yourbasic/graph"
)

func main() {
	var gridBytes [][]byte
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		gridBytes = append(gridBytes, scanner.Bytes())
	}

	g, s, e := parse(gridBytes)
	path, dist := graph.ShortestPath(g, s, e)

	for _, pos := range path {
		x, y := xy(len(gridBytes[0]), pos)
		gridBytes[y][x] = 'X'
	}

	for _, row := range gridBytes {
		fmt.Println(string(row))
	}
	fmt.Println(path)
	fmt.Println(len(path) - 1)
	fmt.Println(dist)
}

func parse(grid [][]byte) (*graph.Mutable, int, int) {
	w, h := len(grid[0]), len(grid)
	g := graph.New(w * h)
	var s, e int
	for y, row := range grid {
		for x, cell := range row {
			if cell == 'S' {
				s = idx(w, x, y)
				grid[y][x] = 'a'
				fmt.Printf("S: x %d y %d idx %d\n", x, y, s)
			}
			if cell == 'E' {
				e = idx(w, x, y)
				grid[y][x] = 'z'
				fmt.Printf("E: x %d y %d idx %d\n", x, y, e)
			}
		}
	}
	for y, row := range grid {
		for x, cell := range row {
			if x-1 >= 0 {
				diff := int(grid[y][x-1]) - int(cell)
				if diff <= 1 {
					fmt.Printf("Left of (%d, %d): %s %s delta %d\n", x, y, string(grid[y][x-1]), string(cell), diff)
					g.Add(idx(w, x-1, y), idx(w, x, y))
				}
			}
			if x+1 < w {
				diff := int(grid[y][x+1]) - int(cell)
				if diff <= 1 {
					fmt.Printf("Right of (%d, %d): %s %s delta %d\n", x, y, string(grid[y][x+1]), string(cell), diff)
					g.Add(idx(w, x+1, y), idx(w, x, y))
				}
			}
			if y-1 >= 0 {
				diff := int(grid[y-1][x]) - int(cell)
				if diff <= 1 {
					fmt.Printf("North of (%d, %d): %s %s delta %d\n", x, y, string(grid[y-1][x]), string(cell), diff)
					g.Add(idx(w, x, y-1), idx(w, x, y))
				}
			}
			if y+1 < h {
				diff := int(grid[y+1][x]) - int(cell)
				if diff <= 1 {
					fmt.Printf("South of (%d, %d): %s %s delta %d\n", x, y, string(grid[y+1][x]), string(cell), diff)
					g.Add(idx(w, x, y+1), idx(w, x, y))
				}
			}
		}
	}
	return g, s, e
}

func idx(w, x, y int) int      { return y*w + x }
func xy(w, idx int) (int, int) { return idx % w, idx / w }

package day6

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"strings"
)

func Roots(ti, di int) (int, int) {
	t, d := float64(ti), float64(di)
	r := math.Sqrt(t*t - 4*d)
	min, max := (t-r)/2, (t+r)/2
	if math.Trunc(min) == min {
		min++
	}
	if math.Trunc(max) == max {
		max--
	}
	return int(math.Ceil(min)), int(math.Floor(max))
}

func LabelledNums(text, label string) []int {
	t, found := strings.CutPrefix(text, label)
	if !found {
		panic("invalid times")
	}

	var nums []int
	r := strings.NewReader(t)
	for {
		var num int
		_, err := fmt.Fscanf(r, "%d", &num)
		if err != nil {
			if err == io.EOF {
				return nums
			}
			panic(err)
		}
		nums = append(nums, num)
	}
}

type Races struct {
	Time []int
	Dist []int
}

func ParseRaces(r io.Reader) Races {
	s := bufio.NewScanner(r)
	if !s.Scan() {
		panic("no times in stream")
	}
	t := LabelledNums(s.Text(), "Time:")
	if !s.Scan() {
		panic("no distances in stream")
	}
	d := LabelledNums(s.Text(), "Distance:")
	return Races{t, d}
}

func ParseRaces2(r io.Reader) Races {
	s := bufio.NewScanner(r)
	if !s.Scan() {
		panic("no times in stream")
	}
	tStr := strings.ReplaceAll(s.Text(), " ", "")
	t := LabelledNums(tStr, "Time:")
	if !s.Scan() {
		panic("no distances in stream")
	}
	dStr := strings.ReplaceAll(s.Text(), " ", "")
	d := LabelledNums(dStr, "Distance:")
	return Races{t, d}
}

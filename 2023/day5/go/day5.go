package day5

import "fmt"

type Map struct {
	//Seeds    []int
	ItemMaps []*ItemMap
}

func (m *Map) To(kind string, num int) (string, int, bool) {
	for _, im := range m.ItemMaps {
		if im.From == kind {
			return im.To, im.Map(num), true
		}
	}
	return "", 0, false
}

type ItemMap struct {
	From   string
	To     string
	Ranges []Range
}

func (i ItemMap) Map(num int) int {
	for _, r := range i.Ranges {
		if r.Contains(num) {
			return r.Map(num)
		}
	}
	return num
}

func (i ItemMap) GoString() string {
	return fmt.Sprintf("{From: %s, To: %s, Ranges: %v}", i.From, i.To, i.Ranges)
}

type Range struct {
	From int
	To   int
	Len  int
}

func (r Range) GoString() string {
	return fmt.Sprintf("{%d %d %d}", r.To, r.From, r.Len)
}

func (r Range) Contains(i int) bool {
	return r.From <= i &&
		i < r.From+r.Len
}

func (r Range) Map(i int) int {
	return i - r.From + r.To
}

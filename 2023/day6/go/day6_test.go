package day6

import (
	"fmt"
	"testing"
)

func TestRoots(t *testing.T) {
	tests := []struct {
		t int
		d int
	}{
		{7, 9},
		{15, 40},
		{30, 200},
	}

	for _, test := range tests {
		fmt.Println(Roots(test.t, test.d))
	}
}

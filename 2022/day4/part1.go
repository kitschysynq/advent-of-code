package main

import (
	"fmt"
	"io"
	"log"
)

func main() {
	var count int
	for {
		var s1, e1, s2, e2 int
		if _, err := fmt.Scanf("%d-%d,%d-%d", &s1, &e1, &s2, &e2); err != nil {
			if err == io.EOF {
				break
			}
			log.Fatalf(err.Error())
		}
		if contains(s1, e1, s2, e2) || contains(s2, e2, s1, e1) {
			count++
		}
	}
	fmt.Println(count)
}

func contains(s1, e1, s2, e2 int) bool {
	return s1 <= s2 && e1 >= e2
}

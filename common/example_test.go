package common

import (
	"fmt"
	"testing"
)

func TestTemp(t *testing.T) {
	// testing.Benchmark(f func(b *B))
	var m = map[string]int{
		"A": 21,
		"B": 22,
		"C": 23,
		"D": 23,
		"E": 23,
		"F": 23,
		"G": 23,
		"H": 23,
	}
	for i := 0; i < 10; i++ {

		counter := 0
		for k, _ := range m {
			if counter == 0 {
				delete(m, "A")
			}
			t.Logf("%v,%v", k, &k)
			m["I"] = 33

			counter++
		}
		fmt.Println(counter)
	}

}

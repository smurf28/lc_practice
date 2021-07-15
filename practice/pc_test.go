package practice

import (
	"fmt"
	"sync"
	"testing"
)

func TestPc(t *testing.T) {
	goods := make(chan int, 5)
	var wg sync.WaitGroup

	go func() {
		for i := 0; i < 3; i++ {
			wg.Add(1)
			go product(&wg, goods, i)
		}
		wg.Wait()
		fmt.Printf("product goods closed\n")
		close(goods)
	}()

	var c sync.WaitGroup
	for i := 0; i < 5; i++ {
		c.Add(1)
		go consumer(&c, goods, i)
	}

	c.Wait()
}


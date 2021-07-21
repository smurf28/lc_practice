package practice

import (
	"context"
	"fmt"
	"sync"
	"testing"
)

var pool = 100

func print1(ch chan int, wg *sync.WaitGroup) {
	for i := 1; i < pool; i ++ {
		ch <- i
		if i%2 == 1 {
            fmt.Println("groutine-1:", i)
        }
	}
	wg.Done()
}

func print2(ch chan int, wg *sync.WaitGroup) {
	for i := 1; i <= pool; i ++{
		<-ch
		if i%2 == 0 {
            fmt.Println("groutine-2:", i)
        }
	}
	wg.Done()
}

func print3(ch chan int, ctx context.Context) {
	for i := 1; i < pool; i += 2 {
		<-ch
		fmt.Println(i)
	}

}

func print4(ch chan int, ctx context.Context) {
	for i := 0; i < pool; i += 2 {
		ch <- i
		fmt.Println(i)
	}
}

func TestPrint(t *testing.T) {
	ch := make(chan int)
	wg := &sync.WaitGroup{}
	wg.Add(2)
	go print1(ch, wg)
	go print2(ch, wg)
	wg.Wait()
}

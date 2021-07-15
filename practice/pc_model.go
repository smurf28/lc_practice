// 生产者消费者模型
package practice

import (
	"context"
	"fmt"
	"sync"
	"time"
)

// func main() {
// 	//go f1()
// 	//go f2()
// 	//go f3()
// 	//go f4()
// 	//go f5()
// 	go f6()

// 	fmt.Println("let's go")
// 	select {}
// }

// normal
func f1() {
	c := make(chan int, 4) // buffer channel
	go func(c chan int) {
		var i = -1
		for {
			i += 1
			ctx, _ := context.WithTimeout(context.Background(), time.Duration(5*time.Second)) // if sending to channel will take too long, then timeout

			fmt.Printf("\ncurrent num is : %d\n", i)
			select {
			case c <- i:
				fmt.Printf("sent success : %d\n", i)
			case <-ctx.Done():
				fmt.Printf("sent failed : %d, since timeout\n", i)
			}
		}
	}(c)

	go func(c chan int) {
		for e := range c {
			fmt.Printf("receive : %d\n", e)
		}
	}(c)
}

// timeout since channel buffer full-blocking
func f2() {
	c := make(chan int, 4)
	go func(c chan int) {
		var i = -1
		for {
			i += 1
			ctx, _ := context.WithTimeout(context.Background(), time.Duration(5*time.Second))

			fmt.Printf("\ncurrent num is : %d\n", i)
			select {
			case c <- i:
				fmt.Printf("sent success : %d\n", i)
			case <-ctx.Done():
				fmt.Printf("sent failed : %d, since timeout\n", i)
			}
		}
	}(c)
}

// blocking but without timeout since go through default directly
// will lost data, if going to `default`
func f3() {
	c := make(chan int, 4)
	go func(c chan int) {
		var i = -1
		for {
			i += 1
			ctx, _ := context.WithTimeout(context.Background(), time.Duration(5*time.Second))

			fmt.Printf("\ncurrent num is : %d\n", i)
			select {
			case c <- i:
				fmt.Printf("sent success : %d\n", i)
			case <-ctx.Done():
				fmt.Printf("sent failed : %d, since timeout\n", i)
			default:
				fmt.Printf("sent unknown : %d, since default\n", i)
			}
		}
	}(c)

	go func(c chan int) {
		for {
			num := <-c
			fmt.Printf("receive : %d\n", num)
		}
	}(c)
}

// 有部分会跑到default去，导致没有sent，也就没有receive
func f4() {
	c := make(chan int, 7)
	go func(c chan int) {
		var i = -1
		for {
			i += 1
			ctx, _ := context.WithTimeout(context.Background(), time.Duration(5*time.Second))

			fmt.Printf("\ncurrent num is : %d\n", i)
			select {
			case c <- i:
				time.Sleep(1 * time.Second) //每秒产生一个
				fmt.Printf("sent success : %d\n", i)
			case <-ctx.Done():
				fmt.Printf("sent failed : %d, since timeout\n", i)
			default:
				time.Sleep(1 * time.Second) //每秒产生一个
				fmt.Printf("sent unknown : %d, since default\n", i)
			}
		}
	}(c)

	go func(c chan int) {
		for e := range c {
			time.Sleep(2 * time.Second) //每2秒才消费一个
			fmt.Printf("receive : %d\n", e)
		}
	}(c)
}

// 直接丢掉被buff的部分，直到consumer把buffer降低，才能再塞入一个
func f5() {
	c := make(chan int, 7)
	go func(c chan int) {
		var i = -1
		for {
			i += 1
			ctx, _ := context.WithTimeout(context.Background(), time.Duration(5*time.Second))

			select {
			case c <- i:
				fmt.Printf("sent success : %d\n", i)
			case <-ctx.Done():
				fmt.Printf("sent failed : %d, since timeout\n", i)
			default:
			}
		}
	}(c)

	go func(c chan int) {
		for e := range c {
			time.Sleep(2 * time.Second) //每2秒才消费一个
			fmt.Printf("receive : %d\n", e)
		}
	}(c)
}

// 不丢弃，blocking
func f6() {
	c := make(chan int, 7)
	go func(c chan int) {
		var i = -1
		for {
			i += 1
			ctx, _ := context.WithTimeout(context.Background(), time.Duration(5*time.Second))

			select {
			case c <- i:
				fmt.Printf("sent success : %d\n", i)
			case <-ctx.Done():
				fmt.Printf("sent failed : %d, since timeout\n", i)
				//default:
			}
		}
	}(c)

	go func(c chan int) {
		for e := range c {
			time.Sleep(2 * time.Second) //每2秒才消费一个
			fmt.Printf("receive : %d\n", e)
		}
	}(c)
}
func product(wg *sync.WaitGroup, good chan int, index int) {
	for i := 1; i <= 10; i++ {
		good <- index*10 + i
		fmt.Printf("product-%d good %d\n", index, i)
	}
	fmt.Printf("product-%d done\n", index)
	wg.Done()
}

func consumer(wg *sync.WaitGroup, good chan int, index int) {
	for {
		count, ok := <-good
		if !ok {
			fmt.Printf("consumer-%d closed\n", index)
			wg.Done()
			break
		}
		fmt.Printf("consumer-%d count %d\n", index, count)
	}
}

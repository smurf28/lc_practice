package practice

import (
	"context"
	"fmt"
	"testing"
	"time"
)

// 大切片 查找topk

func task(datas []int, target int, ctx context.Context, resChan chan bool) {
	for _, data := range datas {
		select {
		case <-ctx.Done():
			return
		default:
		}
		if data == target {
			resChan <- true
			return
		}
	}
	
}

func Solution(datas []int, target int) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	resChan := make(chan bool)

	// 分段长度
	segmentLen := 2

	dataLen := len(datas)
	for i := 0; i < dataLen; i += segmentLen {
		end := i + segmentLen
		if end >= dataLen {
			end = dataLen - 1
		}

		go task(datas[i:end], target, ctx, resChan)
	}

	select {
	case <-resChan:
		ant := fmt.Sprintf("find target: %d", target)
		fmt.Println(ant)
		cancel()
		return
	case <-ctx.Done():
		fmt.Println("can not find")
	}
}

func Test_find(t *testing.T) {
	datas := []int{1, 2, 3, 4, 5, 6, 7, 8, 3}
	Solution(datas, 7)
	
	// case2 TODO找不到的情况, 考虑如何退出fatal error: all goroutines are asleep - deadlock!
	datas2 := []int{1, 2, 3, 4, 5, 6, 7, 8, 3}
	Solution(datas2, 77)
}

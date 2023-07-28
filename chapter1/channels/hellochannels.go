package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func printer(ch chan int) {
	// 从 channel 中获取值
	// 如果没有获取到值，就等待，直到通道关闭
	for i := range ch {
		fmt.Printf("Received %d ", i)
	}

	// 通知计算器 -1
	wg.Done()
}

func main() {
	// 创建一个无缓冲 channel
	c := make(chan int)

	// 启动一个 goroutine
	go printer(c)

	// 等待计数加 1
	wg.Add(1)

	for i := 1; i <= 10; i++ {
		// 开始给 channel 传递值
		c <- i
	}

	// 关闭 channel
	close(c)

	// 计数器清零
	wg.Wait()
}

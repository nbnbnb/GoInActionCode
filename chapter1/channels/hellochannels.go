package main

import (
	"fmt"
	"sync"
)

// 创建一个等待计数器
var wg sync.WaitGroup

func printer(ch chan int) {
	// 从 channel 中获取值
	// 如果没有获取到值，就等待，直到通道关闭
	for i := range ch {
		fmt.Printf("Received %d \n", i)
	}

	// 等待计数器 - 1
	wg.Done()
}

func main() {
	// 创建一个无缓冲 channel
	c := make(chan int)

	// 启动一个 goroutine
	// 传递创建的 channel
	go printer(c)

	// 设置等待计数器 + 1
	wg.Add(1)

	for i := 1; i <= 10; i++ {
		// 开始给 channel 传递值
		c <- i
	}

	// 关闭 channel
	close(c)

	// 等待计数器 = 0 时返回
	wg.Wait()
}

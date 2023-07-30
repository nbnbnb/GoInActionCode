// 这个示例程序展示如何在程序里造成竞争状态
// 实际上不希望出现这种情况
package main

import (
	"fmt"
	"runtime"
	"sync"
)

var (
	// counter 是所有 goroutine 都要增加其值的变量
	counter int

	// wg 用来等待程序结束
	wg sync.WaitGroup
)

func main() {
	// 计数加 2，表示要等待两个 goroutine
	wg.Add(2)

	// 创建两个 goroutine
	go incCounter(1)
	go incCounter(2)

	// 等待 goroutine 结束
	wg.Wait()
	// 期望输出是  counter == 2000
	// 实际不是，因为 counter++ 不是一个原子操作
	fmt.Println("Final Counter:", counter)
}

// incCounter 增加 counter 变量的值
func incCounter(id int) {
	// 在函数退出时调用 Done 来通知 main 函数工作已经完成
	defer wg.Done()

	for count := 0; count < 1000; count++ {
		// 捕获 counter 的值
		// 这里是将 counter 变量的副本存入一个叫作 value 的本地变量
		value := counter

		// yield the thread and be placed back in queue
		// 当前 goroutine 从线程退出，并放回到队列
		runtime.Gosched()

		// 增加本地 value 变量的值
		value++

		// 将该值保存回 counter
		// 将这个新值存回到 counter 变量
		counter = value
	}
}

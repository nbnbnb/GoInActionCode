// 这个示例程序展示 goroutine 调度器是如何在单个线程上切分时间片的
package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

func main() {
	// 分配一个逻辑处理器给调度器使用
	runtime.GOMAXPROCS(1)

	wg.Add(2)

	fmt.Println("Create Goroutines")

	// 创建两个 goroutine
	go printPrime("A")
	go printPrime("B")

	fmt.Println("Waiting To Finish")

	// 等待 goroutine 结束
	wg.Wait()

	fmt.Println("Terminating Program")
}

func printPrime(prefix string) {
	// 在函数退出时调用 Done 来通知 main 函数工作已经完成
	defer wg.Done()

	// 显示 5000 以内的素数值
	// 此处类似 goto 语法
next:
	for outer := 2; outer < 5000; outer++ {
		for inner := 2; inner < outer; inner++ {
			if outer%inner == 0 {
				continue next
			}
		}
		fmt.Printf("%s:%d\n", prefix, outer)
	}

	fmt.Println("Completed", prefix)
}

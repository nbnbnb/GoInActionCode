// 这个示例程序展示如何创建 goroutine 以及调度器的行为
package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	// 分配一个逻辑处理器给调度器使用
	runtime.GOMAXPROCS(1)

	// wg 用来等待程序完成
	var wg sync.WaitGroup

	// 计数加 2 表示要等待两次 wg.Done
	wg.Add(2)

	fmt.Println("Start Goroutines")

	// 声明一个匿名函数，并创建一个 goroutine
	go func() {
		// 注意 defer 的用法
		// 在函数退出时调用 Done 来通知 main 函数工作已经完成
		defer wg.Done()

		// 显示字母表 3 次
		for count := 0; count < 3; count++ {
			// 输出小写字母表
			for char := 'a'; char < 'a'+26; char++ {
				fmt.Printf("%c ", char)
			}
		}
	}()

	// 与上面的 goroutine 一样
	go func() {
		defer wg.Done()
		for count := 0; count < 3; count++ {
			// 输出大写字母表
			for char := 'A'; char < 'A'+26; char++ {
				fmt.Printf("%c ", char)
			}
		}
	}()

	// 立马就会执行到这里，因为上面的 goroutine 是异步执行的
	fmt.Println("Waiting To Finish")

	// 等待 goroutine 结束：wg.Done 需要被调用两次
	wg.Wait()

	fmt.Println("\nTerminating Program")
}

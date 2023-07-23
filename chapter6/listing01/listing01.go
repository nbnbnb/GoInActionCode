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
	// 计数加 2，表示要等待两个 goroutine
	var wg sync.WaitGroup
	wg.Add(2)

	fmt.Println("Start Goroutines")

	// 声明一个匿名函数，并创建一个 goroutine
	go func() {
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

	// 等待 goroutine 结束
	fmt.Println("Waiting To Finish")
	wg.Wait()

	// 第一个 goroutine 完成所有显示需要花时间太短了，以至于在调度器切换到第二个 goroutine 之前，就完成了所有任务
	// 这也是为什么会看到先输出了所有的大写字母，之后才输出小写字母
	fmt.Println("\nTerminating Program")
}

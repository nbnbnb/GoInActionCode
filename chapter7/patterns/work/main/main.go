// 这个示例程序展示如何使用 work 包创建一个 goroutine 池并完成工作
package main

import (
	"GoInActionCode/chapter7/patterns/work"
	"log"
	"sync"
	"time"
)

// names 提供了一组用来显示的名字
var names = []string{
	"steve",
	"bob",
	"mary",
	"therese",
	"jason",
}

// namePrinter 使用特定方式打印名字
type namePrinter struct {
	name string
}

// Task 实现 Worker 接口
func (m *namePrinter) Task() {
	log.Println(m.name)
	// 等待 100ms
	time.Sleep(time.Millisecond * 100)
}

func main() {
	// 使用 2 个 goroutine 来创建工作池
	pool := work.New(2)

	var wg sync.WaitGroup

	// 总共迭代 500 个 goroutine 来完成这些任务
	wg.Add(100 * len(names))

	for i := 0; i < 100; i++ {
		// 迭代 names 切片
		for _, name := range names {
			// 创建一个 namePrinter 并提供指定的名字
			np := namePrinter{
				name: name,
			}

			go func() {
				// 将任务提交执行
				// 当 Run 返回时我们就知道任务已经处理完成
				pool.Run(&np)
				wg.Done()
			}()
		}
	}

	wg.Wait()

	//  让工作池停止工作，等待所有现有的工作完成
	pool.Shutdown()
}

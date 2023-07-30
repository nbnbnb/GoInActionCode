// 这个示例程序展示如何使用有缓冲的通道和固定数目的 goroutine 来处理一堆工作
package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// 注意：常量包的命名方式
const (
	// 要使用的 goroutine 的数量
	numberGoroutines = 4
	// 要处理的工作的数量
	taskLoad = 10
)

// wg 用来等待程序完成
var wg sync.WaitGroup

// init 初始化包
// Go 语言运行时会在其他代码执行之前优先执行这个函数
func init() {
	// 初始化随机数种子
	rand.Seed(time.Now().Unix())
}

func main() {
	// 创建一个有缓冲的通道来管理工作
	// 第二个值表示缓冲区大小
	tasks := make(chan string, taskLoad)

	wg.Add(numberGoroutines)

	for gr := 1; gr <= numberGoroutines; gr++ {
		// 启动 goroutine
		go worker(tasks, gr)
	}

	// 增加一组要完成的工作
	for post := 1; post <= taskLoad; post++ {
		// 向通道中写入数据
		tasks <- fmt.Sprintf("Task : %d", post)
	}

	// 当所有工作都处理完时关闭通道以便所有 goroutine 退出
	// 当通道关闭后， goroutine 依旧可以从通道接收数据，但是不能再向通道里发送数据
	// 从一个已经关闭且没有数据的通道里获取数据，总会立刻返回，并返回一个通道类型的零值

	// 相较于无缓冲的通道
	// 有缓冲的通道，一定要调用 close 函数 来关闭通道
	close(tasks)

	// 等待所有工作完成
	wg.Wait()
}

// worker 作为 goroutine 启动来处理从有缓冲的通道传入的工作
func worker(tasks chan string, worker int) {
	defer wg.Done()

	for {
		// 等待分配工作
		// 从通道中获取数据（每个 goroutine 都会在这个位置阻塞，等待从通道中接收数据）
		task, ok := <-tasks
		// 从一个已经关闭且没有数据的通道里获取数据，总会立刻返回，并返回一个通道类型的零值
		// 所以，如果 ok 的值为 false，表示通道已经关闭
		if !ok {
			// 这意味着通道已经空了，并且已被关闭
			fmt.Printf("Worker: %d : Shutting Down\n", worker)
			return
		}

		// 显示我们开始工作了
		fmt.Printf("Worker: %d : Started %s\n", worker, task)

		// 随机等一段时间来模拟工作
		sleep := rand.Int63n(100)
		time.Sleep(time.Duration(sleep) * time.Millisecond)

		// 显示我们完成了工作
		fmt.Printf("Worker: %d : Completed %s\n", worker, task)
	}
}

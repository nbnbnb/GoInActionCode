// 这个示例程序展示如何用无缓冲的通道来模拟 4 个 goroutine 间的接力比赛
package main

import (
	"fmt"
	"sync"
	"time"
)

// wg 用来等待程序结束
var wg sync.WaitGroup

func main() {
	// 创建一个无缓冲的通道
	baton := make(chan int)

	// 等待最后一个接力者完成
	wg.Add(1)

	//  第一位跑步者持有接力棒
	go Runner(baton)

	// 向通道发送数据
	// 开始比赛
	baton <- 1

	// 等待比赛结束（最后一位跑步者）
	wg.Wait()
}

// Runner 模拟接力比赛中的一位跑步者
func Runner(baton chan int) {
	var newRunner int

	// 程序暂停，从通道接收数据
	// 等待接力棒
	// 从通道读取数据
	runner := <-baton

	// 开始绕着跑道跑步
	fmt.Printf("Runner %d Running With Baton\n", runner)

	// 创建下一位跑步者
	if runner != 4 {
		newRunner = runner + 1
		fmt.Printf("Runner %d To The Line\n", newRunner)
		// 创建一个新的 goroutine 来跑步
		go Runner(baton)
	}

	// 围绕跑道跑
	time.Sleep(100 * time.Millisecond)

	// 比赛结束了吗？
	if runner == 4 {
		fmt.Printf("Runner %d Finished, Race Over\n", runner)
		// 最后一个跑步者完成
		wg.Done()
		return
	}

	fmt.Printf("Runner %d Exchange With Runner %d\n", runner, newRunner)

	// 将接力棒交给下一位跑步者
	// 数据发送到通道
	baton <- newRunner
}

// 这个示例程序展示如何用无缓冲的通道来模拟 2 个 goroutine 间的网球比赛
package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// wg 用来等待程序结束
var wg sync.WaitGroup

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	// 创建一个无缓冲的通道
	court := make(chan int)

	// 计数加 2，表示要等待两个 goroutine
	wg.Add(2)

	// 启动两个选手
	go player("Nadal", court)
	go player("Djokovic", court)

	// 将 1 发送到通道中
	// 发球
	court <- 1

	// 等待游戏结束
	wg.Wait()
}

// player 模拟一个选手在打网球
func player(name string, court chan int) {
	// 在函数退出时调用 Done 来通知 main 函数工作已经完成
	defer wg.Done()

	// 这是一个无限循环
	for {
		// 程序暂停，从通道中等待值
		// 等待球被击打过来
		ball, ok := <-court

		// 如果通道关闭了，ok == false
		if !ok {
			// 如果通道被关闭，我们就赢了
			fmt.Printf("Player %s Won\n", name)
			return
		}

		// 选随机数，然后用这个数来判断我们是否丢球
		n := rand.Intn(100)

		// 如果是 13 的倍数，表示我们丢球了
		if n%13 == 0 {

			fmt.Printf("Player %s Missed\n", name)

			// 关闭通道，表示我们输了
			close(court)

			return
		}

		// 显示击球数
		fmt.Printf("Player %s Hit %d\n", name, ball)

		// 将击球数加 1
		ball++

		// 向通道里面发送一个值
		// 将球打向对手
		court <- ball
	}
}

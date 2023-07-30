// work 包管理一个 goroutine 池来完成工作
package work

import "sync"

// 定义一个接口
// Worker 必须满足接口类型，才能使用工作池
type Worker interface {
	Task()
}

// Pool 提供一个 goroutine 池， 这个池可以完成任何已提交的 Worker 任务
type Pool struct {
	// 一个通道
	work chan Worker
	wg   sync.WaitGroup
}

// New 创建一个新工作池
func New(maxGoroutines int) *Pool {
	pool := Pool{
		// 初始化无缓冲通道
		work: make(chan Worker),
	}

	// 最多 maxGoroutines 个 goroutine 来完成这些任务
	pool.wg.Add(maxGoroutines)

	// 创建 maxGoroutines 个 goroutine
	for i := 0; i < maxGoroutines; i++ {
		go func() {
			// 程序阻塞在这里，直到有任务传入
			// 从通道中获取 work
			// 注意：此处是用迭代器的方式
			for work := range pool.work {
				// 执行 Task 方法
				work.Task()
			}
			// 设置 goroutine 执行完毕
			pool.wg.Done()
		}()
	}

	// 返回指针
	return &pool
}

// Run 提交工作到工作池
func (pool *Pool) Run(work Worker) {
	// 将 work 发送到通道中
	pool.work <- work
}

// Shutdown 等待所有 goroutine 停止工作
func (pool *Pool) Shutdown() {
	// 关闭通道
	close(pool.work)
	// 所有的 goroutine 执行完毕
	pool.wg.Wait()
}

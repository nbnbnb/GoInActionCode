// work 包管理一个 goroutine 池来完成工作
package work

import "sync"

// Worker 必须满足接口类型，才能使用工作池
type Worker interface {
	Task()
}

// Pool 提供一个 goroutine 池， 这个池可以完成任何已提交的 Worker 任务
type Pool struct {
	work chan Worker
	wg   sync.WaitGroup
}

// New 创建一个新工作池
func New(maxGoroutines int) *Pool {
	pool := Pool{
		work: make(chan Worker),
	}

	// 最多 maxGoroutines 个 goroutine 来完成这些任务
	pool.wg.Add(maxGoroutines)

	// 创建 maxGoroutines 个 goroutine
	for i := 0; i < maxGoroutines; i++ {
		go func() {
			for work := range pool.work {
				// 执行 Task 方法
				work.Task()
			}
			pool.wg.Done()
		}()
	}

	return &pool
}

// Run 提交工作到工作池
func (pool *Pool) Run(work Worker) {
	pool.work <- work
}

// Shutdown 等待所有 goroutine 停止工作
func (pool *Pool) Shutdown() {
	close(pool.work)
	pool.wg.Wait()
}

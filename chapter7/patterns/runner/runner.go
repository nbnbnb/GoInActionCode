// runner 包管理处理任务的运行和生命周期
package runner

import (
	"errors"
	"os"
	"os/signal"
	"time"
)

// Runner 在给定的超时时间内执行一组任务，并且在操作系统发送中断信号时结束这些任务
type Runner struct {
	// interrupt 通道报告从操作系统发送的信号
	interrupt chan os.Signal

	// complete 通道报告处理任务已经完成
	complete chan error

	// timeout 通道报告处理任务已经超时
	timeout <-chan time.Time

	// tasks 持有一组以索引顺序依次执行的函数
	tasks []func(int)
}

// ErrTimeout 会在任务执行超时时返回
var ErrTimeout = errors.New("received timeout")

// ErrInterrupt 会在接收到操作系统的事件时返回
var ErrInterrupt = errors.New("received interrupt")

// New 返回一个新的准备使用的 Runner
func New(d time.Duration) *Runner {
	return &Runner{
		interrupt: make(chan os.Signal, 1),
		complete:  make(chan error),
		// After 函数返回一个 time.Time 类型的通道
		// 语言运行时会在指定的 duration 时间到期之后，向这个通道发送一个 time.Time 的值
		timeout: time.After(d),
		// 因为 task 字段的零值是 nil，已经满足初始化的要求，所以没有被明确初始化
	}
}

// Add 将一个任务附加到 Runner 上
// 可变参数
// 这个任务是一个接收一个 int 类型的 ID 作为参数的函数
func (r *Runner) Add(tasks ...func(int)) {
	r.tasks = append(r.tasks, tasks...)
}

// Start 执行所有任务，并监视通道事件
func (r *Runner) Start() error {
	// 我们希望接收所有中断信号
	signal.Notify(r.interrupt, os.Interrupt)

	// 用不同的 goroutine 执行不同的任务
	go func() {
		// 方法返回的 error 接口值发送到 complete 通道
		r.complete <- r.run()
	}()

	select {
	// 当任务处理完成时发出的信号
	// 从 complete 通道获取数据
	case err := <-r.complete:
		return err

	// 当任务处理程序运行超时时发出的信号
	// 从 timeout 通道获取数据
	case <-r.timeout:
		return ErrTimeout
	}
}

// run 执行每一个已注册的任务
func (r *Runner) run() error {
	for id, task := range r.tasks {
		// 检测操作系统的中断信号
		if r.gotInterrupt() {
			return ErrInterrupt
		}

		// 执行已注册的任务
		task(id)
	}

	return nil
}

// gotInterrupt 验证是否接收到了中断信号
func (r *Runner) gotInterrupt() bool {
	// 一般来说， select 语句在没有任何要接收的数据时会阻塞，不过有了 default 分支就不会阻塞了
	// default 分支会将接收 interrupt 通道的阻塞调用转变为非阻塞的

	select {
	// 当中断事件被触发时发出的信号
	// 从 interrupt 通道获取数据
	case <-r.interrupt:
		// 停止接收后续的任何信号
		signal.Stop(r.interrupt)
		return true

	// 继续正常运行
	default:
		return false
	}
}

// runner 包管理处理任务的运行和生命周期
package runner

import (
	"errors"
	"os"
	"os/signal"
	"time"
)

// 一个自定义类型
// Runner 在给定的超时时间内执行一组任务，并且在操作系统发送中断信号时结束这些任务
type Runner struct {
	// interrupt 通道
	// 报告从操作系统发送的信号
	interrupt chan os.Signal

	// complete 通道
	// 报告处理任务已经完成
	complete chan error

	// timeout (通道)类型
	// 类型是 <-chan time.Time
	timeout <-chan time.Time

	// 一个函数数组
	// tasks 持有一组以索引顺序依次执行的函数
	tasks []func(int)
}

// ErrTimeout 会在任务执行超时时返回
var ErrTimeout = errors.New("received timeout")

// ErrInterrupt 会在接收到操作系统的事件时返回
var ErrInterrupt = errors.New("received interrupt")

// New 返回一个新的准备使用的 Runner
func New(d time.Duration) *Runner {
	// 返回 Runner 类型的指针
	return &Runner{
		// 创建一个带缓冲的通道
		interrupt: make(chan os.Signal, 1),
		// 创建一个不带缓冲的通道
		complete: make(chan error),
		// After 函数返回一个 time.Time 类型的通道
		// 语言运行时会在指定的 duration 时间到期之后，向这个通道发送一个 time.Time 的值
		timeout: time.After(d),
		// 因为 task 字段的零值是 nil，已经满足初始化的要求，所以没有被明确初始化
	}
}

// Add 将一个任务附加到 Runner 上
// 可变参数
// 这个任务是一个接收一个 int 类型的 ID 作为参数的函数
func (runner *Runner) Add(tasks ...func(int)) {
	runner.tasks = append(runner.tasks, tasks...)
}

// Start 执行所有任务，并监视通道事件
func (runner *Runner) Start() error {
	// 将 os.Interrupt 中断发送到 interrupt 通道
	// 我们希望接收所有中断信号，相当于此处是订阅中断信号
	signal.Notify(runner.interrupt, os.Interrupt)

	// 用不同的 goroutine 执行不同的任务
	go func() {
		// 方法返回的 error 接口值发送到 complete 通道
		runner.complete <- runner.run()
	}()

	// 从 complete 或者 timeout 通道获取值
	// 谁先收到数据，谁就返回
	select {
	// complete 是无缓冲区的通道
	// 当任务处理完成时发出的信号
	// 从 complete 通道获取数据
	case err := <-runner.complete:
		// 如果无错误，就返回 nil
		return err

	// 当任务处理程序运行超时时发出的信号
	// 从 timeout 通道获取数据
	case <-runner.timeout:
		// 返回 timeout 错误
		return ErrTimeout
	}
}

// run 执行每一个已注册的任务
func (runner *Runner) run() error {
	for id, task := range runner.tasks {
		// 每次运行的时候
		// 检测操作系统的中断信号
		if runner.gotInterrupt() {
			// 返回 interrupt 错误
			return ErrInterrupt
		}

		// 执行已注册的任务
		task(id)
	}

	return nil
}

// gotInterrupt 验证是否接收到了中断信号
func (runner *Runner) gotInterrupt() bool {
	// 一般来说， select 语句在没有任何要接收的数据时会阻塞，不过有了 default 分支就不会阻塞了
	// default 分支会将接收 interrupt 通道的阻塞调用转变为非阻塞的

	select {
	// 从 interrupt 通道中接收数据
	// 当中断事件被触发时发出的信号
	case <-runner.interrupt:
		// 停止接收后续的任何信号
		signal.Stop(runner.interrupt)
		return true

	// 继续正常运行
	default:
		return false
	}
}

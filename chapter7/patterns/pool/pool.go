// 包 pool 管理用户定义的一组资源
package pool

import (
	"errors"
	"io"
	"log"
	"sync"
)

// Pool 管理一组可以安全地在多个 goroutine 间共享的资源
// 被管理的资源必须实现 io.Closer 接口
type Pool struct {
	// 这个互斥锁用来保证在多个 goroutine 访问资源池时，池内的值是安全的
	mutex sync.Mutex
	// 有缓冲的通道，用来管理资源的队列
	resources chan io.Closer
	factory   func() (io.Closer, error)
	closed    bool
}

// ErrPoolClosed 表示请求（Acquire）了一个已经关闭的池
var ErrPoolClosed = errors.New("pool has been closed")

// New 创建一个用来管理资源的池
// 这个池需要一个可以分配新资源的函数，并规定池的大小
func New(fn func() (io.Closer, error), size uint) (*Pool, error) {
	if size <= 0 {
		return nil, errors.New("size value too small")
	}

	return &Pool{
		factory: fn,
		// 缓冲通道，大小为 size
		resources: make(chan io.Closer, size),
	}, nil
}

// Acquire 从池中获取一个资源
func (pool *Pool) Acquire() (io.Closer, error) {
	select {
	// 检查是否有空闲的资源
	case resource, ok := <-pool.resources:
		log.Println("Acquire:", "--------------------------- Shared Resource ---------------------------")
		if !ok {
			return nil, ErrPoolClosed
		}
		return resource, nil

	// 因为没有空闲资源可用，所以提供一个新资源
	default:
		log.Println("Acquire:", "New Resource")
		return pool.factory()
	}
}

// Release 将一个使用后的资源放回池里
func (pool *Pool) Release(r io.Closer) {
	// 保证本操作和 Close 操作的安全
	pool.mutex.Lock()
	defer pool.mutex.Unlock()

	// 如果池已经被关闭，销毁这个资源
	if pool.closed {
		r.Close()
		return
	}

	select {
	// 试图将这个资源放入队列
	case pool.resources <- r:
		log.Println("Release:", "In Queue")

	// 如果队列已满，则关闭这个资源
	default:
		log.Println("Release:", "Closing")
		r.Close()
	}
}

// Close 会让资源池停止工作，并关闭所有现有的资源
func (pool *Pool) Close() {
	// 保证本操作与 Release 操作的安全
	pool.mutex.Lock()
	defer pool.mutex.Unlock()

	// 如果 pool 已经被关闭，什么也不做
	if pool.closed {
		return
	}

	// 将池关闭
	pool.closed = true

	// 在清空通道里的资源之前，将通道关闭
	// 如果不这样做，会发生死锁
	close(pool.resources)

	// 关闭资源
	for resource := range pool.resources {
		resource.Close()
	}
}

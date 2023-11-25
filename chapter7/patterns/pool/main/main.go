// 这个示例程序展示如何使用 pool 包来共享一组模拟的数据库连接
package main

import (
	"io"
	"log"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"

	"goinaction.zhangjin.me/chapter7/patterns/pool"
)

const (
	// 要使用的 goroutine 的数量
	maxGoroutines = 100
	// 池中的资源的数量
	pooledResources = 2
)

// init 初始化包
// Go 语言运行时会在其他代码执行之前优先执行这个函数
func init() {
	// 初始化随机数种子
	rand.Seed(time.Now().Unix())
}

// dbConnection 模拟要共享的资源
type dbConnection struct {
	ID int32
}

// Close 实现了 io.Closer 接口，以便 dbConnection 可以被池管理
// Close 用来完成任意资源的释放管理
func (dbConn *dbConnection) Close() error {
	log.Println("Close: Connection", dbConn.ID)
	return nil
}

// idCounter 用来给每个连接分配一个独一无二的 id
var idCounter int32

// createConnection 是一个工厂函数
// 当需要一个新连接时，资源池会调用这个函数
func createConnection() (io.Closer, error) {
	// 使用原子方式保证唯一性
	id := atomic.AddInt32(&idCounter, 1)
	log.Println("Create: New Connection", id)

	// 创建一个 dbConnection 类型，并返回其地址
	return &dbConnection{id}, nil
}

func main() {
	var wg sync.WaitGroup

	wg.Add(maxGoroutines)

	// 创建用来管理连接的池
	// 指定工厂函数和池容量
	pool, err := pool.New(createConnection, pooledResources)
	if err != nil {
		log.Println(err)
	}

	// 使用池里的连接来完成查询
	for query := 0; query < maxGoroutines; query++ {

		// 添加一个等待时间
		// 让 pool 中的 resources 可以被是否
		// 用于模拟 Shared Resource 的效果

		// 注意：不能放在 performQueries 中等待，因为这是一个闭包
		// 全部请求会立即执行
		time.Sleep(time.Duration(rand.Intn(1000)) * time.Microsecond)

		// 每个 goroutine 需要自己复制一份要查询值的副本
		// 不然所有的查询会共享同一个查询变量
		go func(q int) {
			performQueries(q, pool)
			wg.Done()
		}(query)
	}

	// 等待 goroutine 结束
	wg.Wait()

	log.Println("Shutdown Program.")

	// 关闭池
	pool.Close()
}

// performQueries 用来测试连接的资源池
func performQueries(query int, pool *pool.Pool) {
	// 从池里请求一个连接
	conn, err := pool.Acquire()
	if err != nil {
		log.Println(err)
		return
	}

	// 方法返回的时候
	// 将该连接释放回池里
	defer pool.Release(conn)

	// 用等待来模拟查询响应
	time.Sleep(time.Duration(rand.Intn(1000)) * time.Microsecond)

	log.Printf("Query: QID[%d] CID[%d]\n", query, conn.(*dbConnection).ID)
}

package search

import (
	"log"
	"sync"
)

// 这个变量没有定义在任何函数作用域内，所以会被当成包级变量

// 注意：这里的 map 的键是 string 类型，值是 Matcher 类型
// 变量声明以小写字母开头，表示不公开
// 以大写字母开头的，表示公开
// 使用特殊的内置函数 make 初始化了变量 matchers
var matchers = make(map[string]Matcher)

func Run(searchTerm string) {
	// 从配置 json 中获取数据
	feeds, err := RetrieveFeeds()
	if err != nil {
		log.Fatal(err)
	}

	// 根据经验，如果需要声明初始值为 "零值" 的变量，应该使用 var 关键字声明变量
	// 如果提供确切的非零值初始化变量或者使用函数返回值创建变量，应该使用 "简化变量声明运算符"

	// 创建一个无缓冲的通道，接收匹配后的结果
	results := make(chan *Result)

	// 构造一个 waitGroup，以便处理所有的数据源
	var waitGroup sync.WaitGroup

	// 设置需要等待处理的数据源的数量
	waitGroup.Add(len(feeds))

	// 为每个数据源启动一个 goroutine 来查找结果
	for _, feed := range feeds {

		// 获取一个匹配器用于查找
		// type 都是 rss
		matcher, exists := matchers[feed.Type]
		if !exists {
			matcher = matchers["default"]
		}

		// 启动一个 goroutine 执行查找
		go func(matcher Matcher, feed *Feed) {
			// 将 channel 穿进入
			Match(matcher, feed, searchTerm, results)
			waitGroup.Done()
		}(matcher, feed)
	}

	// 启动一个 goroutine 来监控是否所有的工作都做完了
	go func() {
		// 等候所有任务完成
		waitGroup.Wait()

		// 用关闭通道的方式，通知 Display 函数可以退出程序了
		close(results)
	}()

	// 启动函数，显示返回的结果
	// 并且在最后一个结果显示完后返回
	Display(results)
}

// Register 调用时，会注册一个匹配器，提供给后面的程序使用
func Register(feedType string, matcher Matcher) {
	if _, exists := matchers[feedType]; exists {
		log.Fatalln(feedType, "Matcher already registered")
	}

	log.Println("Register", feedType, "matcher")
	matchers[feedType] = matcher
}

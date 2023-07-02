package search

import (
	"log"
	"sync"
)

// 这个变量没有定义在任何函数作用域内，所以会被当成包级变量
// A map of registered matchers for searching.

// 注意：这里的 map 的键是 string 类型，值是 Matcher 类型
// 变量声明以小写字母开头，表示不公开
// 以大写字母开头的，表示公开
// 使用特殊的内置函数 make 初始化了变量 matchers
var matchers = make(map[string]Matcher)

// Run performs the search logic.
func Run(searchTerm string) {
	// Retrieve the list of feeds to search through.
	feeds, err := RetrieveFeeds()
	if err != nil {
		log.Fatal(err)
	}

	// 根据经验，如果需要声明初始值为零值的变量，应该使用 var 关键字声明变量
	// 如果提供确切的非零值初始化变量或者使用函数返回值创建变量，应该使用简化变量声明运算符

	// Create an unbuffered channel to receive match results to display.
	results := make(chan *Result)

	// Setup a wait group so we can process all the feeds.
	var waitGroup sync.WaitGroup

	// Set the number of goroutines we need to wait for while
	// they process the individual feeds.
	waitGroup.Add(len(feeds))

	// Launch a goroutine for each feed to find the results.
	// 为每个数据源启动一个 goroutine 来查找结果
	for _, feed := range feeds {

		// Retrieve a matcher for the search.
		// 获取一个匹配器用于查找
		matcher, exists := matchers[feed.Type]
		if !exists {
			matcher = matchers["default"]
		}

		// Launch the goroutine to perform the search.
		// 启动一个 goroutine 执行查找
		go func(matcher Matcher, feed *Feed) {
			Match(matcher, feed, searchTerm, results)
			waitGroup.Done()
		}(matcher, feed)
	}

	// Launch a goroutine to monitor when all the work is done.
	// 启动一个 goroutine 来监控是否所有的工作都做完了
	go func() {
		// Wait for everything to be processed.
		// 等候所有任务完成
		waitGroup.Wait()

		// Close the channel to signal to the Display
		// function that we can exit the program.
		// 用关闭通道的方式，通知 Display 函数
		// 可以退出程序了
		close(results)
	}()

	// Start displaying results as they are available and
	// return after the final result is displayed.
	// 启动函数，显示返回的结果
	// 并且在最后一个结果显示完后返回
	Display(results)
}

// Register is called to register a matcher for use by the program.
// Register 调用时，会注册一个匹配器，提供给后面的程序使用
func Register(feedType string, matcher Matcher) {
	if _, exists := matchers[feedType]; exists {
		log.Fatalln(feedType, "Matcher already registered")
	}

	log.Println("Register", feedType, "matcher")
	matchers[feedType] = matcher
}

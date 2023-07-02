package search

import (
	"log"
)

// Result contains the result of a search.
// Result 保存搜索的结果
type Result struct {
	Field   string
	Content string
}

// Matcher defines the behavior required by types that want to implement a new search type.
// Matcher 定义了要实现的新搜索类型的行为
type Matcher interface {
	Search(feed *Feed, searchTerm string) ([]*Result, error)
}

// Match is launched as a goroutine for each individual feed to run searches concurrently.
// Match 函数，为每个数据源单独启动 goroutine 来执行这个函数并发地执行搜索
func Match(matcher Matcher, feed *Feed, searchTerm string, results chan<- *Result) {
	// Perform the search against the specified matcher.
	// 对特定的匹配器执行搜索
	searchResults, err := matcher.Search(feed, searchTerm)
	if err != nil {
		log.Println(err)
		return
	}

	// Write the results to the channel.
	// 将结果写入通道
	for _, result := range searchResults {
		results <- result
	}
}

// Display writes results to the console window as they are received by the individual goroutines.
// Display 从每个单独的 goroutine 接收到结果后在终端窗口输出
func Display(results chan *Result) {
	// The channel blocks until a result is written to the channel.
	// Once the channel is closed the for loop terminates.
	// 通道会一直被阻塞，直到有结果写入
	// 一旦通道被关闭， for 循环就会终止
	for result := range results {
		log.Printf("%s:\n%s\n\n", result.Field, result.Content)
	}
}

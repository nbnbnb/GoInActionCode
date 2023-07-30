package search

import "log"

// 自定义类型
// 封装搜索结果
type Result struct {
	Engine      string
	Title       string
	Description string
	Link        string
}

// 自定义接口
type Searcher interface {
	Search(searchTerm string, searchResults chan<- []Result)
}

// 自定义类型
// 封装搜索会话
type searchSession struct {
	searchers  map[string]Searcher
	first      bool
	resultChan chan []Result
}

// 查询 Google
func Google(s *searchSession) {
	log.Println("search : Submit : Info : Adding Google")
	// 添加到 map 中
	s.searchers["google"] = google{}
}

// 查询 Bing
func Bing(s *searchSession) {
	log.Println("search : Submit : Info : Adding Bing")
	// 添加到 map 中
	s.searchers["bing"] = bing{}
}

// 查询 Yahoo
func Yahoo(s *searchSession) {
	log.Println("search : Submit : Info : Adding Yahoo")
	// 添加到 map 中
	s.searchers["yahoo"] = yahoo{}
}

// 设置是否只返回第一个结果
func OnlyFirst(s *searchSession) {
	// 设置 flag
	s.first = true
}

// 提交查询
func Submit(query string, options ...func(*searchSession)) []Result {
	// 本地变量
	// 零值初始化
	var session searchSession

	// 创建一个 map[string]
	session.searchers = make(map[string]Searcher)

	// 创建一个通道
	// 通道值是 []Result
	session.resultChan = make(chan []Result)

	// options 是是个配置函数列表
	for _, opt := range options {
		// 执行每一个 option 函数
		opt(&session)
	}

	// 执行并行查询
	for _, s := range session.searchers {
		// 通过 goroutine 执行查询
		go s.Search(query, session.resultChan)
	}

	var results []Result

	// 等待结果返回
	for search := 0; search < len(session.searchers); search++ {
		// 如果只需要第一个结果
		if session.first && search > 0 {
			// 丢弃结果
			go func() {
				r := <-session.resultChan
				log.Printf("search : Submit : Info : Results Discarded : Results[%d]\n", len(r))
			}()
			continue
		}

		log.Println("search : Submit : Info : Waiting For Results...")
		// 从通道中等待结果
		result := <-session.resultChan

		log.Printf("search : Submit : Info : Results Used : Results[%d]\n", len(result))
		// 获取到结果
		// 将结果添加到 results 中
		results = append(results, result...)
	}

	log.Printf("search : Submit : Completed : Found [%d] Results\n", len(results))

	return results
}

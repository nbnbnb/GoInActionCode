package search

// defaultMatcher 实现了默认匹配器
type defaultMatcher struct{}

// init 函数将默认匹配器注册到程序里
func init() {
	var matcher defaultMatcher
	Register("default", matcher)
}

// Search 实现了默认匹配器的行为

// 如果声明函数的时候带有接收者，则意味着声明了一个方法
// 这个方法会和指定的接收者的类型绑在一起
// 这里，Search 方法与 defaultMatcher 类型的值绑在一起
func (m defaultMatcher) Search(feed *Feed, searchTerm string) ([]*Result, error) {
	return nil, nil
}

// 这个示例程序展示如何访问另一个包的未公开的标识符的值
package main

import (
	// 注意包引用语法
	"GoInActionCode/chapter5/listing68/counters"
	"fmt"
)

func main() {
	// 使用 counters 包公开的 New 函数来创建一个未公开的类型的变量
	// New 首字母是大写的，表示是公开的
	counter := counters.New(10)

	fmt.Printf("Counter: %d\n", counter)
}

// 这个示例程序展示无法从另一个包里访问未公开的标识符
package main

import "fmt"

func main() {
	// 当要写的代码属于某个包时，好的实践是使用与代码所在文件夹一样的名字作为包名
	// counters 包就是一个很好的例子

	// 创建一个未公开的类型的变量，并赋初值为 10
	// 错误：不能引用未公开的名字 counters.alertCounter
	// counter := counters.alertCounter(10)
	// fmt.Printf("Counter: %d\n", counter)
	fmt.Println("nothing")
}

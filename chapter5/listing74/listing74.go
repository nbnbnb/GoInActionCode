// 这个示例程序展示公开的结构类型中如何访问未公开的内嵌类型的例子
package main

import (
	"fmt"

	"goinaction.zhangjin.me/chapter5/listing74/entities"
)

func main() {
	a := entities.Admin{
		Rights: 10,
		// 由于内部类型 user 是未公开的，此处代码无法直接通过结构字面量的方式初始化该内部类型
	}

	// 不过，即便内部类型是未公开的，内部类型里声明的字段依旧是公开的（大写字母开头）

	// 既然内部类型的标识符提升到了外部类型，这些公开的字段也可以通过外部类型的字段的值来访问

	// 内嵌类型未公开字段的值，提升到外部类型访问
	// 设置未公开的内部类型的公开字段的值
	a.Name = "Bill"
	a.Email = "bill@email.com"

	fmt.Printf("User: %v\n", a)
}

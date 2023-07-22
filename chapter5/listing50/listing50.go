// 这个示例程序展示如何将一个类型嵌入另一个类型，以及内部类型和外部类型之间的关系
package main

import (
	"fmt"
)

type user struct {
	name  string
	email string
}

func (u *user) notify() {
	fmt.Printf("Sending user email to %s<%s>\n", u.name, u.email)
}

type admin struct {
	// 要嵌入一个类型，只需要声明这个类型的名字就可以了，要注意这个和字段声明的差异
	// 嵌入类型
	user
	level string
}

func main() {
	ad := admin{
		user: user{
			name:  "john smith",
			email: "john@yahoo.com",
		},
		level: "super",
	}

	// 我们可以直接访问内部类型的方法
	ad.user.notify()

	// 内部类型提升
	// 内部类型的方法也被提升到外部类型
	ad.notify()
}

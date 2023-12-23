// 这个示例程序展示如何将嵌入类型应用于接口
package main

import (
	"fmt"
)

// 这是一个接口
type notifier interface {
	notify()
}

// 这是一个用户定义类型
type user struct {
	name  string
	email string
}

// 给用户定义类型，添加接口功能呢
func (user *user) notify() {
	fmt.Printf("Sending user email to %s<%s>\n", user.name, user.email)
}

// 这是一个用户定义类型
type admin struct {
	// 注意：这里就是一个嵌入类型，声明方式不需要类型名
	user
	level string
}

func main() {
	// 内部类型提升
	// admin 也实现了 notifier 能力
	admin := admin{
		// 给嵌入类型赋值
		user: user{
			name:  "john smith",
			email: "john@yahoo.com",
		},
		level: "super",
	}

	// 由于内部类型的提升，内部类型实现的接口会自动提升到外部类型
	// 这意味着由于内部类型的实现，外部类型也同样实现了这个接口

	// 给 admin 用户发送一个通知用于实现接口的内部类型的方法，被提升到外部类型

	// admin 也实现了 notifier 能力
	admin.notify()

	// 内部类型 user 实现了 notifier 接口
	// 提升到外部类型 admin
	// 所以此处可以传递一个 admin 类型的值给 sendNotification 函数
	sendNotification(&admin)
}

func sendNotification(notifier notifier) {
	notifier.notify()
}

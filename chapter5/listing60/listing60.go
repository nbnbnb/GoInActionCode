// 这个示例程序展示当内部类型和外部类型要实现同一个接口时的做法
package main

import (
	"fmt"
)

// 一个接口
type notifier interface {
	notify()
}

// 一个自定义类型
type user struct {
	name  string
	email string
}

// user 类型实现了 notifier 接口
func (u *user) notify() {
	fmt.Printf("Sending user email to %s<%s>\n", u.name, u.email)
}

// 一个自定义类型
type admin struct {
	// 内嵌类型
	user
	level string
}

// admin 类型实现了 notifier 接口
func (admin *admin) notify() {
	fmt.Printf("Sending admin email to %s<%s>\n", admin.name, admin.email)
}

func main() {
	admin := admin{
		user: user{
			name:  "john smith",
			email: "john@yahoo.com",
		},
		level: "super",
	}

	// 这表明，如果外部类型实现了 notify 方法，内部类型的实现就不会被提升

	// 给 admin 用户发送一个通知接口的嵌入的内部类型实现并 "没有" 提升到外部类型
	// Sending admin email to john smith<john@yahoo.com>
	sendNotification(&admin)

	// 我们可以直接访问内部类型的方法
	// Sending user email to john smith<john@yahoo.com>
	admin.user.notify()

	// 验证内部类型的方法 "没有" 被提升
	// Sending admin email to john smith<john@yahoo.com>
	admin.notify()
}

func sendNotification(notifier notifier) {
	notifier.notify()
}

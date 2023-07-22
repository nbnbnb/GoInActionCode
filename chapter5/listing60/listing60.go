// 这个示例程序展示当内部类型和外部类型要实现同一个接口时的做法
package main

import (
	"fmt"
)

type notifier interface {
	notify()
}

type user struct {
	name  string
	email string
}

// 通过 user 类型值的指针调用的方法
func (u *user) notify() {
	fmt.Printf("Sending user email to %s<%s>\n",
		u.name,
		u.email)
}

type admin struct {
	user
	level string
}

// 通过 admin 类型值的指针调用的方法
func (a *admin) notify() {
	fmt.Printf("Sending admin email to %s<%s>\n",
		a.name,
		a.email)
}

func main() {
	ad := admin{
		user: user{
			name:  "john smith",
			email: "john@yahoo.com",
		},
		level: "super",
	}

	// 这表明，如果外部类型实现了 notify 方法，内部类型的实现就不会被提升

	// 给 admin 用户发送一个通知接口的嵌入的内部类型实现并 "没有" 提升到外部类型
	sendNotification(&ad)

	// 我们可以直接访问内部类型的方法
	ad.user.notify()

	// 内部类型的方法 "没有" 被提升
	ad.notify()
}

func sendNotification(n notifier) {
	n.notify()
}

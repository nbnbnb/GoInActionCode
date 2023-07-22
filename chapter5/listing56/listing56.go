// 这个示例程序展示如何将嵌入类型应用于接口
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

func (u *user) notify() {
	fmt.Printf("Sending user email to %s<%s>\n",
		u.name,
		u.email)
}

type admin struct {
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

	// 由于内部类型的提升，内部类型实现的接口会自动提升到外部类型
	// 这意味着由于内部类型的实现，外部类型也同样实现了这个接口

	// 给 admin 用户发送一个通知用于实现接口的内部类型的方法，被提升到外部类型
	sendNotification(&ad)
}

func sendNotification(n notifier) {
	n.notify()
}

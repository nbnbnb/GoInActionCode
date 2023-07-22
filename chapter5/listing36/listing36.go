// 这个示例程序展示 Go 语言里如何使用接口
package main

import (
	"fmt"
)

// notifier 是一个定义了通知类行为的接口
type notifier interface {
	notify()
}

// user 在程序里定义一个用户类型
type user struct {
	name  string
	email string
}

// 如果使用 “指针接收者” 来实现一个接口，那么只有指向那个类型的指针才能够实现对应的接口
// 如果使用 “值接收者“ 来实现一个接口，那么那个类型的值和指针都能够实现对应的接口

// 此次使用的是 “指针接收者” 来实现接口
func (u *user) notify() {
	fmt.Printf("Sending user email to %s<%s>\n",
		u.name,
		u.email)
}

func main() {
	// 创建一个 user 类型的值，并发送通知
	u := user{"Bill", "bill@email.com"}

	// 传递 user 类型的 "指针" 给 sendNotification 没有问题
	sendNotification(&u)

	// 传递 user 类型的 "值" 给 sendNotification 会引发一个编译错误
	// sendNotification(u)
}

// sendNotification 接受一个实现了 notifier 接口的值，并发送通知
func sendNotification(n notifier) {
	n.notify()
}

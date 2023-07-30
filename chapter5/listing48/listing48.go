// 这个示例程序使用接口展示多态行为
package main

import (
	"fmt"
)

// notifier 是一个定义了通知类行为的接口
type notifier interface {
	notify()
}

type user struct {
	name  string
	email string
}

// notify 使用 "指针接收者" 实现了 notifier 接口
func (u *user) notify() {
	fmt.Printf("Sending user email to %s<%s>\n",
		u.name,
		u.email)
}

type admin struct {
	name  string
	email string
}

// notify 使用 "指针接收者" 实现了 notifier 接口
func (a *admin) notify() {
	fmt.Printf("Sending admin email to %s<%s>\n", a.name, a.email)
}

func main() {
	bill := user{"Bill", "bill@email.com"}
	sendNotification(&bill)

	lisa := admin{"Lisa", "lisa@email.com"}
	sendNotification(&lisa)
}

// sendNotification 接受一个实现了 notifier 接口的值并发送通知
func sendNotification(n notifier) {
	n.notify()
}

// 这个示例程序展示如何声明并使用方法
package main

import (
	"fmt"
)

// 定义一个 user 用户类型
type user struct {
	name  string
	email string
}

// notify 方法 “值接收者” 是 user
func (u user) notify() {
	fmt.Printf("Sending User Email To %s<%s>\n", u.name, u.email)
}

// changeEmail 方法的 “指针接收者” 是 *user
func (u *user) changeEmail(email string) {
	u.email = email
}

// main 是应用程序的入口
func main() {
	// 值类型
	value_type := user{"Bill", "bill@email.com"}

	// 指针类型
	pointer_type := &user{"Lisa", "lisa@email.com"}

	// 值类型 -> 值接收者
	// notify 方法，收到 value_type 值的一个副本
	value_type.notify()

	// 指针类型 -> 值接收者
	// notify 方法，收到的是 pointer_type（指针）值的一个副本
	pointer_type.notify()

	// 值类型 -> 指针接收者
	// Go 编译器实际值的是 (&value_type).changeEmail("bill@newdomain")
	value_type.changeEmail("bill@newdomain.com")

	// 指针类型 -> 指针接收者
	pointer_type.changeEmail("lisa@newdomain.com")
}

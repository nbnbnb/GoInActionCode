// 这个示例程序展示不是总能获取地址的值
package main

import "fmt"

// duration 是一个基于 int 类型的类型
type duration int

// 使用更可读的方式格式化 duration 值
func (d *duration) pretty() string {
	return fmt.Sprintf("Duration: %d", *d)
}

func main() {
	// 不能获取到 42 这个值的地址
	// 因为不是总能获取一个值的地址，所以值的方法集只包括了使用值接收者声明的方法
	// duration(42).pretty()

	// ./listing46.go:17: 不能通过指针调用 duration(42) 的方法
	// ./listing46.go:17: 不能获取 duration(42) 的地址

	// Values               Methods Receivers
	// -----------------------------------------------
	// T					(t T)
	// *T					(t T) and (t *T)

	// Methods Receivers    Values
	// -----------------------------------------------
	// (t T)				T and *T
	// (t *T)				*T

	println("nothing")
}

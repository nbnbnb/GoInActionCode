// 这个示例程序展示不是总能获取地址的值
package main

import "fmt"

// duration 是一个基于 int 类型的类型
type duration int

// ”指针“ 接收者
func (d *duration) prettyPoint() string {
	return fmt.Sprintf("Duration: %d", *d)
}

// ”值“ 接收者
func (d duration) prettyValue() string {
	return fmt.Sprintf("Duration: %d", d)
}

func main() {
	// 不能获取到 42 这个值的地址
	// 因为不是总能获取一个值的地址，所以值的方法集只包括了使用值接收者声明的方法
	// duration(42).prettyPoint()

	// Values               Methods Receivers
	// -----------------------------------------------
	// T					(t T)
	// *T					(t T) and (t *T)

	// Methods Receivers    Values
	// -----------------------------------------------
	// (t T)				T and *T
	// (t *T)				*T

	println(duration(42).prettyValue())
}

// entities 包包含系统中与人有关的类型
package entities

// User 在程序里定义一个用户类型
type User struct {
	// 大写字母开头，公开的
	Name string
	// 小写字母开头，未公开的
	email string
}

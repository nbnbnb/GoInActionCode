package entities

// 小写字母开头，未公开的类型
type user struct {
	// 但是内部的字段是公开的
	// 可以提升到外部类型中访问
	Name  string
	Email string
}

type Admin struct {
	// 嵌入的类型未公开
	user
	Rights int
}

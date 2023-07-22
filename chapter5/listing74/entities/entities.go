package entities

type user struct {
	Name  string
	Email string
}

type Admin struct {
	// 嵌入的类型未公开
	user
	Rights int
}

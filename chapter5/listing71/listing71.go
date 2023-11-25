// 这个示例程序展示公开的结构类型中未公开的字段无法直接访问
package main

import (
	"fmt"

	"goinaction.zhangjin.me/chapter5/listing71/entities"
)

func main() {
	// 创建 entities 包中的 User 类型的值
	u := entities.User{
		Name: "Bill",
		// 由于 email 是未公开的字段，所以无法直接通过结构字面量的方式初始化该字段
		// unknown entities.User field 'email' in struct literal
		// email: "bill@email.com",
	}

	fmt.Printf("User: %v\n", u)
}

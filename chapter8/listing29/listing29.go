// 这个示例程序展示如何解码 JSON 字符串
package main

import (
	"encoding/json"
	"fmt"
	"log"
)

// JSON 包含要反序列化的样例字符串
var JSON = `{
	"name": "Gopher",
	"title": "programmer",
	"contact": {
		"home": "415.333.3333",
		"cell": "415.555.5555"
	}
}`

func main() {
	// 将 JSON 字符串反序列化到 map 变量

	// 变量 c 声明为一个 map 类型，其键是 string 类型，其值是 interface{} 类型
	var c map[string]interface{}
	err := json.Unmarshal([]byte(JSON), &c)
	if err != nil {
		log.Println("ERROR:", err)
		return
	}

	fmt.Println("Name:", c["name"])
	fmt.Println("Title:", c["title"])
	fmt.Println("Contact")
	// 需要强制转换为 map[string]interface{} 类型
	fmt.Println("H:", c["contact"].(map[string]interface{})["home"])
	fmt.Println("C:", c["contact"].(map[string]interface{})["cell"])
}

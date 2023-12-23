// 这个示例程序展示如何反序列化 JSON 字符串
package main

import (
	"encoding/json"
	"fmt"
	"log"
)

// Info 包含要反序列化的样例字符串
var Info = `{
	"name": "Gopher",
	"title": "programmer",
	"contact": {
		"home": "415.333.3333",
		"cell": "415.555.5555"
	}
}`

func main() {
	// 将 JSON 字符串反序列化到 map 变量

	// 变量 info 声明为一个 map 类型，其键是 string 类型，其值是 interface{} 类型
	var info map[string]interface{}
	err := json.Unmarshal([]byte(Info), &info)
	if err != nil {
		log.Println("ERROR:", err)
		return
	}

	fmt.Println("Name:", info["name"])
	fmt.Println("Title:", info["title"])
	fmt.Println("Contact")
	// 需要强制转换为 map[string]interface{} 类型
	fmt.Println("Home:", info["contact"].(map[string]interface{})["home"])
	fmt.Println("Cell:", info["contact"].(map[string]interface{})["cell"])
}

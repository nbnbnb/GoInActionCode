// 这个示例程序展示如何反序列化 JSON 字符串
package main

import (
	"encoding/json"
	"fmt"
	"log"
)

// Info 结构代表我们的 JSON 字符串
type Info struct {
	// 必须得是大写字母开头
	// 表示公共字段
	Name    string `json:"name"`
	Title   string `json:"title"`
	Contact struct {
		Home string `json:"home"`
		Cell string `json:"cell"`
	} `json:"contact"`
}

// JSON 包含用于反序列化的演示字符串
var JSON = `{
	"name": "Gopher",
	"title": "programmer",
	"contact": {
		"home": "415.333.3333",
		"cell": "415.555.5555"
	}
}`

func main() {
	// 将 JSON 字符串反序列化到变量
	var info Info

	// 将字符串反序列化到变量
	err := json.Unmarshal([]byte(JSON), &info)
	if err != nil {
		log.Println("ERROR:", err)
		return
	}

	fmt.Println(info)
}

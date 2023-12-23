// 这个示例程序展示如何序列化 JSON 字符串
package main

import (
	"encoding/json"
	"fmt"
	"log"
)

func main() {
	// 创建一个保存键值对的映射
	info := make(map[string]interface{})
	info["name"] = "Gopher"
	info["title"] = "programmer"
	info["contact"] = map[string]interface{}{
		"home": "415.333.3333",
		"cell": "415.555.5555",
	}

	// 将这个映射序列化到 JSON 字符串
	// MarshalIndent 很像 Marshal，只是用缩进对输出进行格式化
	data, err := json.MarshalIndent(info, "-", "    ")
	if err != nil {
		log.Println("ERROR:", err)
		return
	}

	fmt.Println(string(data))
}

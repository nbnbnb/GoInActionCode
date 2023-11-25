// 示例并行请求
// 等待全部返回或者第一个返回
package main

import (
	"log"

	"goinaction.zhangjin.me/chapter7/patterns/search"
)

func main() {
	// Submit the search and display the results.
	results := search.Submit(
		"golang",
		// option 配置
		// 只返回第一个结果
		search.OnlyFirst,
		search.Google,
		search.Bing,
		search.Yahoo,
	)

	for _, result := range results {
		log.Printf("main : Results : Info : %+v\n", result)
	}

	log.Println("--------------------------------------------------")

	// 等待所有的结果返回
	results = search.Submit(
		"golang",
		search.Google,
		search.Bing,
		search.Yahoo,
	)

	for _, result := range results {
		log.Printf("main : Results : Info : %+v\n", result)
	}
}

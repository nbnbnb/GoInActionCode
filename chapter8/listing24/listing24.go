// 这个示例程序展示如何使用 json 包和 NewDecoder 函数来解码 JSON 响应
package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type (
	// result 映射到从搜索拿到的结果文档

	// 每个字段最后使用单引号声明了一个字符串
	// 这些字符串被称作标签（ tag），是提供每个字段的元信息的一种机制，将 JSON 文档和结构类型里的字段一一映射起来
	// 如果不存在标签，编码和解码过程会试图以大小写无关的方式，直接使用字段的名字进行匹配
	// 如果无法匹配，对应的结构类型里的字段就包含其零值
	result struct {
		GsearchResultClass string `json:"GsearchResultClass"`
		UnescapedURL       string `json:"unescapedUrl"`
		URL                string `json:"url"`
		VisibleURL         string `json:"visibleUrl"`
		CacheURL           string `json:"cacheUrl"`
		Title              string `json:"title"`
		TitleNoFormatting  string `json:"titleNoFormatting"`
		Content            string `json:"content"`
	}

	// response 包含顶级的文档
	response struct {
		// 必须得是大写字母开头
		// 表示公共字段
		ResponseData struct {
			// 必须得是大写字母开头
			// 表示公共字段
			Results []result `json:"results"`
		} `json:"responseData"`
	}
)

func main() {
	uri := "https://zhangjin.tk/dl/go-search.json?v=1.0&rsz=8&q=golang"

	resp, err := http.Get(uri)
	if err != nil {
		log.Println("ERROR:", err)
		return
	}

	defer resp.Body.Close()

	// 将 JSON 响应解码到结构类型
	var gr response
	err = json.NewDecoder(resp.Body).Decode(&gr)
	if err != nil {
		log.Println("ERROR:", err)
		return
	}

	fmt.Println(gr)

	// 执行格式化输出
	pretty, err := json.MarshalIndent(gr, "-", "    ")
	if err != nil {
		log.Println("ERROR:", err)
		return
	}

	fmt.Println(string(pretty))
}

// 这个示例程序展示如何使用 io.Reader 和 io.Writer 接口写一个简单版本的 curl 程序
package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

// init 在 main 函数之前调用
func init() {
	if len(os.Args) != 2 {
		os.Args = append(os.Args, "https://www.baidu.com")
	}
}

// main 是应用程序的入口
func main() {
	// 从 Web 服务器获取数据
	res, err := http.Get(os.Args[1])
	if err != nil {
		fmt.Println(err)
		return
	}

	// 从 Body 复制到 Stdout
	io.Copy(os.Stdout, res.Body)

	if err := res.Body.Close(); err != nil {
		fmt.Println(err)
	}
}

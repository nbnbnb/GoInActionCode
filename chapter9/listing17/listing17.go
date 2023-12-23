// 这个示例程序实现了简单的网络服务
package main

// 以 _test 结尾得文件，不执行 init 函数（在 test 运行时才执行）
// init 执行顺序是按照文件名排序的

import (
	"log"
	"net/http"

	"goinaction.zhangjin.me/chapter9/listing17/handlers"
)

func main() {
	handlers.Routes()

	log.Println("listener : Started : Listening on :4000")

	// handlers 内部注册了处理器
	// 所以此处第二个参数传递 nil
	http.ListenAndServe(":4000", nil)
}

// 这个示例程序实现了简单的网络服务
package main

import (
	"GoInActionCode/chapter9/listing17/handlers"
	"log"
	"net/http"
)

func main() {
	handlers.Routes()

	log.Println("listener : Started : Listening on :4000")
	http.ListenAndServe(":4000", nil)
}

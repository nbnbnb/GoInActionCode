// 使用标准 log 库
package main

import (
	"log"
)

func init() {
	// 设置前缀
	log.SetPrefix("TRACE: ")
	// 设置前缀头信息
	log.SetFlags(log.Ldate | log.Lmicroseconds | log.Llongfile)
}

func main() {
	// Println 写到标准日志记录器
	log.Println("message")

	// Fatalln 在调用 Println() 之后会接着调用 os.Exit(1)
	// 也就是一个异常终止

	// 输出类似如下信息：
	// Process 21172 has exited with status 1
	// log.Fatalln("fatal message")

	// Panicln 在调用 Println() 之后会接着调用 panic()
	// Panic 系列函数用来写日志消息，然后触发一个 panic
	// 除非程序执行 recover 函数，否则会导致程序打印调用栈后终止
	// log.Panicln("panic message")
}

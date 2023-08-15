package main

import (
	"log"
	"os"

	"goinaction.zhangjin.me/search"

	// 导入前面的 _ 表示
	// 让 Go 语言对包做初始化操作，但是并不使用包里的标识符
	// 只调用包里面的 init 函数
	_ "goinaction.zhangjin.me/matchers"
)

// init is called prior to main.
func init() {
	// Change the device for logging to stdout.
	// 从默认的标准错误（ stderr），设置为标准输出（ stdout）设备
	log.SetOutput(os.Stdout)
}

// main is the entry point for the program.
func main() {
	// Perform the search for the specified term.
	search.Run("president")
}

// 这个示例程序展示来自不同标准库的不同函数是如何使用 io.Writer 接口的
package main

import (
	"bytes"
	"fmt"
	"os"
)

func main() {
	// 创建一个 Buffer 值，并将一个字符串写入 Buffer
	// 使用实现 io.Writer 的 Write 方法
	var buf bytes.Buffer
	buf.Write([]byte("Hello "))

	// 使用 Fprintf 来将一个字符串拼接到 Buffer 里
	// 将 bytes.Buffer 的地址作为 io.Writer 类型值传入
	fmt.Fprintf(&buf, "World! %s", "KKKing!")

	// 将 Buffer 的内容输出到标准输出设备
	// 将 os.Stdout 值的地址作为 io.Writer 类型值传入
	buf.WriteTo(os.Stdout)
}

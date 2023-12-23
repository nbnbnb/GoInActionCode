// 示例在 io.Copy 函数中使用 bytes.Buffer
package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
)

func main() {
	// 创建一个 Buffer 对象
	// 用零值初始化（有内存地址）
	var buf bytes.Buffer

	// 给 Buffer 中写入内容
	buf.Write([]byte("Hello"))

	// 使用 Fprintf 将字符串拼接到 Buffer 中
	fmt.Fprintf(&buf, " World! %s", "KKKing")

	// 将 Buffer 输出到标准输出设备
	io.Copy(os.Stdout, &buf)
}

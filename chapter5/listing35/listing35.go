// 示例在 io.Copy 函数中使用 bytes.Buffer
package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
)

func main() {
	var b bytes.Buffer

	// 给 Buffer 中写入内容
	b.Write([]byte("Hello"))

	// 使用 Fprintf 将字符串拼接到 Buffer 中
	fmt.Fprintf(&b, " World!")

	// 将 Buffer 输出到标准输出设备
	io.Copy(os.Stdout, &b)
}

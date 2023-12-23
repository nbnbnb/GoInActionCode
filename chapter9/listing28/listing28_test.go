// 用来检测要将整数值转为字符串，使用哪个函数会更好的基准测试示例
// 先使用 fmt.Sprintf 函数
// 然后使用 strconv.FormatInt 函数
// 最后使用 strconv.Itoa

// * 和单元测试文件一样，基准测试的文件名也必须以 _test.go 结尾

package listing05_test

import (
	"fmt"
	"strconv"
	"testing"
)

// * 基准测试函数必须以 Benchmark 开头，接受一个指向 testing.B 类型的指针作为唯一参数

// BenchmarkSprintf 对 fmt.Sprintf 函数进行基准测试
func BenchmarkSprintf(b *testing.B) {
	number := 10

	// 初始化变量之后，使用 b.ResetTimer() 函数来重置计时器
	b.ResetTimer()

	// 使用 b.N 值
	for i := 0; i < b.N; i++ {
		fmt.Sprintf("%d", number)
	}
}

// BenchmarkSprintf 对 strconv.FormatInt 函数进行基准测试
func BenchmarkFormat(b *testing.B) {
	number := int64(10)

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		strconv.FormatInt(number, 10)
	}
}

// BenchmarkSprintf 对 strconv.Itoa 函数进行基准测试
func BenchmarkItoa(b *testing.B) {
	number := 10

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		strconv.Itoa(number)
	}
}

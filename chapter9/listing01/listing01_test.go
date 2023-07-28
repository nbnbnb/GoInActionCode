// Go 语言的测试工具只会认为以 _test.go 结尾的文件是测试文件
// 这个示例程序展示如何写基础单元测试
package listing01

import (
	"net/http"
	"testing"
)

const (
	checkMark = "\u2713"
	ballotX   = "\u2717"
)

// 一个测试函数必须是公开的函数，并且以 Test 单词开头
// 函数的签名必须接收一个指向 testing.T 类型的指针，并且不返回任何值

// TestDownload 确认 http 包的 Get 函数可以下载内容
func TestDownload(t *testing.T) {
	url := "https://anotherdayu.com/?feed=rss"
	statusCode := 200

	t.Log("Given the need to test downloading content.")
	{
		t.Logf("\tWhen checking \"%s\" for status code \"%d\"", url, statusCode)
		{
			resp, err := http.Get(url)
			if err != nil {
				t.Fatal("\t\tShould be able to make the Get call.", ballotX, err)
			}
			t.Log("\t\tShould be able to make the Get call.", checkMark)

			defer resp.Body.Close()

			if resp.StatusCode == statusCode {
				t.Logf("\t\tShould receive a \"%d\" status. %v", statusCode, checkMark)
			} else {
				t.Errorf("\t\tShould receive a \"%d\" status. %v %v", statusCode, ballotX, resp.StatusCode)
			}
		}
	}
}

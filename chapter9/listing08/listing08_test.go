// 这个示例程序展示如何写一个基本的 ”表组“ 测试
package listing08

import (
	"net/http"
	"testing"
)

const (
	checkMark = "\u2713"
	ballotX   = "\u2717"
)

// TestDownload 确认 http 包的 Get 函数可以下载内容并正确处理不同的状态
func TestDownload(t *testing.T) {
	urls := []struct {
		url        string
		statusCode int
	}{
		{
			"https://anotherdayu.com/?feed=rss",
			http.StatusOK,
		},
		{
			"http://rss.cnn.com/rss/cnn_topstbadurl.rss",
			http.StatusNotFound,
		},
	}

	t.Log("Given the need to test downloading different content.")
	{
		// 迭代 ”表组“
		for _, u := range urls {
			t.Logf("\tWhen checking \"%s\" for status code \"%d\"", u.url, u.statusCode)
			{
				resp, err := http.Get(u.url)
				if err != nil {
					t.Fatal("\t\tShould be able to Get the url.", ballotX, err)
				}
				t.Log("\t\tShould be able to Get the url.", checkMark)

				defer resp.Body.Close()

				if resp.StatusCode == u.statusCode {
					t.Logf("\t\tShould have a \"%d\" status. %v", u.statusCode, checkMark)
				} else {
					t.Errorf("\t\tShould have a \"%d\" status. %v %v", u.statusCode, ballotX, resp.StatusCode)
				}
			}
		}
	}
}

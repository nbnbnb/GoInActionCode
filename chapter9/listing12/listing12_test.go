// 这个示例程序展示如何内部模仿 HTTP GET 调用
package listing12

import (
	"encoding/xml"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

const (
	// ✓
	checkMark = "\u2713"
	// ✗
	ballotX = "\u2717"
)

// feed 模仿了我们期望接收的 XML 文档
var feed = `<?xml version="1.0" encoding="UTF-8"?>
<rss>
<channel>
    <title>Going Go Programming</title>
    <description>Golang : https://github.com/goinggo</description>
    <link>http://www.goinggo.net/</link>
    <item>
        <pubDate>Sun, 15 Mar 2015 15:04:00 +0000</pubDate>
        <title>Object Oriented Programming Mechanics</title>
        <description>Go is an object oriented language.</description>
        <link>http://www.goinggo.net/2015/03/object-oriented</link>
    </item>
</channel>
</rss>`

// mockServer 返回用来处理请求的服务器的指针
func mockServer() *httptest.Server {
	f := func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		w.Header().Set("Content-Type", "application/xml")
		fmt.Fprintln(w, feed)
	}

	// 使用 httptest
	return httptest.NewServer(http.HandlerFunc(f))
}

// TestDownload 确认 http 包的 Get 函数可以下载内容
// 并且内容可以被正确地反序列化并关闭
func TestDownload(t *testing.T) {
	statusCode := http.StatusOK

	server := mockServer()
	defer server.Close()

	t.Log("Given the need to test downloading content.")
	{
		t.Logf("\tWhen checking \"%s\" for status code \"%d\"", server.URL, statusCode)
		{
			// http.Get 方法调用时并不知道我们的调用是否经过互联网
			// 这个 URL 是 MockServer 提供的
			resp, err := http.Get(server.URL)
			if err != nil {
				t.Fatal("\t\tShould be able to make the Get call.", ballotX, err)
			}
			t.Log("\t\tShould be able to make the Get call.", checkMark)

			defer resp.Body.Close()

			if resp.StatusCode != statusCode {
				t.Fatalf("\t\tShould receive a \"%d\" status. %v %v", statusCode, ballotX, resp.StatusCode)
			}
			t.Logf("\t\tShould receive a \"%d\" status. %v", statusCode, checkMark)

			// 从响应体中解码 XML 文档
			var d Document
			if err := xml.NewDecoder(resp.Body).Decode(&d); err != nil {
				t.Fatal("\t\tShould be able to unmarshal the response.", ballotX, err)
			}
			t.Log("\t\tShould be able to unmarshal the response.", checkMark)

			// 验证有一个 Item 对象
			if len(d.Channel.Items) == 1 {
				t.Log("\t\tShould have \"1\" item in the feed.", checkMark)
			} else {
				t.Error("\t\tShould have \"1\" item in the feed.", ballotX, len(d.Channel.Items))
			}
		}
	}
}

type Item struct {
	XMLName     xml.Name `xml:"item"`
	Title       string   `xml:"title"`
	Description string   `xml:"description"`
	Link        string   `xml:"link"`
}

type Channel struct {
	XMLName     xml.Name `xml:"channel"`
	Title       string   `xml:"title"`
	Description string   `xml:"description"`
	Link        string   `xml:"link"`
	PubDate     string   `xml:"pubDate"`
	Items       []Item   `xml:"item"`
}

type Document struct {
	XMLName xml.Name `xml:"rss"`
	Channel Channel  `xml:"channel"`
	URI     string
}

// 次包的名字也使用 _test 结尾
// 如果包使用这种方式命名，测试代码只能访问包里公开的标识符
// 即便测试代码文件和被测试的代码放在同一个文件夹中，也只能访问公开的标识符

// 这个示例程序展示如何测试内部服务端点的执行效果
package handlers_test

import (
	"GoInActionCode/chapter9/listing17/handlers"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

const (
	checkMark = "\u2713"
	ballotX   = "\u2717"
)

// 为服务端点初始化路由
func init() {
	handlers.Routes()
}

// TestSendJSON 测试 /sendjson 内部服务端点
func TestSendJSON(t *testing.T) {
	t.Log("Given the need to test the SendJSON endpoint.")
	{
		req, err := http.NewRequest("GET", "/sendjson", nil)
		if err != nil {
			t.Fatal("\tShould be able to create a request.", ballotX, err)
		}
		t.Log("\tShould be able to create a request.", checkMark)

		// 使用 httptest 包的 NewRecorder 函数创建一个 ResponseRecorder
		// 用于记录返回的响应
		rw := httptest.NewRecorder()

		// 调用服务默认的多路选择器（mux）的 ServeHttp 方法
		// 调用这个方法模仿了外部客户端对 /sendjson 服务端点的请求
		http.DefaultServeMux.ServeHTTP(rw, req)

		// 一旦 ServeHTTP 方法调用完成，http.ResponseRecorder 值就包含了 SendJSON 处理函数的响应

		if rw.Code != 200 {
			t.Fatal("\tShould receive \"200\"", ballotX, rw.Code)
		}
		t.Log("\tShould receive \"200\"", checkMark)

		u := struct {
			Name  string
			Email string
		}{}

		if err := json.NewDecoder(rw.Body).Decode(&u); err != nil {
			t.Fatal("\tShould decode the response.", ballotX)
		}
		t.Log("\tShould decode the response.", checkMark)

		if u.Name == "Bill" {
			t.Log("\tShould have a Name.", checkMark)
		} else {
			t.Error("\tShould have a Name.", ballotX, u.Name)
		}

		if u.Email == "bill@ardanstudios.com" {
			t.Log("\tShould have an Email.", checkMark)
		} else {
			t.Error("\tShould have an for Email.", ballotX, u.Email)
		}
	}
}

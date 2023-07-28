// 这个示例程序展示如何编写基础示例
package handlers_test

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
)

// * 需要使用 Example 代替 Test 作为函数名的开始
// * 示例代码的函数名字必须基于已经存在的公开的函数或者方法
// 此处是基于 handlers 包里公开的 SendJSON 函数
// ExampleSendJSON 提供了基础示例
func ExampleSendJSON() {
	r, _ := http.NewRequest("GET", "/sendjson", nil)
	w := httptest.NewRecorder()
	http.DefaultServeMux.ServeHTTP(w, r)

	var u struct {
		Name  string
		Email string
	}

	if err := json.NewDecoder(w.Body).Decode(&u); err != nil {
		log.Println("ERROR:", err)
	}

	// 使用 fmt 将结果写到 stdout 来检测输出
	// * 这个 Output: 标记用来在文档中标记出示例函数运行后期望的输出
	fmt.Println(u)
	// Output:
	// {Bill bill@ardanstudios.com}
}

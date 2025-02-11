// handlers 包提供了用于网络服务的服务端点
package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func init() {
	fmt.Println("init handlers")
}

// Routes 为网络服务设置路由
func Routes() {
	http.HandleFunc("/sendjson", SendJSON)
}

// SendJSON 返回一个简单的 JSON 文档
func SendJSON(rw http.ResponseWriter, r *http.Request) {
	u := struct {
		Name  string
		Email string
	}{
		Name:  "Bill",
		Email: "bill@ardanstudios.com",
	}

	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(200)
	json.NewEncoder(rw).Encode(&u)
}

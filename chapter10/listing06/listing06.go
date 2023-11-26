// 测试 mock
package main

import (
	"goinaction.zhangjin.me/chapter10/listing06/pubsub"
)

// mock pubsub 包的接口
type publisher interface {
	Publish(key string, v interface{}) error
	Subscribe(key string) error
}

// 这是需要 mock 的类型
type mock struct{}

// 给 mock 类型添加方法
func (m *mock) Publish(key string, v interface{}) error {
	// ADD YOUR MOCK FOR THE PUBLISH CALL.
	return nil
}

// 给 mock 类型添加方法
func (m *mock) Subscribe(key string) error {
	// ADD YOUR MOCK FOR THE SUBSCRIBE CALL.
	return nil
}

// 最终，mock 类型实现了 publisher 接口

func main() {
	pubs := []publisher{
		// PubSub 实现了 Publish 接口
		pubsub.New("localhost"),
		// mock 实现了 Subscribe 接口
		&mock{},
	}

	for _, p := range pubs {
		// 注意：原始的 Publish 方法和 mock 的 Publish 方法都会调用
		p.Publish("key", "value")
		// 调用 mock 的 Subscribe 方法
		p.Subscribe("key")
	}
}

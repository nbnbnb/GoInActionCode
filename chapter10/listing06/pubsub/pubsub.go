package pubsub

type PubSub struct {
	host string
}

// 创建一个 PubSub 类型的指针
func New(host string) *PubSub {
	ps := PubSub{
		host: host,
	}
	return &ps
}

// 原始的 Publish 方法
func (ps *PubSub) Publish(key string, v interface{}) error {
	// PRETEND THERE IS A SPECIFIC IMPLEMENTATION.
	return nil
}

// 原始的 Subscribe 方法
func (ps *PubSub) Subscribe(key string) error {
	// PRETEND THERE IS A SPECIFIC IMPLEMENTATION.
	return nil
}

package counters

// alertCounter 是一个未公开的类型
// 这个类型用于保存告警计数
type alertCounter int

// New 创建并返回一个未公开的 alertCounter 类型的值
func New(value int) alertCounter {
	// 相当于返回一个私有变量
	return alertCounter(value)
}

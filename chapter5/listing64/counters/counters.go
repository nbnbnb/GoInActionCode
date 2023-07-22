// counters 包提供告警计数器的功能
package counters

// alertCounter 是一个未公开的类型
// 这个类型用于保存告警计数

// 当一个标识符的名字以 "小写字母开头" 时
// 这个标识符就是未公开的，即包外的代码不可见
type alertCounter int

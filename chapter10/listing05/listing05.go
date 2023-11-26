// 编译时隐式接口转换
package main

import "fmt"

// =============================================================================

// 用户定义接口类型
type Mover interface {
	Move()
}

// 用户定义接口类型
type Locker interface {
	Lock()
	Unlock()
}

// 用户定义接口类型（聚合两个接口）
type MoveLocker interface {
	Mover
	Locker
}

// =============================================================================

// 用户定义类型
type bike struct{}

// =============================================================================

// 方法 - 给 bike 添加 Move 方法
func (bike) Move() {
	fmt.Println("Moving the bike")
}

// 方法 - 给 bike 添加 Lock 方法
func (bike) Lock() {
	fmt.Println("Locking the bike")
}

// 方法 - 给 bike 添加 Lock 方法
func (bike) Unlock() {
	fmt.Println("Unlocking the bike")
}

// 注意：方法是一个一个添加的，根据添加的方法，决定实现了那些接口能力

// =============================================================================

func main() {
	// 声明变量，nil 值
	var moveLocker MoveLocker
	var mover Mover

	// moveLocker 实现了 Mover/Locker/MoveLocker 接口
	moveLocker = bike{}

	// 可以赋值成功
	mover = moveLocker

	// 构建的是一个 MoveLocker 对象
	moveLockerAnother := mover.(bike)
	moveLocker = moveLockerAnother
}

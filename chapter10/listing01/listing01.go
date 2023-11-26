// 结构使用
package main

import (
	"errors"
	"fmt"
	"io"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

// 用户定义的类型
type Data struct {
	Line string
}

// 用户定义的类型
type System struct {
	Xenia
	Pillar
}

// 用户定义的类型
type Xenia struct{}

// 用户定义的类型
type Pillar struct{}

// 方法
// 给用户定义的类型添加行为 —— 方法（方法实际上也是函数）
// 在关键字 func 和方法名之间增加了一个参数
func (Xenia) Pull(d *Data) error {
	switch rand.Intn(10) {
	case 1, 9:
		return io.EOF

	case 5:
		return errors.New("Error reading data from Xenia")

	default:
		d.Line = "Data"
		fmt.Println("In:", d.Line)
		return nil
	}
}

// 方法
// 给用户定义的类型添加行为 —— 方法（方法实际上也是函数）
// 在关键字 func 和方法名之间增加了一个参数
func (Pillar) Store(d Data) error {
	fmt.Println("Out:", d.Line)
	return nil
}

// 函数 - 入参是 Xenia 自定义类型指针
func pull(x *Xenia, data []Data) (int, error) {
	for i := range data {
		// 调用 Pull 方法
		if err := x.Pull(&data[i]); err != nil {
			return i, err
		}
	}

	// 返回切片元素的实际大小，而不是总容量
	return len(data), nil
}

// 函数 - 入参是 Pillar 自定义类型指针
func store(p *Pillar, data []Data) (int, error) {
	for i, d := range data {
		// 调用 Store 方法
		if err := p.Store(d); err != nil {
			return i, err
		}
	}

	return len(data), nil
}

// 函数
func Copy(sys *System, batch int) error {
	// 初始化切片
	data := make([]Data, batch)

	// 隔离一个上下文
	for {
		// 调用 pull 函数
		i, err := pull(&sys.Xenia, data)

		if i > 0 {
			// 然后调用 store 函数
			if _, err := store(&sys.Pillar, data[:i]); err != nil {
				return err
			}
		}

		if err != nil {
			return err
		}
	}
}

// =============================================================================

func main() {
	// 声明 sys 类型的变量，并初始化所以字段
	// 注意结尾的 , 不能省略
	sys := System{
		Xenia:  Xenia{},
		Pillar: Pillar{},
	}

	// 调用 Copy 函数
	if err := Copy(&sys, 10); err != io.EOF {
		fmt.Println(err)
	}
}

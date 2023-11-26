// 接口使用
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

// 用户定义类型
type Data struct {
	Line string
}

// 用户定义类型（这个地方与 04 不同，包含的是两个类型）
type System struct {
	Xenia
	Pillar
}

// 用户定义类型
type Xenia struct{}

// 用户定义类型
type Pillar struct{}

// =============================================================================

// 用户定义接口
type Puller interface {
	Pull(d *Data) error
}

// 用户定义接口
type Storer interface {
	Store(d Data) error
}

// 用户定义接口类型（同时实现两个接口）
type PullStorer interface {
	Puller
	Storer
}

// =============================================================================

// 方法 - Xenia 实现了 Puller 接口
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

// 方法 - Pillar 实现了 Storer 接口
// 给用户定义的类型添加行为 —— 方法（方法实际上也是函数）
// 在关键字 func 和方法名之间增加了一个参数
func (Pillar) Store(d Data) error {
	fmt.Println("Out:", d.Line)
	return nil
}

// =============================================================================

// 函数 - 入参是 Puller 接口类型
func pull(p Puller, data []Data) (int, error) {
	for i := range data {
		if err := p.Pull(&data[i]); err != nil {
			return i, err
		}
	}

	return len(data), nil
}

// 函数 - 入参是 Storer 接口类型
func store(s Storer, data []Data) (int, error) {
	for i, d := range data {
		if err := s.Store(d); err != nil {
			return i, err
		}
	}

	return len(data), nil
}

// 函数 - 入参是 PullStorer 自定义接口（包括两个接口）
func Copy(ps PullStorer, batch int) error {
	// 初始化切片
	data := make([]Data, batch)

	// 隔离一个上下文
	for {
		// PullStorer 同时实现了 Puller/Storer 接口，所以符合 pull 函数签名
		// 调用 pull 函数
		i, err := pull(ps, data)
		if i > 0 {
			// PullStorer 同时实现了 Puller/Storer  接口，所以符合 store 函数签名
			// 然后调用 store 函数
			if _, err := store(ps, data[:i]); err != nil {
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

	// 隐含
	// 由于 System 有 Xenia 和 Pillar 两个字段，
	// 而 Xenia 和 Pillar 又分别实现了 Puller 和 Storer 接口
	// 所以 System 类型隐含实现了 PullStorer 接口

	// GO 语言里面的一种思想：给类型挂载方法，就等于类型实现了对应的接口
	sys := System{
		Xenia:  Xenia{},
		Pillar: Pillar{},
	}

	if err := Copy(&sys, 3); err != io.EOF {
		fmt.Println(err)
	}
}

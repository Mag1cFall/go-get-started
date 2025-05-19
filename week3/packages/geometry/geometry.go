// package geometry 定义了一个包含几何计算函数的包。
// 包名通常与其所在的目录名相同。
package geometry

import (
	"fmt"
	"math"
)

// Pi 是一个导出的常量 (首字母大写)
const Pi = math.Pi

// Rect struct (未导出，首字母小写)
// 这个结构体只能在 geometry 包内部使用。
type rect struct {
	width, height float64
}

// Circle struct (导出，首字母大写)
// 这个结构体可以在包外部使用。
type Circle struct {
	Radius float64
}

// Area 函数计算矩形面积 (导出，首字母大写)
// 注意：这个函数接收具体尺寸，而不是 rect 结构体，因为 rect 未导出。
// 如果我们想基于一个结构体计算，那个结构体也需要是导出的，或者提供一个导出的构造函数。
func Area(width, height float64) float64 {
	return width * height
}

// Perimeter 函数计算矩形周长 (导出)
func Perimeter(width, height float64) float64 {
	return 2*width + 2*height
}

// CircleArea 方法计算圆面积 (为导出的 Circle 类型定义的方法，因此也是导出的)
func (c Circle) CircleArea() float64 {
	return Pi * c.Radius * c.Radius
}

// internalHelperFunction (未导出，首字母小写)
// 这个函数只能在 geometry 包内部被调用。
func internalHelperFunction() {
	fmt.Println("这是一个内部帮助函数。")
	// 在实际应用中，fmt.Println 这样的副作用函数通常不直接放在库代码中，
	// 除非是明确的调试或日志输出。这里仅为演示。
}

// init 函数
// 每个包可以包含任意数量的 init 函数。
// init 函数不能被直接调用，它们在程序开始执行时，在 main 函数之前，
// 按照它们在包中声明的顺序自动执行。
// 如果一个包导入了其他包，则会先执行被导入包的 init 函数。
// init 函数通常用于执行包级别的初始化任务。
func init() {
	fmt.Println("geometry 包的 init 函数被调用了。")
	// internalHelperFunction() // 可以在这里调用包内函数
}

func init() {
	fmt.Println("geometry 包的第二个 init 函数被调用了 (按声明顺序)。")
}

// NewRect 是一个导出的构造函数，用于创建未导出的 rect 结构体的实例。
// 这是一个常见的模式，用于隐藏内部实现细节，同时提供可控的创建方式。
// 但由于 rect 结构体本身未导出，返回 *rect 意味着包外也无法直接使用其字段。
// 通常，如果想让外部使用结构体，结构体本身及其需要访问的字段都应该是导出的。
//
// 为了更好地演示包的使用，我们将修改 rect 为导出的 Rectangle，
// 或者提供操作这些内部结构体的导出函数。
//
// 让我们将 rect 修改为导出的 Rectangle 来简化示例。
// （下面的代码将不会使用 NewRect 和 rect，而是使用下面定义的 Rectangle）
/*
func NewRect(w, h float64) *rect {
    if w < 0 || h < 0 {
        return nil // 或者返回错误
    }
    return &rect{width: w, height: h}
}
*/

// Rectangle 结构体 (导出)
type Rectangle struct {
	Width  float64
	Height float64
}

// NewRectangle 是一个构造函数，用于创建 Rectangle 实例
func NewRectangle(width, height float64) (Rectangle, error) {
	if width < 0 || height < 0 {
		return Rectangle{}, fmt.Errorf("宽度和高度不能为负数: width=%.2f, height=%.2f", width, height)
	}
	return Rectangle{Width: width, Height: height}, nil
}

// Methods for exported Rectangle

// Area method for Rectangle (exported)
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

// Perimeter method for Rectangle (exported)
func (r Rectangle) Perimeter() float64 {
	return 2*r.Width + 2*r.Height
}

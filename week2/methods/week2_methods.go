package main

import (
	"fmt"
	"math"
)

// --- 1. 定义结构体 ---
// 我们将为这些结构体定义方法

type Rectangle struct {
	Width  float64
	Height float64
}

type Circle struct {
	Radius float64
}

type Point struct {
	X, Y int
}

// --- 2. 定义方法 ---
// 方法是带有特殊接收者参数的函数。
// 接收者出现在 func 关键字和方法名之间。
// func (receiverName ReceiverType) MethodName(parameters) (returnTypes) { ... }

// Area 方法，接收者是 Rectangle 类型的值
// (r Rectangle) 被称为接收者，r 是接收者变量名，Rectangle 是接收者类型。
// 当调用 r.Area() 时，r 的一个副本会传递给这个方法。
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

// Perimeter 方法，接收者是 Rectangle 类型的值
func (r Rectangle) Perimeter() float64 {
	return 2*r.Width + 2*r.Height
}

// Scale 方法，接收者是 Rectangle 类型的指针
// (r *Rectangle) 表示接收者是一个指向 Rectangle 的指针。
// 这允许方法修改原始的 Rectangle 实例。
func (r *Rectangle) Scale(factor float64) {
	if r == nil {
		fmt.Println("不能对nil的Rectangle进行缩放")
		return
	}
	r.Width *= factor
	r.Height *= factor
	// 注意：即使接收者是指针，调用时仍然使用 r.Scale(2) 而不是 (*r).Scale(2)
	// Go 会自动处理指针的解引用。
}

// Area 方法，接收者是 Circle 类型的值
func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

// Circumference 方法，接收者是 Circle 类型的值
func (c Circle) Circumference() float64 {
	return 2 * math.Pi * c.Radius
}

// ChangeRadius 方法，接收者是 Circle 类型的指针
func (c *Circle) ChangeRadius(newRadius float64) {
	if c == nil {
		return
	}
	c.Radius = newRadius
}

// --- 值接收者 vs 指针接收者 ---
// 1. 值接收者 (e.g., func (r Rectangle) Area()):
//    - 方法操作的是接收者变量的一个副本。
//    - 对副本的修改不会影响原始值。
//    - 如果类型是结构体且较大，或者方法不需要修改原始值，可以使用值接收者。
//    - 可以用值类型或指针类型调用值接收者的方法 (Go会自动解引用指针)。

// 2. 指针接收者 (e.g., func (r *Rectangle) Scale()):
//    - 方法操作的是原始值的引用（通过指针）。
//    - 对接收者的修改会影响原始值。
//    - 当方法需要修改接收者时，必须使用指针接收者。
//    - 对于大型结构体，使用指针接收者可以避免复制整个结构体，从而提高效率。
//    - 如果一个类型的一些方法使用指针接收者，那么通常该类型的所有方法都应该使用指针接收者，以保持一致性。
//    - 可以用值类型或指针类型调用指针接收者的方法 (Go会自动取地址)。

// Distance 方法，接收者是 Point 类型的值
func (p Point) Distance(q Point) float64 {
	dx := float64(p.X - q.X)
	dy := float64(p.Y - q.Y)
	return math.Sqrt(dx*dx + dy*dy)
}

// Move 方法，接收者是 Point 类型的指针
func (p *Point) Move(dx, dy int) {
	if p == nil {
		return
	}
	p.X += dx
	p.Y += dy
}

func main() {
	fmt.Println("--- 第2周学习：方法 (Methods) ---")

	// --- 3. 调用方法 ---
	fmt.Println("\n--- 3. 调用方法 ---")
	rect1 := Rectangle{Width: 10, Height: 5}
	fmt.Printf("矩形 rect1: %+v\n", rect1) // %+v 会打印字段名
	fmt.Println("rect1 的面积:", rect1.Area())
	fmt.Println("rect1 的周长:", rect1.Perimeter())

	// 调用指针接收者的方法
	// 可以直接在值类型上调用指针接收者的方法，Go会自动取地址 (&rect1).Scale(2)
	rect1.Scale(2)
	fmt.Printf("rect1 缩放2倍后: %+v\n", rect1)
	fmt.Println("缩放后 rect1 的面积:", rect1.Area())

	// 也可以显式使用指针
	rect2Ptr := &Rectangle{Width: 3, Height: 4}
	fmt.Printf("矩形 rect2Ptr: %+v (这是一个指针)\n", rect2Ptr)
	fmt.Println("rect2Ptr 指向的矩形面积:", rect2Ptr.Area()) // Go会自动解引用 (*rect2Ptr).Area()
	rect2Ptr.Scale(3)
	fmt.Printf("rect2Ptr 缩放3倍后: %+v\n", rect2Ptr)

	// 对于 nil 指针接收者
	var nilRect *Rectangle
	// fmt.Println("nilRect 的面积 (会panic吗?):", nilRect.Area()) // 这行会引发 panic，如下解释
	// 上一行代码会引发 panic，因为：
	// 1. nilRect 是一个 *Rectangle 类型的 nil 指针。
	// 2. Area() 方法是 func (r Rectangle) Area() float64，它是一个值接收者。
	// 3. 当通过一个指针调用值接收者方法时 (nilRect.Area())，Go 会尝试解引用指针 (*nilRect) 来获取值。
	// 4. 对 nil 指针解引用会导致 "invalid memory address or nil pointer dereference" panic。
	// 正确的做法是先检查指针是否为 nil。
	if nilRect != nil {
		fmt.Println("nilRect 的面积:", nilRect.Area())
	} else {
		fmt.Println("nilRect 是 nil，无法调用其 Area() 方法获取面积 (因为 Area 是值接收者)。")
	}

	nilRect.Scale(2) // 我们在 Scale 方法内部加了 nil 检查，所以这行是安全的

	circ1 := Circle{Radius: 5}
	fmt.Printf("圆形 circ1: %+v\n", circ1)
	fmt.Printf("circ1 的面积: %.2f\n", circ1.Area())
	fmt.Printf("circ1 的周长: %.2f\n", circ1.Circumference())

	circ1.ChangeRadius(7) // 值类型调用指针接收者方法
	fmt.Printf("circ1 半径改变后: %+v\n", circ1)
	fmt.Printf("改变半径后 circ1 的面积: %.2f\n", circ1.Area())

	// Point 示例
	pA := Point{X: 1, Y: 2}
	pB := Point{X: 4, Y: 6}
	fmt.Printf("点 pA: %+v, 点 pB: %+v\n", pA, pB)
	fmt.Printf("pA 和 pB 之间的距离: %.2f\n", pA.Distance(pB))

	pA.Move(10, 20) // 值类型调用指针接收者方法
	fmt.Printf("pA 移动后: %+v\n", pA)

	fmt.Println("\n--- 方法学习结束 ---")
}

// 注意: 对于值接收者的方法，如 Rectangle.Area()
// 如果你有一个指针 *Rectangle，你仍然可以调用 rectPtr.Area()
// Go 会自动将其解引用为 (*rectPtr).Area()

// 同样，对于指针接收者的方法，如 (*Rectangle).Scale()
// 如果你有一个值 Rectangle，你仍然可以调用 rectValue.Scale()
// Go 会自动取其地址为 (&rectValue).Scale()

// 唯一的例外是当接收者本身就是一个接口类型时，规则会更严格。
// (接口将在下一周学习)

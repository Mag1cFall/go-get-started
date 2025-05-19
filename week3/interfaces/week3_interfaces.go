package main

import (
	"fmt"
	"math"
)

// --- 1. 什么是接口？---
// 接口类型是由一组方法签名定义的集合。
// 一个类型如果实现了接口中定义的所有方法，那么它就隐式地实现了该接口。
// 不需要像其他语言那样显式声明 "implements InterfaceName"。

// 定义一个 Shape 接口，它有一个 Area() 方法
type Shape interface {
	Area() float64 // 方法签名：方法名 Area，无参数，返回 float64
}

// 定义一个 Stringer 接口 (类似 fmt.Stringer)
// 如果一个类型实现了 String() string 方法，它就可以自定义其字符串表示形式
// fmt.Println 等函数会自动调用这个方法
type Stringer interface {
	String() string
}

// --- 2. 实现接口 ---

// Rectangle 结构体
type Rectangle struct {
	Width  float64
	Height float64
}

// Rectangle 类型实现 Shape 接口的 Area() 方法
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

// Rectangle 类型实现 Stringer 接口的 String() 方法
func (r Rectangle) String() string {
	return fmt.Sprintf("Rectangle (Width: %.2f, Height: %.2f)", r.Width, r.Height)
}

// Circle 结构体
type Circle struct {
	Radius float64
}

// Circle 类型实现 Shape 接口的 Area() 方法
func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

// Circle 类型实现 Stringer 接口的 String() 方法
func (c Circle) String() string {
	return fmt.Sprintf("Circle (Radius: %.2f)", c.Radius)
}

// Triangle 结构体 (不实现 Stringer 接口，用于对比)
type Triangle struct {
	Base   float64
	Height float64
}

// Triangle 类型实现 Shape 接口的 Area() 方法
func (t Triangle) Area() float64 {
	return 0.5 * t.Base * t.Height
}

// --- 3. 使用接口类型 ---
// 接口类型可以作为函数的参数、返回值或变量的类型。
// 这允许我们编写更通用的代码，可以处理实现了特定接口的任何类型。

// printShapeInfo 函数接收一个 Shape 接口类型的参数
// 它可以接收任何实现了 Area() float64 方法的类型
func printShapeInfo(s Shape) {
	fmt.Printf("  形状的面积是: %.2f\n", s.Area())
	// 我们不能直接访问 s.Width 或 s.Radius，因为 Shape 接口只定义了 Area() 方法
	// 要访问具体类型的字段，需要使用类型断言 (后面会讲)

	// 尝试打印形状的字符串表示
	// 如果 s 也实现了 Stringer 接口，fmt.Println 会自动调用其 String() 方法
	fmt.Printf("  形状的描述: %s\n", s) // %s 会尝试调用 String()
}

func main() {
	fmt.Println("--- 第3周学习：接口 (Interfaces) ---")

	rect := Rectangle{Width: 10, Height: 5}
	circ := Circle{Radius: 7}
	tri := Triangle{Base: 4, Height: 6}

	fmt.Println("\n--- 使用具体类型调用方法 ---")
	fmt.Println(rect.String(), "面积:", rect.Area()) // 调用 Rectangle 的 String() 和 Area()
	fmt.Println(circ.String(), "面积:", circ.Area()) // 调用 Circle 的 String() 和 Area()
	// fmt.Println(tri.String()) // 这会报错，因为 Triangle 没有 String() 方法
	fmt.Printf("Triangle (Base: %.2f, Height: %.2f) 面积: %.2f\n", tri.Base, tri.Height, tri.Area())

	fmt.Println("\n--- 使用接口类型 ---")
	// 创建一个 Shape 接口类型的变量
	var s1 Shape
	s1 = rect // Rectangle 实现了 Shape 接口，所以可以赋值
	fmt.Println("s1 (Rectangle) 的面积:", s1.Area())

	s2 := Shape(circ) // Circle 实现了 Shape 接口
	fmt.Println("s2 (Circle) 的面积:", s2.Area())

	// 使用 printShapeInfo 函数
	fmt.Println("调用 printShapeInfo(rect):")
	printShapeInfo(rect)

	fmt.Println("调用 printShapeInfo(circ):")
	printShapeInfo(circ)

	fmt.Println("调用 printShapeInfo(tri):")
	printShapeInfo(tri) // Triangle 也实现了 Shape 接口

	// --- 4. 空接口 interface{} ---
	// 空接口类型 interface{} 不包含任何方法。
	// 因此，任何类型都隐式地实现了空接口。
	// 空接口可以用来存储任何类型的值，类似于其他语言中的 Object 或 any 类型。
	fmt.Println("\n--- 4. 空接口 interface{} ---")
	var anyType interface{}

	anyType = 100
	fmt.Printf("空接口存储 int: 值=%v, 类型=%T\n", anyType, anyType)

	anyType = "Hello Go"
	fmt.Printf("空接口存储 string: 值=%v, 类型=%T\n", anyType, anyType)

	anyType = Circle{Radius: 3}
	fmt.Printf("空接口存储 Circle: 值=%v, 类型=%T\n", anyType, anyType) // Circle有String()方法

	// --- 5. 类型断言 (Type Assertion) ---
	// 当我们有一个接口类型的值时，有时需要将其转换回其原始的具体类型，以便访问其特有的字段或方法。
	// 类型断言的语法：value, ok := interfaceValue.(ConcreteType)
	// - 如果 interfaceValue 持有的确实是 ConcreteType，则 ok 为 true，value 是转换后的具体类型的值。
	// - 如果不是，则 ok 为 false，value 是 ConcreteType 的零值 (不会 panic)。
	// 另一种语法：value := interfaceValue.(ConcreteType)
	// - 如果转换失败，会直接 panic。通常不推荐，除非你非常确定类型。
	fmt.Println("\n--- 5. 类型断言 ---")

	var shapeForAssertion Shape = Circle{Radius: 5.5}

	// 尝试断言为 Circle
	c, ok := shapeForAssertion.(Circle)
	if ok {
		fmt.Printf("断言成功: 这是一个 Circle，半径是 %.2f\n", c.Radius)
		// 现在可以访问 Circle 特有的字段，如 c.Radius
	} else {
		fmt.Println("断言失败: 不是 Circle 类型")
	}

	// 尝试断言为 Rectangle (会失败)
	r, ok := shapeForAssertion.(Rectangle)
	if ok {
		fmt.Printf("断言成功: 这是一个 Rectangle，宽度是 %.2f\n", r.Width)
	} else {
		fmt.Println("断言失败: 不是 Rectangle 类型 (shapeForAssertion 当前是 Circle)")
	}

	// 使用 panic 版本的类型断言 (如果类型不匹配会 panic)
	// c2 := shapeForAssertion.(Rectangle) // 这行会 panic，因为 shapeForAssertion 是 Circle
	// fmt.Println(c2)

	// --- 6. Type Switch (类型选择) ---
	// Type Switch 是一种更优雅地处理多种可能的具体类型的方式。
	// 语法类似普通的 switch 语句，但在 case 中使用类型。
	fmt.Println("\n--- 6. Type Switch ---")
	checkType(123)
	checkType("Go Language")
	checkType(Rectangle{Width: 2, Height: 3})
	checkType(Circle{Radius: 1.5})
	checkType(3.14)
	checkType(nil) // 注意 nil 的情况

	// 接口值可以是 nil
	var nilShape Shape
	fmt.Println("nilShape:", nilShape, "是否为 nil?", nilShape == nil) // true
	// nilShape.Area() // 这会导致 panic: runtime error: invalid memory address or nil pointer dereference

	if nilShape != nil {
		fmt.Println("nilShape 的面积:", nilShape.Area())
	}

	fmt.Println("\n--- 接口学习结束 ---")
}

func checkType(i interface{}) { // i 是一个空接口，可以接收任何类型
	fmt.Printf("  检查类型: 值=%v, ", i)
	switch v := i.(type) { // v 会是转换后的具体类型的值
	case int:
		fmt.Printf("是 int 类型, 值为 %d\n", v)
	case string:
		fmt.Printf("是 string 类型, 值为 \"%s\"\n", v)
	case Rectangle:
		fmt.Printf("是 Rectangle 类型, 面积为 %.2f\n", v.Area()) // v 是 Rectangle 类型
	case Circle:
		fmt.Printf("是 Circle 类型, 面积为 %.2f\n", v.Area()) // v 是 Circle 类型
	case nil:
		fmt.Println("是 nil") // 当接口变量本身为 nil 时
	default:
		fmt.Printf("是未知类型 %T\n", v) // %T 打印类型
	}
}

// 接口也可以嵌入其他接口
type ReadWriter interface {
	Reader
	Writer
}
type Reader interface {
	Read(p []byte) (n int, err error)
}
type Writer interface {
	Write(p []byte) (n int, err error)
}

// 任何实现了 Read 和 Write 方法的类型都实现了 ReadWriter 接口。

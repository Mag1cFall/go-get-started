package main // 或一个可测试的包名，例如 "mathops"

import "fmt" // <--- 添加 fmt 包导入

// DivisionByZeroError 是一个自定义错误类型，用于表示除以零的错误
type DivisionByZeroError struct{}

// Error 使 DivisionByZeroError 类型实现了 error 接口
func (e *DivisionByZeroError) Error() string {
	return "division by zero"
}

// Add 函数返回两个整数的和
func Add(a, b int) int {
	return a + b
}

// Subtract 函数返回两个整数的差
func Subtract(a, b int) int {
	return a - b
}

// Multiply 函数返回两个整数的积
func Multiply(a, b int) int {
	return a * b
}

// Divide 函数返回两个整数的商。
// 如果除数为0，它会返回0和一个错误。
func Divide(a, b int) (int, error) {
	if b == 0 {
		// 通常我们会返回一个更具体的错误类型或使用 errors.New
		// 为了简单起见，这里直接用 fmt.Errorf
		// 但在一个真正的库中，我们可能不希望直接依赖 fmt 打印错误
		// 因此，更好的做法是返回一个 sentinel error 或自定义错误类型
		// (不过，为了让这个文件能独立运行为 main 包，暂时这样写)
		// return 0, fmt.Errorf("division by zero")
		//
		// 如果这个文件不是 main 包，而是库，则应该：
		// import "errors"
		// return 0, errors.New("division by zero")

		// 为了让 main 包能运行而不必处理 fmt 的导入（如果 main 函数为空），
		// 并且不引入 errors 包（因为我们还没写 _test.go 来使用它），
		// 我们暂时让它在除以0时 panic，测试代码会处理或预期这个panic。
		// 或者，我们可以在这里返回一个预定义的错误变量。
		//
		// 让我们选择返回一个错误，并确保 main 函数（如果存在）或测试能处理它。
		// 为了让这个文件本身不依赖 fmt 或 errors（如果它只是被测试），
		// 并且如果它被作为 main 包运行，main 函数可以打印。
		// 我们暂时不导出错误，让测试来定义预期的行为。
		//
		// 修正：为了让这个文件能独立编译（即使没有main函数），
		// 并且能被测试，我们还是用 errors.New。
		// 测试文件会导入 errors。
		//
		// 再次修正：如果这个文件是 package main，并且没有 main 函数，它不会被编译。
		// 如果有 main 函数，它需要导入 errors。
		// 如果它是一个独立的包 (e.g. package mathops)，那么它可以导入 errors。
		//
		// 为了简单起见，并专注于测试，我们假设这个文件是 `package main`，
		// 但主要目的是被 `_test.go` 文件测试。
		// 我们会在测试中处理错误。
		//
		// 最终决定：为了这个文件能被简单地 `go run` (如果添加main函数)，
		// 并且能被测试，我们还是让它返回 error。
		// 在 _test.go 中，我们会导入 errors。
		// 如果单独运行此文件（添加main），main中也需要导入 errors。
		//
		// 最简单的做法，为了让这个文件本身不报错（如果单独看），
		// 且能被测试，就是 panic，然后测试用例可以测试 panic。
		// 但大纲要求测试，所以返回 error 更符合测试场景。
		//
		// 我们将在这里返回一个 error，并假设测试文件会处理它。
		// 如果没有 main 函数，这个文件本身不会被直接运行。
		// 它的目的是被 _test.go 文件导入（如果它是不同的包）或一起编译（如果是同一个包）。
		//
		// 假设这是 package main，_test.go 也是 package main。
		return 0, &DivisionByZeroError{} // <--- 返回在包级别定义的错误类型
	}
	return a / b, nil
}

// main 函数用于简单演示，但主要目的是被测试。
// 如果你只想运行测试，可以注释掉或删除 main 函数。
func main() {
	fmt.Println("--- 数学运算函数 (主要用于测试) ---")
	a, b := 10, 5
	fmt.Printf("%d + %d = %d\n", a, b, Add(a, b))
	fmt.Printf("%d - %d = %d\n", a, b, Subtract(a, b))
	fmt.Printf("%d * %d = %d\n", a, b, Multiply(a, b))

	res, err := Divide(a, b)
	if err != nil {
		fmt.Printf("%d / %d 错误: %v\n", a, b, err)
	} else {
		fmt.Printf("%d / %d = %d\n", a, b, res)
	}

	res, err = Divide(a, 0)
	if err != nil {
		fmt.Printf("%d / %d 错误: %v\n", a, 0, err)
	} else {
		fmt.Printf("%d / %d = %d\n", a, 0, res)
	}
}

// 为了让上面的 main 函数能工作，需要导入 fmt
// 但如果我们主要目的是写一个可测试的 "库" (即使是 main 包的一部分被测试)，
// 这个文件本身可能不应该有 main 函数，或者 main 函数应该有条件编译。
//
// 为了让 _test.go 能顺利编译和测试这些函数，
// 这个文件声明为 package main 是可以的，测试文件也会是 package main。
// 我将添加 fmt 的导入，以便 main 函数能工作。

package main // 测试文件通常与其被测试的源文件在同一个包

import (
	"testing" // 导入 testing 包
	// "errors" // 如果我们要与 errors.New 创建的错误比较，可能需要
)

// --- 1. 单元测试 (Unit Tests) ---
// 测试函数必须以 Test 开头，后跟一个大写字母开头的单词或词组。
// 测试函数接收一个参数 t *testing.T。
// 使用 t.Errorf, t.Fatalf, t.Logf 等方法来报告测试结果。
// - t.Errorf(...): 报告错误，但继续执行当前测试函数。
// - t.Fatalf(...): 报告错误，并立即停止当前测试函数的执行。
// - t.Logf(...): 记录信息 (通常在 -v 模式下显示)。

// TestAdd 测试 Add 函数
func TestAdd(t *testing.T) {
	// 可以使用表格驱动测试 (table-driven tests) 来组织多个测试用例
	testCases := []struct {
		name     string // 测试用例名称
		a, b     int    // 输入
		expected int    // 期望输出
	}{
		{"positive numbers", 2, 3, 5},
		{"negative numbers", -2, -3, -5},
		{"mixed numbers", 2, -3, -1},
		{"zero", 0, 0, 0},
		{"add to zero", 5, 0, 5},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) { // t.Run 创建子测试，有助于组织和筛选测试
			result := Add(tc.a, tc.b)
			if result != tc.expected {
				t.Errorf("Add(%d, %d) = %d; want %d", tc.a, tc.b, result, tc.expected)
			}
		})
	}
}

// TestSubtract 测试 Subtract 函数
func TestSubtract(t *testing.T) {
	result := Subtract(5, 3)
	expected := 2
	if result != expected {
		t.Errorf("Subtract(5, 3) = %d; want %d", result, expected)
	}

	result = Subtract(3, 5)
	expected = -2
	if result != expected {
		t.Errorf("Subtract(3, 5) = %d; want %d", result, expected)
	}
}

// TestMultiply 测试 Multiply 函数
func TestMultiply(t *testing.T) {
	tables := []struct {
		a, b, expected int
	}{
		{2, 3, 6},
		{-2, 3, -6},
		{2, 0, 0},
	}

	for _, table := range tables {
		if result := Multiply(table.a, table.b); result != table.expected {
			t.Errorf("Multiply(%d, %d) = %d; want %d", table.a, table.b, result, table.expected)
		}
	}
}

// TestDivide 测试 Divide 函数
func TestDivide(t *testing.T) {
	t.Run("successful division", func(t *testing.T) {
		result, err := Divide(10, 2)
		if err != nil {
			t.Fatalf("Divide(10, 2) returned an unexpected error: %v", err)
		}
		if result != 5 {
			t.Errorf("Divide(10, 2) = %d; want 5", result)
		}
	})

	t.Run("division by zero", func(t *testing.T) {
		_, err := Divide(10, 0)
		if err == nil {
			t.Fatalf("Divide(10, 0) expected an error, but got nil")
		}
		// 我们可以检查错误的具体类型或内容
		// 在 math_operations.go 中，我们返回了 *DivisionByZeroError
		if _, ok := err.(*DivisionByZeroError); !ok {
			t.Errorf("Divide(10, 0) returned error of type %T; want *DivisionByZeroError", err)
		}
		// 或者检查错误消息
		// expectedErrorMsg := "division by zero"
		// if err.Error() != expectedErrorMsg {
		// 	t.Errorf("Divide(10, 0) error message = \"%s\"; want \"%s\"", err.Error(), expectedErrorMsg)
		// }
	})

	t.Run("division result check", func(t *testing.T) {
		result, err := Divide(9, 3)
		if err != nil {
			t.Fatalf("Divide(9, 3) returned an unexpected error: %v", err)
		}
		if result != 3 {
			t.Errorf("Divide(9, 3) = %d; want 3", result)
		}
	})
}

// --- 2. 基准测试 (Benchmark Tests) --- (可选，初步了解)
// 基准测试函数必须以 Benchmark 开头。
// 它们接收一个 *testing.B 参数。
// 循环体 b.N 次来执行被测试的代码。b.N 会由测试框架自动调整。

// BenchmarkAdd 基准测试 Add 函数
func BenchmarkAdd(b *testing.B) {
	// b.ReportAllocs() // 可以报告内存分配情况
	for i := 0; i < b.N; i++ { // 循环 b.N 次
		Add(100, 200)
	}
}

// 运行测试：
// - 在当前目录下 (week8/testing_examples/) 打开终端。
// - 运行所有测试：go test
// - 运行详细输出：go test -v
// - 运行特定测试函数：go test -v -run TestDivide
// - 运行特定子测试：go test -v -run TestDivide/division_by_zero
// - 运行基准测试：go test -bench=.  (或者 go test -bench Add 来运行 BenchmarkAdd)
//   注意：基准测试默认不运行，需要 -bench 标志。
// - 同时运行测试和基准测试：go test -v -bench=.

package main

import "fmt"

// Calculator 简单计算器
type Calculator struct{}

// Add 加法
func (c *Calculator) Add(a, b int) int {
	return a + b
}

// Subtract 减法
func (c *Calculator) Subtract(a, b int) int {
	return a - b
}

// Multiply 乘法
func (c *Calculator) Multiply(a, b int) int {
	return a * b
}

// Divide 除法
func (c *Calculator) Divide(a, b int) int {
	if b == 0 {
		return 0
	}
	return a / b
}

// Power 幂运算
func (c *Calculator) Power(base, exponent int) int {
	result := 1
	for i := 0; i < exponent; i++ {
		result *= base
	}
	return result
}

func main() {
	calc := &Calculator{}

	fmt.Println("=== 计算器测试 ===")
	fmt.Printf("2 + 3 = %d\n", calc.Add(2, 3))
	fmt.Printf("5 - 2 = %d\n", calc.Subtract(5, 2))
	fmt.Printf("3 * 4 = %d (期望 12)\n", calc.Multiply(3, 4))
	fmt.Printf("10 / 2 = %d\n", calc.Divide(10, 2))
	fmt.Printf("2 ^ 3 = %d (期望 8)\n", calc.Power(2, 3))
}

package main

import "fmt"

func main() {
	fmt.Println(add07(30, 60))
}

func add07(num1 int, num2 int) int {
	// GO遇到defer关键字，不会立即执行其后的语句，而是将defer后的语句和相关的值压入栈中，继续执行函数后面的语句。在函数执行完毕后，从栈中取出语句执行
	defer fmt.Println("num1=", num1)
	defer fmt.Println("num2=", num2)
	num1 += 10
	num2 += 10
	var sum int = num1 + num2
	fmt.Println("sum=", sum)
	return sum
}

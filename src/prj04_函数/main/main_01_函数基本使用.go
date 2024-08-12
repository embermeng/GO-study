package main

import "fmt"

/*
	func 函数名(形参列表) (返回值类型列表) {
		...
		return 返回值列表
	}
*/
func add(num1 int, num2 int) int { // 如果返回值类型只有一个，可以省略括号
	return num1 + num2
}

// 计算两个数的和、差
func cal(num1 int, num2 int) (int, int) { // 如果返回值类型只有一个，可以省略括号
	return num1 + num2, num1 - num2
}

func main() {
	sum := add(10, 20)
	fmt.Println(sum)

	res1, res2 := cal(30, 50)
	fmt.Println(res1, res2)
}

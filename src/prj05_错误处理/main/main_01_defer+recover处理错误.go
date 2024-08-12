package main

import "fmt"

func main() {
	testErr()
	fmt.Println("继续执行后续逻辑")
}

func testErr() {
	// 利用defer + recover来捕获错误：defer后加匿名函数调用
	defer func() {
		// recover内置函数可以捕获错误
		err := recover()
		// 如果没有捕获错误，返回nil
		if err != nil {
			fmt.Println("错误已经捕获：", err)
		}
	}()

	num1 := 10
	num2 := 0
	res := num1 / num2
	fmt.Println(res)
}

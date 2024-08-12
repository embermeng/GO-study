package main

import "fmt"

/*执行顺序：1. 全局变量的定义 2. init函数的调用 3. main函数的调用*/
var num int = test05()

func test05() int {
	fmt.Println("test函数被执行")
	return 100
}

func init() {
	fmt.Println("init函数被执行")
}

func main() {
	fmt.Println("main函数被执行")
}

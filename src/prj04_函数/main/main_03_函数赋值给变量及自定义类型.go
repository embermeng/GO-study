package main

import "fmt"

func test2(num int) {
	fmt.Println(num)
}

func main() {
	a := test2
	fmt.Printf("a的类型: %T, test2函数的类型: %T\n", a, test2) //a的类型: func(int), test2函数的类型: func(int)
	a(10)

	//自定义数据类型
	type myInt int
	var num1 myInt = 80
	fmt.Println("num1", num1)

	var num2 int = int(num1) //虽然是别名，但是类型还是不同，要转换类型
	fmt.Println("num2", num2)
}

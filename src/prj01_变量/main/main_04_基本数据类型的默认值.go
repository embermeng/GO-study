package main

import "fmt"

func main() {
	//基本数据类型的默认值
	var a int
	var b float64
	var c bool
	var d string
	fmt.Printf("a = %d, b = %f, c = %t, d = %s\n", a, b, c, d)
	//%v表示按照变量的值输出
	fmt.Printf("a = %d, b = %v, c = %v, d = %v\n", a, b, c, d)
}
